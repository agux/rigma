package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/agux/rigma/rsec"
	"github.com/agux/rigma/tes"
	"github.com/chrislusf/gleam/gio"
)

func main() {
	gio.Init() // If the command line invokes the mapper or reducer, execute it and exit.

	svr := rpc.NewServer()
	//Register RPC services
	svr.Register(new(rsec.IndcScorer))
	svr.Register(new(rsec.DataSync))
	svr.Register(new(tes.GTest))

	l, e := net.Listen("tcp", ":45321")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	log.Printf("PID:%+v PPID:%+v listening on port: %d", os.Getpid(), os.Getppid(), 45321)

	// This statement links rpc server to the socket, and allows rpc server to accept
	// rpc request coming from that socket.
	svr.Accept(l)
}
