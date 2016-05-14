package Components

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tatsushid/go-fastping"
	"net"
	"net/http"
	"strconv"
	"time"
)

func Commands(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Available commands\n")

	//fmt.Fprintf(w,routes[0].Name)
	//for _,route := range routes{
	//	fmt.Fprintf(w, route.Pattern)
	//}

}

func SimplePing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ipaddr := vars["ipaddr"]

	times, err := strconv.Atoi(vars["times"])
	if err != nil {
		times = 1
	}

	success := false

	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ipaddr)
	if err != nil {
		fmt.Fprintf(w, "invalid ip or name")
		return
	}

	if !Filter.FilterIP(ra.IP){
		fmt.Fprintf(w, "this ip is filtered by the server administrator")
		return
	}

	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Fprintf(w, "IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		success = true
	}
	p.OnIdle = func() {
		if !success {
			fmt.Fprintf(w, "%s not reachable", ra.String())
		}
	}

	for i := 0; i < times; i++ {
		err = p.Run()
		if err != nil {
			fmt.Fprintf(w, "Something went wrong")

		}
	}
}


