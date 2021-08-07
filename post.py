import requests, sys

URL = "http://127.0.0.1:8081/api"

json = {
        "Key": "test",
        "Target": "127.0.0.1",
        "Method": "udpplain",
        "Port": 1337,
        "Time": 300
        }
r = requests.post(URL, json=json)

print(r.status_code)
print(r.headers)
print(r.text)
