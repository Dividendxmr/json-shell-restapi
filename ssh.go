package main

import (
  "os/exec"
)

func ConnectSsh(method, target string, port, time int) {
  cmd := exec.Command("python", "ssh.py", method, target, string(port), string(time))
  _, err := cmd.Output()
  if err != nil {
    return
  }
}
