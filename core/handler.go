package core

import "fmt"

func (c *Client) HandleConnection() {
	MainLoop := true
	for MainLoop {
		userInput, err := c.ReceiveData()
		DisplayMsg(err)
		switch {
		case userInput == "stop" || userInput == "99":
			fmt.Println("Exiting")
			MainLoop = false
			break
		case userInput == "1":
			fmt.Println("Receive encrypted file")
			c.DownloadFile()
		default:
			fmt.Println("you entered: ", userInput)
		}
	}
}
