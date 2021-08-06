package main

import (
  "fmt"

  "github.com/sfreiberg/simplessh"
)

func ConnectSsh(method, tIP string, port, time int) {
  var client *simplessh.Client

  client, err := simplessh.ConnectWithPassword(C.SSH.IP, C.SSH.User, C.SSH.Password)
  if err != nil {
    fmt.Println(err)
  }

  cmd := "." + method + " " + tIP + " " + string(port) + " " + string(time)
  _, err = client.Exec(cmd)
  if err != nil {
    fmt.Println(err)
  }
  client.Close()
}

