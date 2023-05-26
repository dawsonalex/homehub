package main

import (
	"context"
	"fmt"
	"github.com/dawsonalex/homehub/cmd/todo/schema"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
}

func (s *Server) GetTodo(ctx context.Context, request *schema.GetTodoRequest) (*schema.GetTodoResponse, error) {
	return &schema.GetTodoResponse{
		Id:          request.Id,
		Description: "DEFAULT TODO RESPONSE",
		CreatedOn:   timestamppb.Now(),
		ModifiedOn:  timestamppb.Now(),
	}, nil
}

func (s *Server) CreateTodo(ctx context.Context, request *schema.CreateTodoRequest) (*schema.GetTodoResponse, error) {
	//TODO implement storing todos
	return &schema.GetTodoResponse{
		Id:          0,
		Description: "DEFAULT TODO RESPONSE",
		CreatedOn:   timestamppb.Now(),
		ModifiedOn:  timestamppb.Now(),
	}, nil
}

func main() {
	fmt.Println("Hello, World!")
}
