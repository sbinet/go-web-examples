# 09-simple-https

Simple HTTPS/TLS example copied from https://gist.github.com/denji/12b3a568f092ab951456.

```
$> go run ./server.go
$> open https://localhost:5555
```

and, in another terminal:

```
$> go run ./client.go
hello world from HTTP/2.0
```

Another source of inspiration: https://github.com/golang/net/tree/master/http2/h2demo

**NOTE** that the server's private and public key files were generated like so:

```
## generate CA
$> openssl genrsa -out rootCA.key 2048
$> openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem

## generate server private key
$> openssl genrsa -out server.key 2048

## generate self-signed public key, based on the private key
$> openssl req -new -key server.key -out server.csr
$> echo subjectAltName = IP:127.0.0.1 > extfile.cnf
$> openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 3650 -extfile extfile.cnf
```
