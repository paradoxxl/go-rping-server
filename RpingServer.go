package main

import (
	"flag"
	"github.com/paradoxxl/go-rping-server/Components"
	"log"
	"net"
	"net/http"
	"strings"
)

var splitstr string = "//"

var Port = flag.Int("p", 8080, "port")
var BlockLocals = flag.Bool("l", true, "block locals")
var CustomPermit = flag.String("ca", "", "custom permit filter. Format:  ip/subnetmask//ip//... empty subnermask equals /32 (host) - higher predescence as allow locals")
var CustomBlock = flag.String("cb", "", "custom block filter. Format:  ip/subnetmask//ip//... empty subnermask equals /32 (host) - higher predescence as custom permit")

func main() {

	flag.Parse()

	if *CustomBlock != "" {
		//add customs to filter
		parseAndAdd(*CustomBlock, false)

	}
	if *CustomPermit != ""  {
		//add customs to filter
		parseAndAdd(*CustomPermit, true)
	}
	if *BlockLocals {
		//add local ipnets
		for _, v := range Components.LocalIps {
			Components.Filter.AddIPNetStringExt(v, false)
		}
	}
	Components.Filter.SetDefaultBehaviour(true)

	router := Components.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

func parseAndAdd(val string, permit bool) {
	//add customs to filter

	//seperate the string
	entries := strings.Split(val, "//")

	for _, v := range entries {
		if _, ipnet, err := net.ParseCIDR(v); err == nil {
			Components.Filter.AddIPNetExt(*ipnet, permit)
		} else if ip := net.ParseIP(v); ip != nil {
			//check whether v4 or v6 and create the /32 or /128 mask accordingly
			var mask net.IPMask
			if ipv4:=ip.To4();ipv4 == nil{
				mask = net.CIDRMask(8*net.IPv4len,8*net.IPv4len)
			}else {
				mask = net.CIDRMask(8*net.IPv6len,8*net.IPv6len)
			}
			Components.Filter.AddIPNetExt(net.IPNet{ip,mask}, permit)
		} else {
			log.Println("could not add the following item to the filter: ", v)
		}
	}

	q:=Components.Filter
	q.SetDefaultBehaviour(true)
}
