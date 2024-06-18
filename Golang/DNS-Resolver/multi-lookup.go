package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	queueMutex sync.Mutex
	fileMutex  sync.Mutex
)

func requesterThreadFunc(arg string) {
	inputfp, err := os.Open(arg)
	if err != nil {
		fmt.Println("Error Opening Input File: ", arg)
		return
	}
	defer inputfp.Close()

	var hostname string
	for {
		_, err := fmt.Fscanf(inputfp, "%s", &hostname)
		if err != nil {
			break
		}
		queueMutex.Lock()
		// Sleep until there is space in the queue
		queueMutex.Unlock()
	}
}

func resolverThreadFunc() {
	for {
		queueMutex.Lock()
		// Sleep after removing hostname from queue
		queueMutex.Unlock()
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments: ", len(os.Args)-1)
		fmt.Println("Usage: ", os.Args[0], " <inputfile1> ... <inputfileN> ")
		return
	}

	num_files := len(os.Args) - 2

	outputfp, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Error opening output file, filename was ", os.Args[len(os.Args)-1])
		return
	}
	defer outputfp.Close()

	var requesterThread [num_files]sync.WaitGroup
	var resolverThread [nprocs]sync.WaitGroup

	for i := 0; i < num_files; i++ {
		requesterThread[i].Add(1)
		go func(i int) {
			requesterThreadFunc(os.Args[i+1])
			requesterThread[i].Done()
		}(i)
	}

	for i := 0; i < nprocs; i++ {
		resolverThread[i].Add(1)
		go func(i int) {
			resolverThreadFunc()
			resolverThread[i].Done()
		}(i)
	}

	for i := 0; i < num_files; i++ {
		requesterThread[i].Wait()
	}

	for i := 0; i < nprocs; i++ {
		resolverThread[i].Wait()
	}
}
