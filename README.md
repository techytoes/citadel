### Citadel: Understanding Databases
---

The purpose of this repo is just to collect snippets from my experimentation with databases.

This makes use of **Unix domain socket** to make call between client and socket. Its helps in making IPC calls.
If we see clearly we are creating a `.sock` file which will be used to communicate between client and server.
Server would be listening to client using the `.sock` file. The server would directly be interacting with database and fetch the results.


### How to run?

Pretty straightforward. Run these 2 commands:

```
cd server && go run main.go
cd client && go run main.go
```