package stress

import (
	"net/http"
	"sync"
	"time"
)

type Tester struct {
	URL         string
	Requests    int
	Concurrency int
	Result
}

type job struct {
	url string
}

type Result struct {
	MapStatusCode map[int]int
	TimeElapsed   time.Duration
	TotalRequests int
}

func (tester *Tester) Stress() {
	defer tester.timeTrack(time.Now())
	tester.MapStatusCode = make(map[int]int)
	requests := tester.Requests
	jobs := make(chan job, requests)
	results := make(chan int, requests)
	concurrency := tester.Concurrency
	wg := sync.WaitGroup{}

	for i := 1; i <= concurrency; i++ {
		go worker(jobs, results, &wg)
	}

	wg.Add(requests)
	for i := 1; i <= requests; i++ {
		jobs <- job{tester.URL}
	}
	close(jobs)
	wg.Wait()

	for i := 1; i <= requests; i++ {
		statusCode := <-results
		tester.MapStatusCode[statusCode] += 1
	}
	tester.TotalRequests = requests
}

func worker(jobs <-chan job, out chan<- int, wg *sync.WaitGroup) {
	client := http.Client{}
	for job := range jobs {
		req, err := http.NewRequest("GET", job.url, nil)
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		out <- resp.StatusCode
		wg.Done()
	}
}

func (tester *Tester) timeTrack(start time.Time) {
	tester.TimeElapsed = time.Since(start)
}
