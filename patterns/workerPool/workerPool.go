package workerPool

import (
	"fmt"
	"sync"
	"time"
)

const (
	numWorkers      = 50    // number of workers (goroutines)
	numTasks        = 1000  // number of tasks
	memoryIntensity = 10000 // size of memory-intensive task (number of elements)
)

func TryWorkerPool() {

	// initialise task queue and worker queue using channels
	taskQueue := make(chan int, numTasks)
	resultQueue := make(chan float64, numTasks)

	// start the workers
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		// call worker function
		go startWorker(taskQueue, resultQueue, &wg)
	}

	// send tasks to the queue
	for i := 0; i < numTasks; i++ {
		taskQueue <- i
	}
	// close the queue if your sent all the tasks, if not, then keep it open for later use
	close(taskQueue)

	// retrieve results from the queue
	go func() {
		wg.Wait()
		// again, keep it open or closed as per your requirement
		close(resultQueue)
	}()

	var totalTime float64

	// process the results
	resultIndex := 0
	for result := range resultQueue {
		resultIndex += 1
		taskTime := result
		totalTime += taskTime
		fmt.Printf("Task[%d] processed: %f ms\n", resultIndex, taskTime)
	}

	fmt.Printf("\n\nworker pool example is now finished!\n\n")
	fmt.Printf("total time taken: %fms\n", totalTime)
	fmt.Printf("average time taken for each task: %fms\n\n", totalTime/numTasks)
}

func startWorker(taskQueue <-chan int, resultQueue chan<- float64, wg *sync.WaitGroup) {

	// signal the waitGroup when you are done in the last
	defer wg.Done()

	for task := range taskQueue {
		// start the timer
		startTime := time.Now()
		_ = performMemoryIntensiveTask(task)

		endTime := time.Since(startTime)
		resultQueue <- endTime.Seconds()
	}
}

func performMemoryIntensiveTask(task int) int {

	// create a large-sized slice
	data := make([]int, memoryIntensity)
	for i := 0; i < memoryIntensity; i++ {
		data[i] = i + task
	}

	// latency imitation
	time.Sleep(10 * time.Millisecond)

	// calculate the result
	result := 0
	for _, value := range data {
		result += value
	}

	// return the result
	return result
}
