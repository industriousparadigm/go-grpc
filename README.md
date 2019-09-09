# Go gRPC

## client-API interaction written with Go

Stephane Maarek's Udemy course is the basis of this repo, where he shows how to write a server-client interaction in Go using the gRPC paradigm.

Calculator RPC's are implemented by me after seeing Maarek's implementation of the Greet RPC's.

Concepts involved are Google's protocol buffers and their advantages as an alternative to JSON/XML; the gRPC approach to designing API's as opposed to REST, integration with MongoDB and general concepts of writing with Golang.

### Types of gRPC covered

Unary (request -> response), Server Streaming (request -> stream of responses), Client Streaming (stream of requests -> response) and Bi-Directional Streaming (stream req -> stream res)
