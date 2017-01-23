package client

import (
	"log"
	"net"
)

type Client struct {
	conn net.Conn
}

func (c *Client) Connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn
}

func (c *Client) SendMessage(msg string) {
	defer c.conn.Close()
	c.conn.Write([]byte(msg))
}
