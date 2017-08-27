package main

import (
	"github.com/cloudbuy/go-pkg-optarg"
	"os"
	"./client"
)

func main() {
	optarg.Add("h", "help", "Displays Commnad Examples.", false)
	optarg.Add("w", "withdraw", "Withdrawing the last announce prefix.", false)

	if len(os.Args) == 1 {
		client.Announce()
	} else {
		for opt := range optarg.Parse() {
			switch opt.ShortName {
			case "h":
				client.Examples()
			case "w":
				client.Last_withdraw()
			}
		}
	}
}
