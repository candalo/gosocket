package main

import (
	"bufio"
	"fmt"
	"gosocket/client"
	"os"
)

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Send messages to the server")
	fmt.Println("------------------------------")
	reader := bufio.NewReader(os.Stdin)
	client := new(client.Client)
	for {
		fmt.Print("->")
		msg, _ := reader.ReadString('\n')
		client.Connect()
		client.SendMessage(msg)
	}
}
