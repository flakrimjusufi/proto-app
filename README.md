# proto-app

This is just a simple gRPC app build with Go and ProtoBuff.

To try it, please make sure you have Go installed:

~~~~
Go [https://golang.org/], any one of the three latest major releases of Go.
For installation instructions, see Go's getting started guide: https://golang.org/doc/install
~~~~

## How to Run it?

1. Clone the repo in your local environment:

~~~~
git clone https://github.com/flakrimjusufi/proto-app.git
~~~~

2. Change directory to the project root and install dependencies with go mod commands:
~~~~
- go mod init server/main.go 
- go mod tidy 

**You should see all dependencies installed in a go.mod and go.sum file 
~~~~

3. Run the server first:
~~~~
go run server/main.go

**You should get a response from the server: server listening at [::]:50051** once the server is up and running.
~~~~

4. Run the client afterwards:

~~~~
go run client/main.go _your_name_ your_age

**You should have a response from client side, for example: **Sent Person data to the server: Name Flakrim, Age 27**

**And then on server side: **Received: Flakrim 27** 
~~~~