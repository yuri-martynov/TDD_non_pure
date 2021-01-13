package racer

import ("net/http"; "time")


func racer(first string, second string, ping func(string) chan bool) string {
    select {
    case <-ping(first):
        return first
    case <-ping(second):
		return second
	// case <- time.After(10 * time.Second):
		// return "timeout"
    }
}


func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}