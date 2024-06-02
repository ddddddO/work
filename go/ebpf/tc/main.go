package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	// "github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
	// "github.com/PraserX/ipconv"
	"github.com/cilium/ebpf"
	"github.com/vishvananda/netlink"

	"golang.org/x/sys/unix"
)

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
	    log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs egress_packetObjects
	if err := loadEgress_packetObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	if err := attachFilter("eth0", objs.egress_packetPrograms.ShowIcmp); err != nil {
		log.Fatal("Failed to attach:", err)
	}

	tick := time.Tick(time.Second)
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)
	for {
		select {
		case <-tick:
			var count uint64
			err := objs.PktCount.Lookup(uint32(0), &count)
			if err != nil {
				log.Fatal("Map lookup:", err)
			}
			log.Printf("Sent %d packets", count)

			var arpCount uint64
			err = objs.ArpPktCount.Lookup(uint32(0), &arpCount)
			if err != nil {
				log.Fatal("ARP Map lookup:", err)
			}
			log.Printf("ARP Sent %d packets", arpCount)
		case <-stop:
			log.Print("Received signal, exiting..")
			return
		}
	}
}

// https://github.com/fedepaol/tc-return/blob/main/main.go
func attachFilter(attachTo string, program *ebpf.Program) error {
	devID, err := net.InterfaceByName(attachTo)
	if err != nil {
		return fmt.Errorf("could not get interface ID: %w", err)
	}

	qdisc := &netlink.GenericQdisc{
		QdiscAttrs: netlink.QdiscAttrs{
			LinkIndex: devID.Index,
			// Handle:    netlink.MakeHandle(0xffff, 0),
			Parent: netlink.HANDLE_CLSACT,
		},
		QdiscType: "clsact",
	}

	err = netlink.QdiscReplace(qdisc)
	if err != nil {
		return fmt.Errorf("could not get replace qdisc: %w", err)
	}

	filter := &netlink.BpfFilter{
		FilterAttrs: netlink.FilterAttrs{
			LinkIndex: devID.Index,
			Parent:    netlink.HANDLE_MIN_EGRESS,
			Handle:    1,
			Protocol:  unix.ETH_P_ALL,
		},
		Fd:           program.FD(),
		Name:         program.String(),
		DirectAction: true,
	}

	if err := netlink.FilterReplace(filter); err != nil {
		return fmt.Errorf("failed to replace tc filter: %w", err)
	}
	return nil
}
