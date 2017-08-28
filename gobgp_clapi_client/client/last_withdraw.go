package client

import (
        "github.com/mattn/go-scan"
       "net/url"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Last_withdraw() {
	if !exists(LCOMMANDFILE) {
		fmt.Println("\nSorry,Not found \".last_command\".")
		fmt.Println("Now, Cannot use this option.\n")
		os.Exit(0)
	}
	if !grep("add", LCOMMANDFILE) {
		fmt.Println("\nSorry, Last command is NOT \"add\" flow.")
		os.Exit(0)
	}
	var js = strings.NewReader(cat(LCOMMANDFILE))
	var s string
	if err := scan.ScanJSON(js, "/command/", &s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lastAnnounce := s
	fmt.Println(lastAnnounce)
	lastAnnounce = strings.Replace(lastAnnounce , "/root/go/bin/" , "" ,1)
	fmt.Print("\n###################################################################\n")
	fmt.Println("\n Gobgp Flowspec client")
	fmt.Printf(" %sBe careful!!%s This Working is Last Announcing to Withdrawing!\n",ENRED,CONSOLE_CLEAR)
	fmt.Printf(" Last Announce: %s%s%s\n",ENYELLOW,lastAnnounce,CONSOLE_CLEAR)
	fmt.Print("\n###################################################################\n")

	changeCommand := strings.Replace(cat(LCOMMANDFILE), "add", "del", 1)
	noJsonchangeCommand := strings.Replace(lastAnnounce, "add", "del", 1)
	currentHash := check_hash()

	fmt.Print("\n######################################\n")
	fmt.Print(" POST The Last announce withdrawing!\n")
	fmt.Print("######################################\n")

	fmt.Println("\n######################################################################\n")
	fmt.Printf(" Current Hash Code: %s%s%s\n",ENBLUE,currentHash,CONSOLE_CLEAR)
	fmt.Printf(" Post Command: %s%s%s\n", ENGREEN,noJsonchangeCommand,CONSOLE_CLEAR)
	fmt.Print("\n######################################################################\n")

	for {
		predoit := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to POST this command??(y/n): ")
		doit, _ := predoit.ReadString('\n')
		doit = strings.Trim(doit, "\n")
		switch doit {
		case "y":
			if !exists(LCOMMANDFILE) {
				os.MkdirAll(GOBGPHOME, 0600)
				os.Create(LCOMMANDFILE)
			}
			dog(changeCommand, LCOMMANDFILE)
			values := url.Values{}
			curl_post_command(values, currentHash)
			t := time.Now()
			var t_tos string
			t_tos = t.String()
			timeAndLog := "[" + t_tos + "] GoBGP POST Command: " + noJsonchangeCommand + "\n"
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
