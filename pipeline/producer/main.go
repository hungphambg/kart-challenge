package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const BatchSize = 10000

func main() {
	fmt.Println("Coupon Producer Service started.")
	if len(os.Args) < 3 {
		fmt.Println("Usage: producer <path_to_gz_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileIdx, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		fmt.Printf("Error converting first argument: %v\n", err1)
		return
	}
	fmt.Println("process file: ", filePath)

	state := NewCouponState("redis:6379")

	start := time.Now()
	processFileAndSend(filePath, fileIdx, nil, state)
	fmt.Println("Coupon Producer Service Done after ", time.Since(start))

	//ticker := time.NewTicker(5 * time.Minute)
	//defer ticker.Stop()
	//
	//// Process the file immediately on startup
	//processFileAndSend(filePath, writer, topic)
	//
	//for {
	//	select {
	//	case <-ticker.C:
	//		fmt.Printf("Ticker ticked, processing file: %s\n", filePath)
	//		processFileAndSend(filePath, writer, topic)
	//	}
	//}
}
func processFileAndSend(filePath string, fileIdx int, writer *kafka.Writer, couponState *CouponState) {
	startTime := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024) // 1MB
	jobs := make(chan []string, 100)               // backpressure

	go readBatches(reader, jobs)

	workers := 10
	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(jobs, writer, &wg, couponState, fileIdx)
	}

	wg.Wait()
	fmt.Println("Time elapsed: ", time.Since(startTime))
}

func readBatches(reader *bufio.Reader, out chan<- []string) {
	defer close(out)

	batch := make([]string, 0, BatchSize)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			if line != "" {
				batch = append(batch, line)
			}
		}

		if len(batch) >= BatchSize {
			out <- batch
			batch = make([]string, 0, BatchSize)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if len(batch) > 0 {
		out <- batch
	}
}

func worker(jobs <-chan []string, writer *kafka.Writer, wg *sync.WaitGroup,
	couponState *CouponState, fileIdx int) {
	defer wg.Done()
	for batch := range jobs {
		processBatch(couponState, writer, batch, fileIdx) // DB / Kafka / validate
	}
}

func processBatch(couponState *CouponState, writer *kafka.Writer, batch []string, fileIdx int) {
	if couponState != nil {
		couponState.AddObservation(batch, fileIdx)
	}
	if writer != nil {
		msgs := make([]kafka.Message, 0, len(batch))
		for _, code := range batch {
			msg := kafka.Message{
				Key:   []byte(code),
				Value: []byte(code),
				Time:  time.Now(),
			}
			msgs = append(msgs, msg)
		}
		err := writer.WriteMessages(context.Background(), msgs...)
		if err != nil {
			fmt.Printf("Failed to write message to Kafka: %v\n", err)
		}
	}
}

func logf(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
