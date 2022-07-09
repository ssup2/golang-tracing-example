package sync

import (
	"log"
	"sync"
)

var mutex = sync.Mutex{}

func Mutex01() {
	for {
		mutex.Lock()
		log.Println("Mutex01")
		mutex.Unlock()
	}
}

func Mutex02() {
	for {
		mutex.Lock()
		log.Println("Mutex02")
		mutex.Unlock()
	}
}

func Mutex03() {
	for {
		mutex.Lock()
		log.Println("Mutex03")
		mutex.Unlock()
	}
}
