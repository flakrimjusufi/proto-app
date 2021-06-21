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

2. Change Directory to the server directory and run the server:

~~~~
cd to _YOUR_LOCAL_PATH_/server and run **go run main.go**

**You should get a response from the server: server listening at [::]:50051** once the server is up and running.
~~~~

3. Change Directory to the client directory and run the client:

~~~~
cd to _YOUR_LOCAL_PATH_/client and run **go run main.go _your_name_ your_age**

You should have a response from client side, for example: **Sent Person data to the server: Name Flakrim, Age 27**

And then on server side: **Received: Flakrim 27** 
~~~~