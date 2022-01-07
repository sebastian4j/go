package main

import (
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	defer savior()
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
	defer trace()()
	time.Sleep(2 * time.Second)
	defer defer2()
	defer func() {
		log.Println("defer 1")
	}()
	_ = defer2()

}

func savior() {
	if p := recover(); p != nil {
		log.Printf("algo salió mal: %v\n", p)
	}
}

func defer2() int {
	defer func() { log.Println("defer2 usada") }()
	log.Println("defer 2")
	panic("ddd panic ddd")
	return 0
}

func trace() func() {
	start := time.Now()
	log.Printf("inicia: %v", start)
	return func() {
		log.Printf("termina en: %s segundos", time.Since(start))
	}
}
