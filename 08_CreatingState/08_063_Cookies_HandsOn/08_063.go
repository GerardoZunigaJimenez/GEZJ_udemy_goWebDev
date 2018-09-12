package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, createCookie(req,"my-cookie")	)
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, createCookie(req,"general")	)
	http.SetCookie(w, createCookie(req,"specific")	)

	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func createCookie(req *http.Request, cName string) *http.Cookie {
	c, err := req.Cookie(cName)
	var v int64
	if err != nil {
		c = &http.Cookie{
			Name:  cName,
			Value: "1",
		}
	} else {
		v, err = strconv.ParseInt((*c).Value, 10, 64)
		if err != nil {
			log.Println(err)
			return nil
		}
		v++
		(*c).Value = fmt.Sprint(v)
	}
	return c
}
