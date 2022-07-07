package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

//nolint:gocognit //Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	resChan := make(chan error)
	jobChan := make(chan Task, len(tasks))

	for i := 0; i < len(tasks); i++ {
		jobChan <- tasks[i]
	}

	failures := 0
	isRunning := true
	freeWorkers := n

	var graceful bool
	var waitList int
	var errOverload error
	var finished int

	for isRunning {
		select {
		case job, ok := <-jobChan:
			if ok && freeWorkers > 0 && !graceful {
				go func() {
					resChan <- job()
				}()
				freeWorkers--
				waitList++
			}
			if !ok && waitList <= 0 {
				graceful = true
				isRunning = false
				close(jobChan)
			}

		case e, ok := <-resChan:
			if ok && freeWorkers < n {
				freeWorkers++
				waitList--
			}
			if e != nil {
				failures++
			}
			if failures >= m {
				graceful = true
			}
			if failures >= m && waitList <= 0 {
				errOverload = ErrErrorsLimitExceeded
				isRunning = false
				close(resChan)
				close(jobChan)
			}
			finished++
			if waitList <= 0 && finished == len(tasks) {
				graceful = true
				isRunning = false
				close(resChan)
				close(jobChan)
			}
		}
	}

	if errOverload != nil {
		return ErrErrorsLimitExceeded
	}

	return nil
}
