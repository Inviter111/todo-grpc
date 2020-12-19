package main

import (
	"context"
	"log"
	"net"

	"todo/db"
	"todo/db/models"
	todo "todo/todo_proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type server struct {
	todo.UnimplementedTodoServiceServer
}

func (s *server) FindAll(ctx context.Context, in *todo.Empty) (*todo.Todos, error) {
	log.Printf("Received request")

	connection := db.Database

	var todos []models.Todo
	err := connection.Model(&todos).Select()

	return &todos, nil
}

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server")
	}

	db.Connect()
	defer db.Database.Close()

	s := grpc.NewServer()
	todo.RegisterTodoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve server")
	}
	log.Printf("Server is listening on port %s", port)
}
