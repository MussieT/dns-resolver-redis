package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"reflect"

	"github.com/dns-resolver-redis/connect"
	"github.com/dns-resolver-redis/parser"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	if len(os.Args) < 2 {
		fmt.Println("Please provide a domain name as an argument.")
		os.Exit(1)
	}

	domain := os.Args[1]

	trimmedDomain := parser.SanitizeQuery(domain)
	
	ips, err := net.LookupIP(trimmedDomain)
	
	if err != nil {
		fmt.Println("DNS resolution failed:", err)
		os.Exit(1)
	}

	ipStrings := make([]string, len(ips))
	
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}

	connect.SetValue(domain, ips)
	
	// redisConf(ips)
	fmt.Printf("Type %s", reflect.TypeOf(ips))
}
