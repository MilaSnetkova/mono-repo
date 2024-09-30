package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type StringWriter struct {
	m      sync.Mutex
	Writer io.Writer
}

const gorutines  = 2 
const lines  = 500000

func (s *StringWriter) WriteString(str string) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.Writer.Write([]byte(str))
}

func do(writer *StringWriter,id int, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	for i:=0; i < lines; i++ {
		<-ch 
		_, err := writer.WriteString(fmt.Sprintf("Строка горутины %d: %d\n", id, i))
		if err != nil {
			fmt.Println("Ошибка записи:", err)
		}
		ch <-struct{}{}
	}

}

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Файл не создан")
	}

	defer file.Close()
	if err != nil {
		fmt.Println("Файл не закрыт")
	}
	writer := &StringWriter{Writer: file}

	var wg sync.WaitGroup

	ch := make(chan struct{}, 1)

	for i :=1; i <= gorutines; i++ {
	wg.Add(1) 
	go do(writer, i, &wg, ch)
	} 

	ch <- struct{}{}

	wg.Wait()
	
	close(ch)
	
	fmt.Println("Done!)")
}
