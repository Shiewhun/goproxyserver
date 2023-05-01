package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	target, err := url.Parse("https://www.google.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("Location", target.String())
		resp.Header.Set("Access-Control-Allow-Origin", "*")
		resp.Header.Set("Origin", "https://outlook.live.com")
		resp.StatusCode = http.StatusTemporaryRedirect
		io.Copy(w, resp.Body)
		resp.Body.Close()
		return nil
	}

	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
