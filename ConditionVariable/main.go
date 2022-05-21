package main

import (
	"convar/dataStructure"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	rec := dataStructure.NewRecord()
	cond := sync.NewCond(rec)

	wg.Add(1)
	go func(r *dataStructure.Record) {
		defer wg.Done()

		rec.Lock()
		cond.Wait()
		fmt.Println("Data:", rec.Data)
		rec.Unlock()
	}(rec)

	time.Sleep(2 * time.Second) // Make a pause
	rec.Lock()
	rec.Data = "Gopher"
	rec.Unlock()
	cond.Signal()

	wg.Wait()
}
