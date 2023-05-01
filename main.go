package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(r.Method, "https://outlook.live.com"+r.RequestURI, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
		w.WriteHeader(resp.StatusCode)
		w.Write(respBody)

		http.Redirect(w, r, "https://www.google.com", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
