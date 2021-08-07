package main

import (
  "net"
  "log"
  "time"
)

func ConnectServer(method, IP string, port, time int) {
  c, err := net.Dial("tcp", C.Server.IP)
  if err != nil {
    log.Fatal(err)
  }

  _, err = c.Write([]byte("\n\n"))
  if err != nil {
    log.Fatal(err)
  }
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

  time.Sleep(120 * time.Second)
  c.Close()
}

