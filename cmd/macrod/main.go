package main

import (
	"fmt"
	"github.com/dawsonalex/homehub/cmd/macrod/repo"
	"github.com/dawsonalex/homehub/cmd/macrod/schema"
	"github.com/dawsonalex/homehub/pkg/db/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	db, err := postgres.OpenConnection(postgres.Config{
		User:   "postgres",
		Dbname: "macrod",
	})
	if err != nil {
		panic(err)
	}

	schema.RegisterMacrodServer(s, &server{
		repo: repo.NewPostgres(*db),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
