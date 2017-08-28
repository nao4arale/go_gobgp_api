package client

import (
	"bufio"
	"fmt"
//	"github.com/mattn/go-pipeline"
//	"golang.org/x/crypto/ssh/terminal"
//	"log"
	"net/url"
	"os"
//	"syscall"
	"strings"
	"time"
)

func Announce() {
	fmt.Println("\n#########################")
	fmt.Println("  Gobgp Flowspec client")
	fmt.Println("#########################\n")

//	cli_vrf     := check_vrf("vrf(MUST): ")
        cli_action := check_action("Do you want to do?(add/del): ")
        cli_dest_ip := address_checker("destination_ip(MUST): ")
	cli_source_ip := address_checker("source_ip(MUST): ")
        cli_protocols := check_protocols("protocols(tcp/udp/any): ")
        cli_dest_port := numbers_checker("destion_port: ")
	cli_source_port := numbers_checker("source_port: ")
	cli_then := check_then("Do you want to then?(accept/discard/rate-limit <ratelimit>): ")
//	vrf	:= "vrf " + cli_vrf
	action := cli_action
        dest_ip := " destination " + cli_dest_ip
	source_ip := " source " + cli_source_ip +" "
	protocols := "protocol " + cli_protocols + " "
	 if cli_protocols == "" {
	 protocols = ""
	}
	dest_port := "destination-port =='" + cli_dest_port + "' "
         if cli_dest_port == "" {
         dest_port = ""
	}
	source_port := "source-port =='" + cli_source_port + "' "
	 if cli_source_port == "" {
	 source_port = ""
	}
	then := "then " + cli_then

        precmd := "gobgp global rib -a ipv4-flowspec" + action + " match" + dest_ip + source_ip + protocols + dest_port + source_port + then

	currentHash := check_hash()

	fmt.Println("\n######################################################################\n")
	fmt.Printf("    Current Hash Code: %s%s%s\n" ,ENGREEN,currentHash,CONSOLE_CLEAR)
	fmt.Printf("	 Post Command: %s%s%s\n\n" ,ENBLUE, precmd,CONSOLE_CLEAR)
	fmt.Println("######################################################################\n")

	for {
		predoit := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to POST this command??(y/n): ")
		doit, err := predoit.ReadString('\n')
		fatal(err)
		doit = strings.Trim(doit, "\n")
		switch doit {
		case "y":
			cmd := `{"command":` + `"` + GOBGP_DIR + precmd + `"}`
			if !exists(LCOMMANDFILE) {
		                os.MkdirAll(GOBGPHOME, 0600)
                                os.Create(LCOMMANDFILE)
			}
			dog(cmd, LCOMMANDFILE)
			values := url.Values{}
			curl_post_command(values, currentHash)
			t := time.Now()
			var t_tos string
			t_tos = t.String()
			timeAndLog := "["+t_tos+"] GoBGP POST Command: "+precmd+"\n"
			addog(timeAndLog, GOBGPCOMMANDLOG)

			fmt.Println("\n####################")
			fmt.Printf("  \x1b[36;1mWorking is Done.\x1b[0m\n" )
			fmt.Println("####################\n")

			os.Exit(0)
		case "n":
			fmt.Print("\nOkay,See you the next working.\n")
			os.Exit(0)
		default:
			fmt.Println("Sorry,cannot understand.")
			fmt.Println("Please,type the \"y\" or \"n\"")
			continue
			}
		}
	}

