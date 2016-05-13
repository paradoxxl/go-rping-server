package Components

import "github.com/xuyu/ipfilter"

var localIps = string{
	"192.168.0.0/16",
	"172.16.0.0/12",
	"10.0.0.0/8"}

var Filter ipfilter.IPFilter
