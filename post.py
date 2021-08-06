import requests

URL = "http://127.0.0.1:8081/api"

r = requests.post(URL)

print(r.status_code)
print(r.headers)
print(r.text)
