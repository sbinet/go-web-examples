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
	//  $> openssl genrsa -out server.key 2048
	serverKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAo7lDaxphoYEpmSVm5Se3tj6lJ6iUXXI8g0hMYCklvyJO2HhJ
SJRCDr5Nli1mKHfMpZ5gqw3Si+OVWQFW2vKqvQEmq/SczbIZ7DNMQ1L66kEcOIMN
gsssmpbnvi0iyvgwACyDPfZ5+Z+6bcuBwJ70btwciSiG0CSJlvbPMLKJ8dtgtlcM
IKsYWxjPIt2JlOoU5ZE+rNzh/oWBD9YFwhLgfvYzzkenMeHSNyLV78JJQ6sJHRA4
nCi0zt83ttogKMTRgKpizHGTHvvHF7tCWWc3MVj2QRipLSNK7+3DB0QyAk/rWeqz
Qn5Bw6JXfKAT4uli/UmCUL+9eva3cd/Mg7gz0wIDAQABAoIBABuqOGwmnwytSJwq
J8Lc/Tv8Ref3ompP3U3Jr64oBcrQP0ZwvOFYu/jFy1JvBW4dICV0J51/zm343MmX
YlfI3XTmduRjhwNy7tJssJxHr0JaEiyyaFwLfNP6X9pQwipN2b6Nxvd3aJD3nobi
9l/X/DGnW+MJSA5vvhgWSFhQMuL9C4K+MtRHBljzChCTWOGBUZJXUT4Y1FcCV5eb
WhocMoSoYNEd42UAxD9hrF/3VS9X/Lcec5LGRtWXGWIP0QA7LHye0dguNu1Dw4SA
kCPuYiVCD5grEHZNw3801D+XRD6Eg76YwySdAs1FQ5oggAQby1DvdFJ0IWBIplx8
7+lEXGECgYEA2Ox6qqauRCIs6dn+ZBjzljZf4njAmFWgkWCy8guFqyavK3ZblI4W
vyV7wHPGqSqiBvPJQ4NJ6dKZSEP4AlJTdua6TvNxWuwoGyKMN7SSe45urnxk2s+Q
XTxGTcZ/Sut6DVJDzEzlLuxf6CQs/5dzz5SNhobmy9MvDeqRSBktJLkCgYEAwTdx
c0/fFffiQWR8udSunyYWUpmvlJwGfRH24jVdnl6jGPQKHQWfH3TTzFmLhsq7s//e
0kzvB72UskD39fuEzVBNagFoi4ZnGF8PCFIbAevGnSh2BbkFOYD5QmKrke4thY7I
kR3Rfjyma3nUE2b5IES/AYYOrBv3aE/Lnk2AbusCgYEApO5gcgHbfBhT99X0ctz5
z9s12jafkPOB6ycPx7L1BeWayDvsc635I4p3ZiNhB9xPZ2PSQg+/khW6z9RF7FX1
+fCB9WdpZ40pDUCeRfh4a5MnGmmgMTIh/JHIjnf+7tLNcPV8cQiCNMQqQ5HF2Oup
dUnotUE6l/zQyQ7xOVVTGpkCgYBEBiICB++xi0jz3s3umsszqrQXNpZTSq8aH1X9
A+7Gz+i540kgOJdx+q6KdkgiF4QH9iBbh8xrvDn6m9bwmjGCGzB2DrLg1Fu9f9IK
CuYJQhn1wIX0s9P8D8Uxsw1WXjVWnRNNy+Kyf+XTVtvsTMeOrcVVYv4NZ/ctmVJF
lGGQKwKBgQC96hgBTtbMkvO2seTCEC6OcMQKxtWnE/FG5Gi1nNagOkviTbifDl8t
1SWOQvZ/PDCc6q1AHrHw0tZI67bXtE+byhBjHYvBJaRlldWFRFFQ8Mx6Uit16znq
ubgfh1Ep7omDV/7BkD+1/fDSWd/sZ+XyB8q+WzuudQrrpE2kTmOxOg==
-----END RSA PRIVATE KEY-----
`)

	// generated with:
	//  $> openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
	serverCert = []byte(`-----BEGIN CERTIFICATE-----
MIIDXTCCAkWgAwIBAgIJAIU8V+U6Yj5YMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQwHhcNMTYxMDI1MTEyOTU5WhcNMjYxMDIzMTEyOTU5WjBF
MQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50
ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB
CgKCAQEAo7lDaxphoYEpmSVm5Se3tj6lJ6iUXXI8g0hMYCklvyJO2HhJSJRCDr5N
li1mKHfMpZ5gqw3Si+OVWQFW2vKqvQEmq/SczbIZ7DNMQ1L66kEcOIMNgsssmpbn
vi0iyvgwACyDPfZ5+Z+6bcuBwJ70btwciSiG0CSJlvbPMLKJ8dtgtlcMIKsYWxjP
It2JlOoU5ZE+rNzh/oWBD9YFwhLgfvYzzkenMeHSNyLV78JJQ6sJHRA4nCi0zt83
ttogKMTRgKpizHGTHvvHF7tCWWc3MVj2QRipLSNK7+3DB0QyAk/rWeqzQn5Bw6JX
fKAT4uli/UmCUL+9eva3cd/Mg7gz0wIDAQABo1AwTjAdBgNVHQ4EFgQU8cMQWrwA
5OSiE7ssTlN9C0Typ+EwHwYDVR0jBBgwFoAU8cMQWrwA5OSiE7ssTlN9C0Typ+Ew
DAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAn+0zHOHLqC4czeZhWHqO
vXdnU6T7Oj+wYN+aYChQHR1z2nOXR88VvCp/lXOiXAmNpLwsqHvrdEydF5OIiyvI
dQcut4wdaono76wA5eqhOHpYN8vm0Hg7Twv1xS2aIagBQHZnIe+Z6BZGTI713Yrz
tq77WGj/HbHptDu4ZrJnIYIlfdEqtosPCRVZOJIiBd7zVxIiad2l2RKJl+19Nxf8
LPM1uIYwkMbkmW0SxO+p6RZllQCbHVjdjQzJ8VAdIM/tOSR/gKXfn/momcHI/vQy
3k0brlGfbO4EbBKVDHurFcPEOzG/68UJHDU/tKnM30MmUyNMpoVOhEy3sIamKzPz
BQ==
-----END CERTIFICATE-----
`)
)
