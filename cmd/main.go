package main

func main() {
	server := NewAPIServer(":9999")
	server.Run()
}
