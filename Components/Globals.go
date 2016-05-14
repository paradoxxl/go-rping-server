package Components

import "github.com/paradoxxl/go-ipfilter"

var LocalIps = []string{
	"192.168.0.0/16",
	"172.16.0.0/12",
	"10.0.0.0/8",
	"224.0.0.0/4",
	"fe80::/10",
	"ff00::/8",
	"fc00::/7"}

var Filter ipfilter.IPFilter
