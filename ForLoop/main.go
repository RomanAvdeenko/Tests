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
	wg.Add(1)

	go func(r *dataStructure.Record) {
		defer wg.Done()

		for {
			r.Lock()
			if rec.Data != "" {
				fmt.Println("Data:", rec.Data)
				return
			}
			r.Unlock()
			time.Sleep(time.Millisecond) // To reduce CPU wasting
		}
	}(rec)

	time.Sleep(2 * time.Second) // Make a pause
	rec.Lock()
	rec.Data = "Gopher"
	rec.Unlock()

	wg.Wait()
}
