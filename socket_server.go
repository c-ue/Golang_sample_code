package main

import (
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
	IP     string
	PORT   int
}

func Cmd_parser() (HOST_INFO, bool) {
	var ret HOST_INFO
	Bind_IP := flag.String("lh", "0.0.0.0", "help message for flagname")
	Bind_Port := flag.Int("lp", 8080, "help message for flagname")
	DEBUGACTIVED := flag.Bool("d", false, "help message for flagname.\nEx:" + os.Args[0] +
		" -lp=8080 -lh=0.0.0.0 -d")
	flag.Parse()
	ret.IP = *Bind_IP
	ret.PORT = *Bind_Port
	return ret, *DEBUGACTIVED
}

func main() {
	//main init
	var myhost, debugactived = Cmd_parser()
	if debugactived {
		out = os.Stdout
	}
	fmt.Fprintf(out, "DebugActived: %v\n", debugactived)
	fmt.Fprintf(out, "Argv Bind IP: %v\n", myhost.IP)
	fmt.Fprintf(out, "Argv Bind PORT: %v\n", myhost.PORT)
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Fprintf(out, "NumCPU: %d\n", runtime.NumCPU())
	fmt.Fprintf(out, "NumGoroutine: %d\n", runtime.NumGoroutine())

	//accept connection
	l, err := net.Listen("tcp", net.JoinHostPort(myhost.IP, strconv.Itoa(myhost.PORT)))
	if err != nil {
		fmt.Fprintln(out, "listen error:", err)
		return
	}
	fmt.Fprintf(out, "Server start listen %s\n", l.Addr().String())
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Fprintln(out, "accept error:", err)
			break
		}
		fmt.Fprintf(out, "Accept %s\n", c.RemoteAddr().String())
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	var data = make([]byte, 1024)
	fmt.Fprintln(out, "Start Receive:")
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Fprintln(out, "Receive error:", err)
			break
		}
		fmt.Fprintf(out, "[Receive %s] %s\n", c.RemoteAddr().String(), string(data[:n]))
		n, err = c.Write(data[:n])
		if err != nil {
			fmt.Fprintln(out, "Write error:", err)
			break
		}
	}
}
