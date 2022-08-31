package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job        Job
	sumofdigit int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digit(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digit(job.randomno)} //i cant understand this line

		results <- output

	}
	wg.Done() //done func mines counter

}
func createWorkerPool(noOworker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOworker; i++ {
		wg.Add(1)
		go worker(&wg)

	}
	wg.Wait()
	close(results)

}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("job id %d;input random no %d,sum of digits %d\n ", result.job.id, result.job.randomno, result.sumofdigit)

	}
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
