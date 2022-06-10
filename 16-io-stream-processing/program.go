package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	fileReadWg := &sync.WaitGroup{}
	fileReadWg.Add(2)
	go source("data1.dat", dataCh, fileReadWg)
	go source("data2.dat", dataCh, fileReadWg)

	processWg := &sync.WaitGroup{}
	evenCh, oddCh := splitter(dataCh, processWg)
	evenSumCh := sum(evenCh, processWg)
	oddSumCh := sum(oddCh, processWg)
	processWg.Add(1)
	go merger("result.txt", evenSumCh, oddSumCh, processWg)

	fileReadWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Done")

}

func source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if val, err := strconv.Atoi(text); err == nil {
			dataCh <- val
		}
	}
}

func splitter(dataCh chan int, wg *sync.WaitGroup) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	wg.Add(1)
	go func(evenCh, oddCh chan<- int) {
		defer wg.Done()
		defer close(evenCh)
		defer close(oddCh)
		for val := range dataCh {
			if val%2 == 0 {
				evenCh <- val
			} else {
				oddCh <- val
			}
		}
	}(evenCh, oddCh)
	return evenCh, oddCh
}

func sum(valCh <-chan int, wg *sync.WaitGroup) <-chan int {
	resultCh := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := 0
		for val := range valCh {
			result += val
		}
		resultCh <- result
	}()
	return resultCh
}

func merger(fileName string, evenSumCh, oddSumCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	for i := 0; i < 2; i++ {
		select {
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even total : %d\n", evenSum))
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd total : %d\n", oddSum))
		}
	}
}
