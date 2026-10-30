package main

import (
	"fmt"
	"log"
	"time"
)

// NOTE: Recovering per job in a background worker
// * A worker loops over jobs, and one bad job must not take the whole process
// * down. Each job runs behind its own recovery boundary: processJob panics,
// * executeJob's deferred function recovers, logs the failure, and returns
// * normally. The loop then moves on to the next job.

func processJob(jobID int) {
	fmt.Printf("processing job %d\n", jobID)
	time.Sleep(time.Second)

	if jobID == 2 {
		panic("corrupted job data")
	}

	fmt.Printf("job %d completed successfully\n", jobID)
}

func executeJob(jobID int) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("job %d failed due to panic: %v", jobID, r)
		}
	}()

	processJob(jobID)
}

func main() {
	for i := range 5 {
		executeJob(i)
	}

	fmt.Println("worker finished processing jobs")
}

// WARNING: A recover only catches panics in its own goroutine. If each job ran
// * in a separate goroutine, this defer would not see the panic; the worker
// * goroutine itself would need the recovery boundary.
//
// Output (the log timestamp reflects the run):
// processing job 0
// job 0 completed successfully
// processing job 1
// job 1 completed successfully
// processing job 2
// 2026/10/30 09:13:18 job 2 failed due to panic: corrupted job data
// processing job 3
// job 3 completed successfully
// processing job 4
// job 4 completed successfully
// worker finished processing jobs
