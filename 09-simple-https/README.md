# 09-simple-https

Simple HTTPS/TLS example copied from https://gist.github.com/denji/12b3a568f092ab951456.

```
$> go run ./main.go
$> open https://localhost:5555
```

Another source of inspiration: https://github.com/golang/net/tree/master/http2/h2demo

**NOTE** that the server's private and public key files were generated like so:

```
## generate private key
$> openssl genrsa -out server.key 2048
## generate self-signed public key, based on the private key
$> openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```
