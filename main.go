package main

import (
	"fmt"
	"net/http"
	"golang.org/x/sync/errgroup"
	"k8s.io/klog"
)


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome my web server!")
}

func main() {
	http.HandleFunc("/", index)
	var eg errgroup.Group
	// 启动http server
	eg.Go(func() error {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			klog.Fatal("ListenAndServe: ", err)
		}
		return nil
	})
	// 启动https server
	eg.Go(func() error {
		err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
		if err != nil {
			klog.Fatal("ListenAndServe: ", err)
		}
		return nil
	})
	eg.Wait()
}
