SmartFridge
===========

By Lv-276.Golang
----------------


### How to start
Run to install glide and gometalinter

    make install-helpers

Run toinstall application dependencies

    make dependencies
    
### Code quality check
    make code-quality


Agent
-----

Agent communicates with server via websocket connection.
After the connection is established agent sends its ID and waits for configuration message from the server. This message sets userID (the owner of the container), the ID of product (content of agent) and the frequency of sending agent data to server.

To run agent with default config file
```go run cmd/agent/main.go```

To run agent and specify other destination
```go run cmd/agent/main.go -config=[path]```