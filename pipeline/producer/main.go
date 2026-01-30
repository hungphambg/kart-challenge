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
)

const BatchSize = 10000

func main() {
	fmt.Println("Coupon Producer Service started.")
	if len(os.Args) < 3 {
		fmt.Println("Usage: producer <path_to_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileIdx, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error converting file index: %v\n", err)
		return
	}
	fmt.Printf("process file: %s, file index %d\n", filePath, fileIdx)

	var (
		ctx    = context.Background()
		writer = NewCouponRedisWriter("redis:6379", fileIdx)
		start  = time.Now()
	)

	processFileAndSend(ctx, filePath, writer)
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

func processFileAndSend(ctx context.Context, filePath string, writer CouponWriter) {
	startTime := time.Now()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 1024*1024) // 1M

	jobs := make(chan []string, 1000) // backpressure
	go readBatches(reader, jobs)

	workers := 10
	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(ctx, jobs, &wg, writer)
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

func worker(ctx context.Context, jobs <-chan []string, wg *sync.WaitGroup, write CouponWriter) {
	defer wg.Done()
	for batch := range jobs {
		processBatch(ctx, write, batch) // DB / Kafka / validate
	}
}

func processBatch(ctx context.Context, write CouponWriter, batch []string) {
	write.Write(ctx, batch)
}

func logf(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
