package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc_project/userpb"
)

type server struct {
	userpb.UnimplementedUserServiceServer
	users map[string]*userpb.User
}

func (s *server) AddUser(ctx context.Context, req *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	id := fmt.Sprintf("user_%d", len(s.users)+1)
	user := &userpb.User{
		Id:    id,
		Name:  req.Name,
		Email: req.Email,
	}
	s.users[id] = user
	return &userpb.AddUserResponse{User: user}, nil
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &userpb.GetUserResponse{User: user}, nil
}

func (s *server) GetAllUsers(ctx context.Context, req *userpb.GetAllUsersRequest) (*userpb.GetAllUsersResponse, error) {
	var users []*userpb.User
	for _, user := range s.users {
		users = append(users, user)
	}
	return &userpb.GetAllUsersResponse{Users: users}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	_, exists := s.users[req.Id]
	if !exists {
		return &userpb.DeleteUserResponse{Success: false}, nil
	}
	delete(s.users, req.Id)
	return &userpb.DeleteUserResponse{Success: true}, nil
}

func main() {
	// Create a TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserService with the server
	userpb.RegisterUserServiceServer(grpcServer, &server{users: make(map[string]*userpb.User)})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
