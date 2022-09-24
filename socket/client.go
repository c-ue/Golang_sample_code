package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strconv"
)

var out = ioutil.Discard

type HOST_INFO struct {
	IP   string
	PORT int
}

func Cmd_parser() (HOST_INFO, bool) {
	var ret HOST_INFO
	Target_Ip := flag.String("rh", "127.0.0.1", "help message for flagname")
	Target_Port := flag.Int("rp", 9090, "help message for flagname")
	DEBUGACTIVED := flag.Bool("d", false, "help message for flagname.\nEx:" + os.Args[0] +
		"-rh=127.0.0.1 -rp=9090 -d")
	flag.Parse()
	ret.IP = *Target_Ip
	ret.PORT = *Target_Port
	return ret, *DEBUGACTIVED
}

func main() {
	//main init
	var remote_host, debugactived = Cmd_parser()
	var n int
	var resp = make([]byte, 4096)
	if debugactived {
		out = os.Stdout
	}
	fmt.Fprintf(out, "DebugActived: %v\n", debugactived)
	fmt.Fprintf(out, "Argv Remote IP: %v\n", remote_host.IP)
	fmt.Fprintf(out, "Argv Remote PORT: %v\n", remote_host.PORT)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Fprintf(out, "NumCPU: %d\n", runtime.NumCPU())
	fmt.Fprintf(out, "NumGoroutine: %d\n", runtime.NumGoroutine())

	conn, err := net.Dial("tcp", net.JoinHostPort(remote_host.IP, strconv.Itoa(remote_host.PORT)))
	if err != nil {
		fmt.Fprintln(out, "connect error:", err)
		return
	}
	fmt.Fprintf(out, "Connect success %s-->%s\n", conn.LocalAddr().String(), conn.RemoteAddr().String())
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Start echo: ")
	for {
		text, _ := reader.ReadString('\n')
		n, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Fprintln(out, "Write error:", err)
			break
		}
		n, err = conn.Read(resp)
		if err != nil {
			fmt.Fprintln(out, "Receive error:", err)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", string(resp[:n]))
	}
}
