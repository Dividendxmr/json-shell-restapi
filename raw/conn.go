package main

import (
  "net"
  "log"
)

func ConnectServer(method, IP string, port, time int) {
  c, err := net.Dial("tcp", C.Server.IP)
  if err != nil {
    log.Fatal(err)
  }
  defer c.Close()

  _, err = c.Write([]byte(C.Server.Username))
  if err != nil {
    log.Fatal(err)
  }

  _, err = c.Write([]byte(C.Server.Password))
  if err != nil {
    log.Fatal(err)
  }

  cmd := "." + method + " " + IP + " " + string(port) + " " + string(time) + "\n"
  _, err = c.Write([]byte(cmd))
  if err != nil {
    log.Fatal(err)
  }
}

