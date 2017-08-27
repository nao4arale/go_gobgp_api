package client

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	//	"io/ioutil"
	"net/url"
	"os"
	"syscall"
	"strings"
)

func check_hash() string {
	fmt.Println("\n##########################")
	fmt.Println("    check the hash key")
	fmt.Println("##########################\n")
	prg()
	if ! exists(HASH_KEY) {
		os.MkdirAll(GOBGPHOME, 0600)
		os.Create(HASH_KEY)
	}
	hash := cat(HASH_KEY)

//	HASHCHECK := curl_check(hash, "unused")
        hash_tmp := "Bearer " + hash
        HASHCHECK := curl_check_jwk("Authorization", hash_tmp)
	if HASHCHECK {
         fmt.Println("\nOK,Current HASH key is not still changed.")
         fmt.Println("Go to Next Process.")

		return hash
	} else {

	 fmt.Print("\nSorry,Current HASH key is changed.\n")
	 fmt.Print("Please,GET HASH.\n")
	 fmt.Print("\n#############################\n")
	 fmt.Println(` Let's Sending GET Token.`)
	 fmt.Print(" Please type user, and pass.\n")
	 fmt.Println("#############################\n")

		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("hash_user: ")
			HASHUSER, _ := reader.ReadString('\n')
			HASHUSER = strings.Trim(HASHUSER, "\n")
			fmt.Print("hash_pass: ")
			hashPass, _ := terminal.ReadPassword(int(syscall.Stdin))
			HASHPASS := string(hashPass)
			HASHPASS = strings.Trim(HASHPASS, "\n")
			HASHCHECK2 := curl_check(HASHUSER, HASHPASS)
			if HASHCHECK2 {
				values := url.Values{}
				hash = curl_get(values,HASHUSER, HASHPASS)
				dog(hash, HASH_KEY)
				break
			} else {
				fmt.Println("\nSorry,Not Complete the GET HASH.")
				fmt.Println("Please,Retry again.\n")
				continue
			}
		}
		//values := url.Values{}
		//hash = curl_get(values)
		//dog(hash, HASH_KEY)
		//	 hash_data := []byte(hash)
		//	 ioutil.WriteFile(hash_key, hash_data, os.ModePerm)

		fmt.Println("\n\n###########################")
		fmt.Println("     Success GET HASH!")
		fmt.Println("###########################")

		return hash

	}
}
