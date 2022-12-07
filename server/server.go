package main

import (
	"disastermanagement/database"
	"disastermanagement/pb"
	"log"

	"google.golang.org/grpc/reflection"

	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserErrorServiceServer
}

func init() {
	database.DbInit()
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":"+os.Getenv("SERVER_PORT"))
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()
	pb.RegisterUserErrorServiceServer(srv, &server{})
	reflection.Register(srv)
	log.Println("connected")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
