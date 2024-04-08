package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"

	"github.com/redis/go-redis/v9"
)

// Explain - further (bytes, string builder and others..)
func removeNonAlphaNumeric(s []byte) []byte {
    n := 0
    for _, b := range s {
        if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') ||
            ('0' <= b && b <= '9') ||
            b == '-' || b == '.' {
            s[n] = b
            n++
        }
    }
    return s[:n]
}

func removePrefixIfExists(s string) string {

	prefixes := []string{"http://", "https://", "www."}

	after := s

	for _, prefix := range prefixes {
		if strings.Contains(s, prefix) {
			after = strings.TrimPrefix(s, prefix)
		}
	}

	return after
}

func redisConf(ipkey []net.IP) {
	client := redis.NewClient(&redis.Options{
        Addr:	  os.Getenv("REDIS_ADDRESS"),
        Password: os.Getenv("PASSWORD"), // no password set
        DB:		  0,  // use default DB
    })

	ctx := context.Background()

	data, _ := json.Marshal(ipkey)

	err := client.Set(ctx, "foo", data, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a domain name as an argument.")
		os.Exit(1)
	}

	domain := os.Args[1]

	trimmedDomain := removePrefixIfExists((domain))

	trimmedDomain = string(removeNonAlphaNumeric([]byte(trimmedDomain)))

	if len(trimmedDomain) > 25 {
		fmt.Println("Length reached max")
	}

	
	fmt.Println(trimmedDomain)

	ips, err := net.LookupIP(trimmedDomain)
	
	
	if err != nil {
		fmt.Println("DNS resolution failed:", err)
		// os.Exit(1)
	}

	ipStrings := make([]string, len(ips))
	
	for i, ip := range ips {
		fmt.Println(ip)
		ipStrings[i] = ip.String()
	}
	
	redisConf(ips)
	fmt.Printf("Type %s", reflect.TypeOf(ips))
}
