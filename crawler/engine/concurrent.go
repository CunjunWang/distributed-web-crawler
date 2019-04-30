package engine

import (
	"distributed-web-crawler/crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			// log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			_, ok := item.(model.Profile)
			if ok {
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}
		}

		// URL de-dup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				// log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, notifier ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			notifier.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visited = make(map[string]bool)

func isDuplicate(url string) bool {
	if visited[url] {
		return true
	}
	visited[url] = true
	return false
}
