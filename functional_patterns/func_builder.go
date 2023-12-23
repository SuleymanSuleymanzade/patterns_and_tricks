package main

import (
	"fmt"
)

type Server struct {
	Opts *Opt
}

type OptFunc func(*Opt)

type Opt struct {
	maxConn int
	id      string
	tls     bool
}

func defautOpt() *Opt {
	return &Opt{
		maxConn: 23,
		id:      "Default",
		tls:     false,
	}
}

func withTLS(opts *Opt) {
	opts.tls = true
}

func setId(id string) OptFunc {
	return func(opt *Opt) {
		opt.id = id
	}
}

func newServer(opt ...OptFunc) *Server {
	o := defautOpt()
	for _, fn := range opt {
		fn(o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	server := newServer(withTLS, setId("MyId"))

	fmt.Printf("%+v\n", server)
	fmt.Println(server.Opts.id)
	fmt.Println(server.Opts.maxConn)
	fmt.Println(server.Opts.tls)

}
