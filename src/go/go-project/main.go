package main

import "embed"
import "io/fs"
import "net/http"
import "fmt"

//go:embed static/*
var content embed.FS

func handler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("exampleCookie")
	if err != nil {
		fmt.Print(err)
	} else {
       fmt.Println(cookie.Value)
	}
	new_cookie := http.Cookie{
        Name:     "exampleCookie",
        Value:    "Hello world, new cookie!",
        Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteStrictMode,
    }
	http.SetCookie(w, &new_cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world! This is a 200 OK response.")
}

func main() {
	data, err := content.ReadFile("static/index.html")
	if err == nil {
		print(string(data))

		fsys, err := fs.Sub(content, "static")
		if err == nil {
			http.Handle("/", http.FileServer(http.FS(fsys)))
			http.HandleFunc("/api", handler)
			http.HandleFunc("/api/", handler)
			http.ListenAndServe(":3000", nil)
		} else {
			fmt.Print(err)
		}
	} else {
		fmt.Print(err)
	}
}
