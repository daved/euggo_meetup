//START1 OMIT
package main

import (
	"flag" //OMIT
	"fmt"  //OMIT
	"net"  //OMIT
	"os"   //OMIT
	// ...
	"github.com/daved/grpcbasic0/idl"
	"google.golang.org/grpc"
)

//END1 OMIT

func listener(port string) (*net.TCPListener, error) {
	a, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve addr: %v\n", err)
	}

	l, err := net.ListenTCP("tcp", a)
	if err != nil {
		return nil, fmt.Errorf("cannot listen: %v\n", err)
	}

	return l, nil
}

func main() {
	var port string
	flag.StringVar(&port, "rcp", ":3323", "rcp port (default: ':3323')")
	flag.Parse()

	//START2 OMIT
	l, err := listener(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot get listener: %v\n", err)
		os.Exit(1)
	}

	srvr, svc := grpc.NewServer(), NewUserService()
	idl.RegisterUserServiceServer(srvr, svc)
	//OMIT
	if err = srvr.Serve(l); err != nil {
		fmt.Fprintf(os.Stderr, "rpc server error: %v\n", err)
		os.Exit(1)
	}
	//END2 OMIT
}
