package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var marks = map[bool]string{true: "✔", false: "✖ "}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		domain := scanner.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exist])
		time.Sleep(1 * time.Second)
	}
}
func exists(domain string) (bool, error) {
	const whoisServer string = "whois.verisign-grs.com"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(strings.ToLower(text), "no match") {
			return false, nil
		}
	}
	return true, nil
}
