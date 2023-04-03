package main

func main() {
	server := NewServer(":3030")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/createUser", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, Logging(), CheckAuth()))
	server.Listen()
}
