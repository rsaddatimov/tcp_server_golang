package main

import "github.com/rsaddatimov/tcp_server_golang"

func main() {
	server := tcp_server.New("localhost:5050")

	server.OnNewClient(func(c *tcp_server.Client) {
		// client connected
		c.Send("Hello")
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// message received
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection lost
	})

	server.Listen()
}
