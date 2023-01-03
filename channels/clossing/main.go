package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, open := <-jobs
			if open {
				fmt.Println("Received a job ", j)
			} else {
				fmt.Println("Received all jobs. Closing...")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("Sent a job ", i)
	}
	close(jobs)

	<-done
}
