package client

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func IsIP(ip string) bool{
	if ipm, err := regexp.MatchString("^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$", ip); ipm {
		return true
		if err != nil {
			return false
		}
	}
	return false
}

func IsCIDR(cidr string) bool{
	if cidrm, err := regexp.MatchString("^([1-9]|[12][0-9]|3[0-2])$", cidr); cidrm {
		return true
		if err != nil {
			return false
		}
	}
	return false
}

func IsNUMBER(num string) bool{
	if numm, err := regexp.MatchString("^[0-9]+$", num); numm {
		return true
		if err != nil {
			return false
		}
	}
	return false
}

func numbers_checker(num_q string) (num_a string) {
	for {
		numb := bufio.NewReader(os.Stdin)
		fmt.Print(num_q)
		NUMB, _ := numb.ReadString('\n')
		NUMB = strings.Trim(NUMB, "\n")
		if NUMB == "" {
			num_a = NUMB
			break
		}
		if ! IsNUMBER(NUMB) {
		//	fmt.Println("NUMB OK")
		//	num_a = NUMB
		//	break
		//} else {
			fmt.Printf("\nSorry, %s is not number\n" ,NUMB)
			fmt.Println("Plerase, retype.\n")
			continue
		}
		num_a = NUMB
		break
	}
	return num_a
}

func address_checker(add_q string) (add_a string) {
	for {
		addr := bufio.NewReader(os.Stdin)
		fmt.Print(add_q)
		ADDR, _ := addr.ReadString('\n')
		ADDR = strings.Trim(ADDR, "\n")
		if strings.Contains(ADDR, "/") {
			addary := strings.SplitN(ADDR, "/", 2)
//			addary[1] = strings.Trim(addary[1], "\n")
			if ! IsIP(addary[0]) {
			//	fmt.Println("ADDR OK")
			//} else {
                        fmt.Printf("\nSorry, %s is bad address format.\n" ,addary[0])
                        fmt.Println("Plerase, retype.\n")
				continue
			}
			if ! IsCIDR(addary[1]) {
			//	fmt.Println("CIDR OK")
			//} else {
			fmt.Printf("\nSorry, %s is bad cidr format.\n" ,addary[1])
                        fmt.Println("Plerase, retype.\n")
				continue
			}
			add_a = ADDR
			break
		} else {
                        fmt.Printf("\nSorry, %s is not cidr.\n" ,ADDR)
                        fmt.Println("Plerase, retype.\n")
			continue
		}
	}
	add_a = strings.Trim(add_a, "\n")
	return add_a
}
