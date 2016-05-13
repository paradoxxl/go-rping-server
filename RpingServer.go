package main

import (
	"flag"
	"strings"
)
var Port = flag.Int("p",8080,"port");
var BlockLocals = flag.Bool("l",true,"block locals")
var CustomPermit = flag.String("ca","","custom permit filter. Format:  ip/subnetmask//ip///... empty subnermask equals /32 (host) - higher predescence as allow locals")
var CustomBlock = flag.String("ca","","custom block filter. Format:  ip/subnetmask//ip///... empty subnermask equals /32 (host) - higher predescence as custom permit")

func main() {

	flag.Parse()

	if CustomPermit !=""{
		//add customs to filter
	}
	if CustomPermit !=""{
		//add customs to filter
	}
	if BlockLocals !=""{
		//add locals
	}

	//router := NewRouter()

	//log.Fatal(http.ListenAndServe(":8080", router))

}
