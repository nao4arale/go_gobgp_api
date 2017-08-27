package client

import (
	"bufio"
	"fmt"
	//	"github.com/mattn/go-pipeline"
	//	"golang.org/x/crypto/ssh/terminal"
	//	"log"
	"os"
	//	"syscall"
	"strings"
	"regexp"
)

func check_action(act_q string) string {
	for {
		precli_action := bufio.NewReader(os.Stdin)
		fmt.Print(act_q)
		cli_action, _ := precli_action.ReadString('\n')
		cli_action = strings.Trim(cli_action, "\n")
		switch cli_action {
		case "add":
			break
		case "del":
			break
		default:
			fmt.Println("\nSorry,cannot understand.")
			fmt.Println("Please,type the \"add\" or \"del\"\n")
			continue
		}
		cli_action =" "+cli_action
		return cli_action
	}
}

func check_protocols(proto_q string)  string {
	for {
		precli_protocols := bufio.NewReader(os.Stdin)
		fmt.Print(proto_q)
		cli_protocols, _ := precli_protocols.ReadString('\n')
		cli_protocols = strings.Trim(cli_protocols, "\n")
		switch cli_protocols {
		case "tcp":
			break
		case "udp":
			break
//		case "unknown":
//			break
		case "any":
			cli_protocols = ""
			break
		case "":

			fmt.Println("\nUmm...Sorry,protocols=\"\" is not support\n")
			fmt.Println("If you Specifies Protocols \"ALL\", Please type \"any\"\n")
			continue
		default:
			fmt.Println("\nSorry,cannot understand.")
			fmt.Println("Please,type the \"tcp\" , \"udp\" , or \"any\"\n")
			continue
		}
		return cli_protocols
	}
}

func check_then(then_q string) string {
	for {
		precli_then := bufio.NewReader(os.Stdin)
		fmt.Print(then_q)
		cli_then, err := precli_then.ReadString('\n')
			if err != nil {
			fmt.Println(err)
				os.Exit(1)	}
		cli_then = strings.Trim(cli_then, "\n")
		 if strings.Contains(cli_then, " "){
		  rateary := strings.SplitN(cli_then, " ", 2)
		  r := regexp.MustCompile(`^rate-limit$`)
		   if ! r.MatchString(rateary[0]) {
//		   if ! r.MatchString(cli_then) {
	                  continue
		   }
		   if ! IsNUMBER(rateary[1]) {
			fmt.Println("\nSorry,cannot understand.")
                        fmt.Println("Please,type the \"accept\" or \"discard\" or \"rate-limit <ratelimit>\"\n")
		   continue
		   } else {
			return cli_then
		  }
		}
		switch cli_then {
		case "accept":
			return cli_then
		case "discard":
			return cli_then
//		case r.MatchString(cli_then):
//			return cli_then
		default:
//			if r.MatchString(cli_then) {
//			return cli_then
//			}
			fmt.Println("\nSorry,cannot understand.")
			fmt.Println("Please,type the \"accept\" or \"discard\" or \"rate-limit <ratelimit>\"\n")
			continue
		}
//		cli_then=" "+cli_then
		return cli_then
	}
}

/*
func check_neighbor(nei_q string) string {
	var neigh string
	for {
		precli_nei := bufio.NewReader(os.Stdin)
		fmt.Print(nei_q)
		cli_nei, _ := precli_nei.ReadString('\n')
		cli_nei = strings.Trim(cli_nei, "\n")
		if ! IsIP(cli_nei) {
//			break
//		} else {
			fmt.Printf("\nSorry,It's %s is bad address format.\n", cli_nei)
			fmt.Println("Please, check the address\n")
			continue
		}
		if ! grep(cli_nei, CONFIGINI) {
//			break
//		} else {
			fmt.Printf("\nSorry,Cannot find \"neighbor %s\" in config.\n", cli_nei)
			fmt.Println("Please,check the address.\n")
			continue
		}
		neigh = "neighbor " + cli_nei
	Q:
		prenei_q := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to Announce Only One neibhgor?(y/n): ")
		nei_q, _ := prenei_q.ReadString('\n')
		nei_q = strings.Trim(nei_q, "\n")
		switch nei_q {
		case "n":
			if strings.Contains(neigh, "neighbor") {
				neigh += " , "
			}
			continue
		case "y":
			break
		default:
			fmt.Println("\nSorry,cannot understand.")
			fmt.Println("Please,type the \"announce\" or \"withdraw\"\n")
			goto Q
		}
		break
	}
	return neigh
}
*/
