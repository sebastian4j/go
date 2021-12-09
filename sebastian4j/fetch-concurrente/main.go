package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// ./fetch-concurrente https://golang.org http://gopl.io https://godoc.org
func main() {
	start := time.Now()
	ch := make(chan string) // crea un channel de string
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // comienza una nueva goroutine por cada url
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // recibido desde el channel, se bloquea hasta que recibe
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // enviar al chanel
		return
	}

	// lee el contenido pero lo descarta, retorna el numero de bytes
	// leidos
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
