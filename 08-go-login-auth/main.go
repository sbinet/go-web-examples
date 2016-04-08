package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
	"sync"

	"github.com/satori/go.uuid"
)

const loginPage = `
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<form action="login" method="post">
			<input type="password" name="password"/>
			<input type="submit" value="login"/>
		</form>
	</body>
</html>
`

func main() {
	http.Handle("/hello", helloWorldHandler{})
	http.Handle("/", authenticate(helloWorldHandler{secure: true}))
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":5555", nil)
}

var (
	sessionStore = make(map[string]Client)
	mu           sync.RWMutex
)

type Client struct {
	loggedIn bool
}

type helloWorldHandler struct {
	secure bool
}

func (h helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.secure {
		fmt.Fprintln(w, "Hello World (Secure)")
	} else {
		fmt.Fprintln(w, "Hello World (NOT SECURE)")
	}
}

type authMiddleware struct {
	handler http.Handler
}

func (auth authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client, _, err := doLogin(w, r)
	if err != nil {
		return
	}

	if !client.loggedIn {
		fmt.Fprint(w, loginPage)
		return
	}
	auth.handler.ServeHTTP(w, r)
}

func authenticate(h http.Handler) authMiddleware {
	return authMiddleware{h}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	client, cookie, err := doLogin(w, r)
	if err != nil {
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	if subtle.ConstantTimeCompare([]byte(r.FormValue("password")), []byte("password123")) != 1 {
		fmt.Fprintln(w, "wrong password")
		return
	}

	client.loggedIn = true
	fmt.Fprintln(w, "thank you. logged in.")
	mu.Lock()
	sessionStore[cookie] = client
	mu.Unlock()
}

func doLogin(w http.ResponseWriter, r *http.Request) (Client, string, error) {
	var (
		ok     = false
		client Client
	)
	cookie, err := r.Cookie("session")
	if err != nil {
		if err != http.ErrNoCookie {
			fmt.Fprint(w, err)
			return client, "", err
		}
		err = nil
	}

	if cookie != nil {
		mu.RLock()
		client, ok = sessionStore[cookie.Value]
		mu.RUnlock()
	}
	if !ok {
		cookie = &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		}
		client = Client{loggedIn: false}
		mu.Lock()
		sessionStore[cookie.Value] = client
		mu.Unlock()
	}

	http.SetCookie(w, cookie)
	return client, cookie.Value, nil
}
