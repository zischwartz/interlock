package main

// interlock plugins
import (
	_ "github.com/zischwartz/interlock/plugins/example"
	_ "github.com/zischwartz/interlock/plugins/haproxy"
	_ "github.com/zischwartz/interlock/plugins/nginx"
	//_ "github.com/zischwartz/interlock/plugins/stats"
)
