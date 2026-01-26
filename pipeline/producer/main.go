package main

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	fmt.Println("Coupon Producer Service started.")

	if len(os.Args) < 2 {
		fmt.Println("Usage: producer <path_to_gz_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileName := filepath.Base(filePath)
	topic := strings.TrimSuffix(fileName, filepath.Ext(fileName)) // e.g., "couponbase1" from "couponbase1.gz"

	fmt.Printf("Attempting to open GZ file: %s\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("Error creating gzip reader for %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer gzr.Close()

	couponRegex := regexp.MustCompile(`\b[a-zA-Z0-9]{8,10}\b`)

	// Initialize Kafka producer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"kafka:9092"}, // Kafka broker address from docker-compose.pipeline.yaml
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Logger:       kafka.LoggerFunc(logf),
		ErrorLogger:  kafka.LoggerFunc(logf),
		Async:        true,        // Publish messages asynchronously
		RequiredAcks: kafka.NoAck, // Fire and forget for maximum throughput
	})
	defer writer.Close()

	scanner := bufio.NewScanner(gzr)
	for scanner.Scan() {
		line := scanner.Text()
		matches := couponRegex.FindAllString(line, -1)

		for _, match := range matches {
			couponCode := strings.TrimSpace(match)
			fmt.Printf("Identified couponLib code: %s, publishing to topic: %s\n", couponCode, topic)

			msg := kafka.Message{
				Key:   []byte(couponCode),
				Value: []byte(couponCode),
				Time:  time.Now(),
			}

			err = writer.WriteMessages(context.Background(), msg)
			if err != nil {
				fmt.Printf("Failed to write message to Kafka: %v\n", err)
			}
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		fmt.Printf("Error reading decompressed content from %s: %v\n", filePath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed GZ file: %s\n", filePath)
}

func logf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
