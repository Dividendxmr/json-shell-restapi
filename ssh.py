#!/usr/local/bin/python
import sys, json, paramiko

if len(sys.argv) < 4:
    print("Usage: python", sys.argv[0], " <method> <IP> <port> <time>")
# Needed arguments for command
method = sys.argv[1]
IP = sys.argv[2]
tPort = sys.argv[3]
time = sys.argv[4]

# Server specifications
cmd = method + " " + IP + " " + tPort + time

def connectSsh():
    ssh = paramiko.SSHClient()
    ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    ssh.connect(host, sPort, user, password)

    stdin, stdout, stderr = ssh.exec_command(cmd)
    ssh.close()

connectSsh()
