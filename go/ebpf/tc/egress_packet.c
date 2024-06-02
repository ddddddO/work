//go:build ignore

// 部分的にコピーしてる
// https://eunomia.dev/tutorials/20-tc/#writing-ebpf-programs

// 以下あたりの定義も拝借. ethhdr/iphdr など
// https://github.com/cilium/ebpf/blob/b8dc0ee25417ce7cd4a6feb48be42c0615ee9043/examples/headers/common.h#L4

// #include <vmlinux.h>
// #include "common.h"
#include <linux/bpf.h>
#include <bpf/bpf_endian.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

#define TC_ACT_OK 0
#define TC_ACT_SHOT 2
#define ETH_P_IP 0x0800
#define ETH_P_ARP 0x0806
#define MAX_ENTRIES 64
#define AF_INET		2

struct ethhdr {
	unsigned char h_dest[6];
	unsigned char h_source[6];
	__be16 h_proto;
};

struct iphdr {
	__u8 ihl: 4;
	__u8 version: 4;
	__u8 tos;
	__be16 tot_len;
	__be16 id;
	__be16 frag_off;
	__u8 ttl;
	__u8 protocol;
	__sum16 check;
	__be32 saddr;
	__be32 daddr;
};

char __license[] SEC("license") = "Dual MIT/GPL";

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY); 
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 1);
} pkt_count SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY); 
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 1);
} arp_pkt_count SEC(".maps");

// __sk_buff について
// https://medium.com/@c0ngwang/understanding-struct-sk-buff-730cf847a722

SEC("tc")
int control_egress(struct __sk_buff *skb)
{
    void *data_end = (void *)(__u64)skb->data_end;
    void *data = (void *)(__u64)skb->data;
    struct ethhdr *eth;
    struct iphdr *iph;

    __u32 key    = 0; 
    __u64 *count = bpf_map_lookup_elem(&pkt_count, &key);
    __u64 *arp_count = bpf_map_lookup_elem(&arp_pkt_count, &key);

    if (count) { 
        __sync_fetch_and_add(count, 1); 
    }

    // やっぱskb自体がethernetフレームなんでは
    if (skb->protocol != bpf_htons(ETH_P_IP)) {
        if (skb->protocol == bpf_htons(ETH_P_ARP)) {
            if (arp_count) { 
                __sync_fetch_and_add(arp_count, 1);
            }
            // return TC_ACT_SHOT;
        }
        return TC_ACT_OK;
    }

    eth = data;

    if ((void *)(eth + 1) > data_end) {
        return TC_ACT_OK;
    }

    iph = (struct iphdr *)(eth + 1);
    if ((void *)(iph + 1) > data_end) {
        return TC_ACT_OK;
    }

    // if (bpf_ntohs(eth->h_proto) == ETH_P_ARP)
    //     if (arp_count) { 
    //         __sync_fetch_and_add(arp_count, 1); 
    //     }

    if (bpf_ntohs(eth->h_proto) != ETH_P_IP) {
        return TC_ACT_OK;
    }

    bpf_printk("Got IP packet!!: tot_len: %d, ttl: %d", bpf_ntohs(iph->tot_len), iph->ttl);

    // key = bpf_ntohl(iph->saddr);
        // nextHop = bpf_map_lookup_elem(&redirect_map_ipv4, &key);
    // if (nextHop == NULL) {
    //     bpf_trace_printk(notfound_str, sizeof(notfound_str), iph->saddr, bpf_ntohl(iph->saddr));
    //     return TC_ACT_OK;
    // }

    return TC_ACT_OK;
}