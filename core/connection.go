package core

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	RemoteIP string
	RemotePort string
	Conn net.Conn
}

func (c *Client)ConnectServer() (err error) {

	c.Conn, err = net.Dial("tcp", c.RemoteIP+":"+c.RemotePort)
	if err != nil {
		fmt.Println("[-] Error: couldn't connect: " + c.RemoteIP)
	} else {
		fmt.Println("[+] Connection established with :" + c.Conn.RemoteAddr().String())

	}
	return err
}

func (c *Client) ReceiveData()(data string, err error){
	data, err = bufio.NewReader(c.Conn).ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
	return data, err
}

func (c *Client) SendData(data string)(err error){
	b, err := c.Conn.Write([]byte(data+"\n"))
	fmt.Println(b, "bytes written")
	return err
}