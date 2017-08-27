package client

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	GOBGPHOME         = "./.gobgp"
        GOBGP_STATUS    = "http://localhost:3000/api/status"
	GOBGP_JWKSTATUS	= "http://localhost:3000/api/jwkstatus"
	GOBGP_TOKEN     = "http://localhost:3000/api/token"
	GOBGP_COMMAND   = "http://localhost:3000/api/command"
	HASH_KEY         = "./.gobgp/.ghash_key"
	LCOMMANDFILE     = "./.gobgp/.last_command"
	GOBGPCOMMANDLOG = "/var/log/gobgp_client/gobgp_client.log"
	CONFIGINI        = "/go/go-honban/gobgpd.conf"
	GOBGP_DIR	= "/root/go/bin/"
)

const (
	ENBLACK = "\x1b[30m"
        ENRED = "\x1b[31m"
	ENGREEN = "\x1b[32m"
	ENYELLOW = "\x1b[33m"
	ENBLUE = "\x1b[34m"
	ENMAGENTA = "\x1b[35m"
	ENCYAN = "\x1b[36m"
	ENWHITE = "\x1b[37m"
	CONSOLE_CLEAR ="\x1b[0m"
)

//func str_paint (conlorCode string, str string) {
//	s := "\x1b["+colorCode+"m"
//        fmt.Printf("%s%s%s\n", s,str,CONSOLE_CLEAR)
//}


func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func cat(filename string) string {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(buff)
}

func dog(text string, filename string) {
	text_data := []byte(text)
	err := ioutil.WriteFile(filename, text_data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func addog(text string, filename string) {
	var writer *bufio.Writer
	text_data := []byte(text)

	write_file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0600)
	writer = bufio.NewWriter(write_file)
	writer.Write(text_data)
	writer.Flush()
	 if err != nil {
                fmt.Println(err)
        }
	defer write_file.Close()
}

func grep(str string, filepath string) (b bool) {
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			break
		}
		//              fmt.Printf("%4d行目: %s\n", i, sc.Text())
		if strings.Contains(sc.Text(), str) {
			return true
		}
	}
	return false
}

func Examples() {
	msg :=
		`
Examples(1):
 This Scriptes BGP Flow Spec command is for example...

     Do you want to do?(announce/withdraw): announce
     source IP(MUST)?: 10.0.0.0/24
     destination IP(MUST): 192.168.0.0/24
     protocols: tcp
     source port: 80
     dest port: 53
     Do you want to flowspec action?(accept/discard/rate-limit <ratelimit>): rate-limit 1000000

Results(1):
     Post Command: announce flow route source 10.0.0.0/24 destination 192.168.0.0/24 protocol [ tcp ] source-port [ =80 ] destination-port [ =53 ] rate-limit 1000000 

Examples(2):
 This Scriptes BGP Flow Spec command is for example...

     Do you want to do?(announce/withdraw): withdraw
     source IP(MUST)?: 20.0.0.1/32
     destination IP(MUST)?: 68.1.212.23/32
     protocols:
     source port:
     dest port:
     Do you want to flowspec action?(accept/discard/rate-limit <ratelimit>): discard 

Results(2):
     Post Command: withdraw flow route source 20.0.0.1/32 destination 68.1.212.23/32 discard
`
	fmt.Printf(msg)
	os.Exit(0)
}
