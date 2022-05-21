package dataStructure

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

type Record struct {
	sync.Mutex
	Data string
}

func NewRecord() *Record {
	return new(Record)
}

type RecordWriter struct {
	Record
	cond    *sync.Cond
	writers []io.Writer
}

func NewRecordWriter(writers ...io.Writer) *RecordWriter {
	rec := new(RecordWriter)
	rec.writers = writers
	rec.cond = sync.NewCond(rec)
	return rec
}

func (r *RecordWriter) Start() error {
	for _, writer := range r.writers {
		go func(w io.Writer) {
			for {
				r.Lock()
				r.cond.Wait()
				fmt.Fprintf(w, "%s\n", r.Data)
				r.Unlock()
			}
		}(writer)
	}
	return nil
}

func (r *RecordWriter) Prompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text:")
	for scanner.Scan() {
		r.Lock()
		r.Data = scanner.Text()
		r.Unlock()
		r.cond.Broadcast()
	}
}
