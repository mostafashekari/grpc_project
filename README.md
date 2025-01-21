Here’s a complete `README.md` file for your gRPC project:

---

# **gRPC Project**

This is a simple **gRPC-based User Management System** implemented in Go. It provides functionality to manage users, including adding, retrieving, listing, and deleting users. The project demonstrates the use of **Protocol Buffers (protobuf)** and **gRPC**.

---

## **Features**
- Add a user.
- Get user details by ID.
- List all users.
- Delete a user by ID.
- Implements a gRPC server and client.

---

## **Project Structure**
```
grpc_project/
├── client.go             # gRPC client implementation
├── server.go             # gRPC server implementation
├── user.proto            # Protobuf definition file
├── go.mod                # Go module file
├── go.sum                # Dependencies file
└── userpb/               # Generated Protobuf Go files
    ├── user.pb.go
    └── user_grpc.pb.go
```

---

## **Requirements**
- **Go**: 1.20 or later
- **Protocol Buffers (protoc)**: 3.0 or later
- **Protobuf Go Plugins**:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

---

## **Setup Instructions**

### **1. Clone the Repository**
```bash
git clone https://github.com/your-username/grpc_project.git
cd grpc_project
```

### **2. Install Dependencies**
```bash
go mod tidy
```

### **3. Generate Protobuf Files**
If `user.pb.go` and `user_grpc.pb.go` are missing, regenerate them:
```bash
protoc --go_out=. --go-grpc_out=. user.proto
```

### **4. Run the Server**
Start the gRPC server:
```bash
go run server.go
```

### **5. Run the Client**
In a separate terminal, run the gRPC client:
```bash
go run client.go
```

---

## **Usage Example**
1. **Start the Server**:
   ```bash
   go run server.go
   ```
   Expected output:
   ```
   Server is running on port 50051...
   ```

2. **Run the Client**:
   ```bash
   go run client.go
   ```
   Expected output:
   ```
   Added User: id:"user_1" name:"John Doe" email:"john@example.com"
   All Users: [id:"user_1" name:"John Doe" email:"john@example.com"]
   Fetched User: id:"user_1" name:"John Doe" email:"john@example.com"
   Deleted User Success: true
   All Users After Deletion: []
   ```

---

## **API Methods**

### **Protobuf Definition**
The gRPC API is defined in `user.proto`:
```proto
service UserService {
    rpc AddUser(AddUserRequest) returns (AddUserResponse);
    rpc GetUser(GetUserRequest) returns (GetUserResponse);
    rpc GetAllUsers(GetAllUsersRequest) returns (GetAllUsersResponse);
    rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse);
}

message User {
    string id = 1;
    string name = 2;
    string email = 3;
}

message AddUserRequest {
    string name = 1;
    string email = 2;
}

message AddUserResponse {
    User user = 1;
}

message GetUserRequest {
    string id = 1;
}

message GetUserResponse {
    User user = 1;
}

message GetAllUsersRequest {}

message GetAllUsersResponse {
    repeated User users = 1;
}

message DeleteUserRequest {
    string id = 1;
}

message DeleteUserResponse {
    bool success = 1;
}
```

---

## **Technologies Used**
- **Go**: Programming language.
- **gRPC**: Remote Procedure Call (RPC) framework.
- **Protocol Buffers**: Serialization format for defining structured data.
- **Go Modules**: Dependency management.

---

## **Contributing**
Contributions are welcome! If you find any issues or have suggestions, feel free to open an issue or submit a pull request.

---

## **License**
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## **Acknowledgments**
- [gRPC Documentation](https://grpc.io/docs/)
- [Protocol Buffers Documentation](https://protobuf.dev/)
