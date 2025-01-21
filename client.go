package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc_project/userpb"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client for the UserService
	client := userpb.NewUserServiceClient(conn)

	// Add a user
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addUserResp, err := client.AddUser(ctx, &userpb.AddUserRequest{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("Added User: %v", addUserResp.User)

	// Get all users
	getAllUsersResp, err := client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
	if err != nil {
		log.Fatalf("could not get all users: %v", err)
	}
	log.Printf("All Users: %v", getAllUsersResp.Users)

	// Get a user by ID
	getUserResp, err := client.GetUser(ctx, &userpb.GetUserRequest{Id: addUserResp.User.Id})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("Fetched User: %v", getUserResp.User)

	// Delete the user
	deleteUserResp, err := client.DeleteUser(ctx, &userpb.DeleteUserRequest{Id: addUserResp.User.Id})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	log.Printf("Deleted User Success: %v", deleteUserResp.Success)

	// Get all users after deletion
	getAllUsersResp, err = client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
	if err != nil {
		log.Fatalf("could not get all users: %v", err)
	}
	log.Printf("All Users After Deletion: %v", getAllUsersResp.Users)
}
