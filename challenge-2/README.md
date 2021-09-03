## How to correctly stop and debug Goroutines?

Check out the code explained below at [main.go](github.com/ebikode/peaq-challenge/challenge2).

To achieve this, there has to be a way of synchronizing and sending signal to the goroutine processes. Channel was used as a messenger between goroutines and other application programs. Channels are one of the powerful feature of the golang language that facilitate communication feature between goroutines.  The [Sync](https://golang.org/pkg/sync/) package was used to manage the synchronization and process workflow.

To correctly stop and debug Goroutines, the following process was considered:

1. Initializing a `sync.WaitGroup` method. This tracks the number of running goroutine and their current status. When a goroutine process is completed, the task is taken away from the waitGroup by calling the `WaitGroup.Done()` method. This continues until the total number of task added has been completed and removed.

2. Created a boolean channel for sending signal to the goroutine. Exited the goroutine by sending a `false` boolean value to notify the program to log the debug error and stop the goroutine by calling the `WaitGroup.Done()` method of the sync package.






