package main

import "time"

func main(){
	quit := make(chan bool)
	for i := 0; i < 100 ; i++ {
		go func() {
			for {
				select {
				case <-quit:
					break
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 15)
	for i := 0; i != 5; i++ {
		quit <- true
	}
}
