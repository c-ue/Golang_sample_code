package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

var out = ioutil.Discard

type HOST_INFO struct {
	Bind_IP     string
	Bind_Port   int
	Target_Ip   string
	Target_Port int
}

func Cmd_parser() (HOST_INFO, bool) {
	var ret HOST_INFO
	Bind_IP := flag.String("lh", "0.0.0.0", "help message for flagname")
	Bind_Port := flag.Int("lp", 8080, "help message for flagname")
	Target_Ip := flag.String("rh", "127.0.0.1", "help message for flagname")
	Target_Port := flag.Int("rp", 9090, "help message for flagname")
	DEBUGACTIVED := flag.Bool("d", false, "help message for flagname.\nEx:" + os.Args[0] +
		" -lp=8080 -lh=0.0.0.0 -rh=127.0.0.1 -rp=9090 -d")
	flag.Parse()
	ret.Bind_IP = *Bind_IP
	ret.Bind_Port = *Bind_Port
	ret.Target_Ip = *Target_Ip
	ret.Target_Port = *Target_Port
	return ret, *DEBUGACTIVED
}

func main() {
	//main init
	var host, debugactived= Cmd_parser()
	if debugactived {
		out = os.Stdout
	}
	fmt.Fprintf(out, "DebugActived: %v\n", debugactived)
	fmt.Fprintf(out, "Argv Bind IP: %v\n", host.Bind_IP)
	fmt.Fprintf(out, "Argv Bind Port: %v\n", host.Bind_Port)
	fmt.Fprintf(out, "Argv Remote IP: %v\n", host.Target_Ip)
	fmt.Fprintf(out, "Argv Remote PORT: %v\n", host.Target_Port)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Fprintf(out, "NumCPU: %d\n", runtime.NumCPU())
	fmt.Fprintf(out, "NumGoroutine: %d\n", runtime.NumGoroutine())
}
