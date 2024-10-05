package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"flag"
)

type StringWriter struct {
	m      sync.Mutex
	Writer io.Writer
}


func (s *StringWriter) WriteString(str string) (n int, err error) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.Writer.Write([]byte(str))
}

func do(writer *StringWriter,id int, wg *sync.WaitGroup, ch chan struct{}, lines int) {
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
	filenamePtr := flag.String("filename", "file.txt", "Название файла для записи")
	parallelPtr := flag.Int("parallel", 2, "Количество горутин")
	linesPtr := flag.Int("lines", 500000, "Количество строк")

	flag.Parse()


	file, err := os.Create(*filenamePtr)
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

	for i :=1; i <= *parallelPtr; i++ {
	wg.Add(1) 
	go do(writer, i, &wg, ch, *linesPtr) 
	} 

	ch <- struct{}{}

	wg.Wait()
	
	close(ch)
	
	fmt.Println("Done!)", *filenamePtr)
}