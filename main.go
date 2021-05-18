package main

import "time"

func main() {
	for i := 0; i < 6; i++ {
		go func() {
			for {
				//fmt.Println("1")
			}

		}()
	}

	time.Sleep(time.Hour)
}
