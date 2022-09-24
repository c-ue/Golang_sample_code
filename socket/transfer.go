package main

import (
	"flag"
	"fmt"
	"io"
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

func Cmd_parser() (HOST_INFO, HOST_INFO, bool) {
	var myhost, remote_host HOST_INFO
	Bind_IP := flag.String("lh", "0.0.0.0", "help message for flagname")
	Bind_Port := flag.Int("lp", 8080, "help message for flagname")
	Target_Ip := flag.String("rh", "127.0.0.1", "help message for flagname")
	Target_Port := flag.Int("rp", 9090, "help message for flagname")
	DEBUGACTIVED := flag.Bool("d", false, "help message for flagname.\nEx:" + os.Args[0] +
		" -lp=8080 -lh=0.0.0.0 -rh=127.0.0.1 -rp=9090 -d")
	flag.Parse()
	myhost.IP = *Bind_IP
	myhost.PORT = *Bind_Port
	remote_host.IP = *Target_Ip
	remote_host.PORT = *Target_Port
	return myhost, remote_host, *DEBUGACTIVED
}

func main() {
	//main init
	var myhost, remote_host, debugactived = Cmd_parser()
	if debugactived {
		out = os.Stdout
	}
	fmt.Fprintf(out, "DebugActived: %v\n", debugactived)
	fmt.Fprintf(out, "Argv Bind IP: %v\n", myhost.IP)
	fmt.Fprintf(out, "Argv Bind Port: %v\n", myhost.PORT)
	fmt.Fprintf(out, "Argv Remote IP: %v\n", remote_host.IP)
	fmt.Fprintf(out, "Argv Remote PORT: %v\n", remote_host.PORT)
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
		go handleConn(c, remote_host)
	}
}

func handleConn(c net.Conn, remote_host HOST_INFO) {
	defer c.Close()
	conn, err := net.Dial("tcp", net.JoinHostPort(remote_host.IP, strconv.Itoa(remote_host.PORT)))
	if err != nil {
		fmt.Fprintln(out, "connect error:", err)
		return
	}
	fmt.Fprintf(out, "Connect success %s<-->%s\n", c.LocalAddr().String(), conn.RemoteAddr().String())
	defer conn.Close()
	go io.Copy(conn, c)
	io.Copy(c, conn)
}
