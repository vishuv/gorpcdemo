package main

import "fmt"
import "net"
import "net/rpc"

import "time"
import "log"

type TellTime struct{}

type Args struct {
	name string
}

type Reply string

func (t *TellTime) TellTimeTo(args *Args, reply *Reply) error {
	*reply = new Reply(time.Now().Local().String())
	return nil
}

func main() {
	var tt = new(TellTime)
	rpcserver := rpc.NewServer()
	rpcserver.Register(tt)
	listener, err := net.Listen("unix", "jellow")
	if err != nil {
		log.Fatal("Error while listening to unix socket")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting a connection")
		}
		go rpcserver.ServeConn(conn)
	}
	fmt.Printf("Hello world!")
}
