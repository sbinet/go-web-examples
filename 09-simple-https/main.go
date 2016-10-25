package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
)

func main() {
	flag.Parse()

	cert, err := tls.X509KeyPair(serverCert, serverKey)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	http.HandleFunc("/", helloWorld)
	srv := http.Server{
		Addr:      *addrFlag,
		TLSConfig: cfg,
	}
	err = srv.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from %v\n", r.Proto)
}

var (
	// generated with:
	//  $> openssl genrsa -out rootCA.key 2048
	//  $> openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem
	rootCA = []byte(`-----BEGIN CERTIFICATE-----
MIIDXTCCAkWgAwIBAgIJAIee3w8TvMpZMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTYxMDI1MTE1MzU4WhcNMTkwODE1MTE1MzU4WjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAlNT8iFvp2HAJmH7HPJlykF1OWNI5xE3+kvQuVzH6s9TwYuqP6sQC5UtO
x+rQKkk/6b/53MLrU6dw9/9MUPYcGljxFMkDIvb4cyxz/WegHggmJ+AD18mDiQe5
Qnln6Gf2elEivdfPP72oWi2LJBycDM1YWVC7za0o6do2QGiSP99uWAa3HQTZrN2y
MYHGRKx8rh1WIvnc1v32VbkdFf13jNFwiLRVReuFFWClcyTkzN3Yx0aMZf4IWj3/
9h1XONJXATksXUiKSI5IM3Au6Ljk8DLx3KhDDUSd6uUTcugmun8LYTLCjRPwRjDj
ywYBCVwH1JSwOZrFjM77suV4aprSmwIDAQABo1AwTjAdBgNVHQ4EFgQUlyVsFHel
kdnwC80RWtLHbfIINpkwHwYDVR0jBBgwFoAUlyVsFHelkdnwC80RWtLHbfIINpkw
DAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAGJd142ObC5XvAQyx64aV
m4WqzqvBRI+e7BWOTBqxP7Hpz0bWxH9tqsQONeYsYgkiXrpTsYf1B+/+bS28AEx5
awlRuJ0u1hCrxwQAJHKjhs1SMU8HFKA6eLMZ7Ko3xTAcMiuoMx2AJkJUxiugw2gK
tcBMBtJy9II/yu3IMdsVK7v3W6RknIwvdYl49Uz6nOWPbxgsEqq1D/wO+GMWF5lm
4VeUxdr9R3Pg8QAOQ37LIZhRYInQ+6O29q1RpGslhuLdL7nou1t80EQSvyQEACj8
hH130jK5lKGLspDfV9FfCZ3Rj275LOA3zdWjXenYDbEwWUHLB/3TXUSRkgV6Ga0i
rw==
-----END CERTIFICATE-----
`)
	// generated with:
	//  $> openssl genrsa -out server.key 2048
	serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEAzsADcua6qCyY1AMwi0JRLw1+bXFv3UX8KOzWUtOGXpjgeq/3
3UTFZBK1M9vEQdN7fO+m5+AOLmfxmKQzLdjdO/hCwQZt4i33jFb6fPiZbPx082Ga
ZC2/g9/Aq9Q6ZNBqoYRkA1123uzHJuxoZ27xTNXg07fEWLIMSwAYTM4EwVEsuDek
d4LIPdMyBlPKSfWXpHqu5wD/SQ5AY+m8LibKoOhqM3TpB4S9B0+xC7eG7/1k/teQ
XkJaYWrYSb/NwfdcnKC47LkaGIk4IF0MvRgo5o7CASLWSRdCuiqapV5B85sji/nD
PUKGURtjsTNR4rL6Ci2xGHwbI+YIPsFQfj2djQIDAQABAoIBAQCG2RQmyDisYdIY
TjBGVC3RS7LPPTKiFvh91OOBSDeW0Y1wV/+JkUZRnCYVudxKtt4PQdwR1sBJIkCg
t1AuMgkyCR24+jGHWUXhggWEpzyiqhK5f0qM1o3YPINVT2n6wTkbOddlnntesP6/
82exNtoprEIktPJai16bOGehCKpvTw9rEPqfbs8gyMYQVimsRjlS5SB0H9xEleFm
yKY6CCvZrEGXTQm0HE8NbXiMAzPBud8HxPhP/iy2y0ojuJEjw1Bg6OpgEVHdE5eM
lc/61qxcvoux6QQF6uhRxVu1oT5qE1N0nTKLgMx3v/00nFwVEeHtwvFDWjJcgr3h
USAIjZedAoGBAOyWiXbd7uRStJvrPEOkSlvp4dEcV3mNLhAJQsMY9IyNWQZPXuUr
cwE0NzOtVR/IFqJnTwWtrXC3GdxkdrkbR+MPav6Gt5sxNwggcxLujV3QHu8bkezz
llq+K3MqEORovTPoDP4ie7bRXG+zXbr3wGXs3BeV5dD4IjHyKeOCsMb7AoGBAN+2
vQPGtthCoaN/H/hsc1cgvKYJ2zX4O2bIwPifoTUcnu+ikxJK0V/3gXQCztptOfLo
SgKT8HlzEuDeAS1xkPXBmAxDI8Jcgts8eP1m/QU4bMpueffNx/hlasYVIUxMVlkM
E2RfOpcUTOkl2Xdb77QwEjMsmaUxj81jH0KeSqcXAoGBANObXT41CGutUZ83ikJR
2rRU/0XiXJb6YF21esI3OtvPvPV43j6JWyaThmAxQzbW5u/BCRviGIP9MSM/tDSi
Iu0CqEcZRXWIUllxdJtVRJnJATIJ7b4KrMti1kf+MveoernHbn+LZi3BGjVKL6Z/
29UOQljSGKiMl0EnALJd/TErAoGBAK3JiJYJZXskb+rR06UpooQ5szMNDxSi7IpR
Q88gOpxGypT/515bbsEtRYw4rjZyiYTQvMW55SKDqVO0Qpfp1CGFrbdA5OOU+DaV
iKvZuI64oZK0Nd2yJzkol8SfH4nk92MfhWUAsUQyCflIZOJbMh/5A/prGESC9uIn
sV4QF5ArAoGBAJWIU+IuTKZdhR0PEPQ6R6HWmzVyYfOQAUlmu+y0Zr+DJAg8rggx
AfiapMBmBYHCnfa9XtmvAmkUXleW936yFK6ersimgdiEWpGTs8U4cNN5+QI/BTV8
zSefGP23eciQeCF5PvTiy//giPkOnfxZkWPnoFpibaQjvd+hUGnjLpjv
-----END RSA PRIVATE KEY-----
`)
	// generated with:
	//  $> openssl req -new -key server.key -out server.csr
	//  $> echo subjectAltName = IP:127.0.0.1 > extfile.cnf
	//  $> openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650 -extfile extfile.cnf
	//  $> openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 3650 -extfile extfile.cnf
	serverCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDIDCCAgigAwIBAgIJAMq1nc6xDi/yMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTYxMDI1MTIwMDA5WhcNMjYxMDIzMTIwMDA5WjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAzsADcua6qCyY1AMwi0JRLw1+bXFv3UX8KOzWUtOGXpjgeq/33UTFZBK1
M9vEQdN7fO+m5+AOLmfxmKQzLdjdO/hCwQZt4i33jFb6fPiZbPx082GaZC2/g9/A
q9Q6ZNBqoYRkA1123uzHJuxoZ27xTNXg07fEWLIMSwAYTM4EwVEsuDekd4LIPdMy
BlPKSfWXpHqu5wD/SQ5AY+m8LibKoOhqM3TpB4S9B0+xC7eG7/1k/teQXkJaYWrY
Sb/NwfdcnKC47LkaGIk4IF0MvRgo5o7CASLWSRdCuiqapV5B85sji/nDPUKGURtj
sTNR4rL6Ci2xGHwbI+YIPsFQfj2djQIDAQABoxMwETAPBgNVHREECDAGhwR/AAAB
MA0GCSqGSIb3DQEBCwUAA4IBAQAsU3uoquRIoiJGM7gquD2SlLscMS3SVIMDe3p/
3XK7FUeemET3IgL68o9qeu7JKUxEgjeGA7zeKNt73ubeBr05S5sYNh8fAR70goke
3TlEFFdeErHt7jtRykSzwL84BTMaKtYlXFV15z2xgoctNkUeOTpSDYab7gBVvnGB
WD/BPnv2sE7/j2feR1qvrire7jbwMe1nP8mw3ovhixrHbtEJz/krIoldylky+Vza
snmUfrvWuC4GZIew+ZWmYu1ozgtQatrSloWHZbAKAdf8piO4FJA2kGNoOUv8LEmH
WAD4c+0FllOU/WwwttDLTT2Vu0tH0255Js9ER5hlUJlzBmLG
-----END CERTIFICATE-----
`)
)
