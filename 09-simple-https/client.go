// +build ignore

package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(rootCA) {
		log.Fatal("invalid cert")
	}

	tlsConf := &tls.Config{
		RootCAs: roots,
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConf,
	}
	http2.ConfigureTransport(tr) // enable HTTP/2 see https://github.com/golang/go/issues/17051
	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("https://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
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
)
