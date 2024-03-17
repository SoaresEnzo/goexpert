package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

// Test with apache benchmark - ab -n 10000 -c 100 http://localhost:3000/
// go run -race main.go        Runs checking for race conditions

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("You had acess to this page %d times \n", number)
		w.Write([]byte("You had acess to this page " + fmt.Sprint(number) + " times \n"))
	})

	http.ListenAndServe(":3000", nil)
}
