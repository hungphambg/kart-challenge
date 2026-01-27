package main

import (
	"bufio"
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
	topic := strings.TrimSuffix(fileName, filepath.Ext(fileName))

	fmt.Printf("Attempting to open file: %s\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	couponRegex := regexp.MustCompile(`\b[a-zA-Z0-9]{8,10}\b`)
	// Initialize Kafka producer
	writer := &kafka.Writer{
		Addr:         kafka.TCP("kafka:9092"),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Logger:       kafka.LoggerFunc(logf),
		ErrorLogger:  kafka.LoggerFunc(logf),
		Async:        true,
		RequiredAcks: kafka.RequireNone,
	}
	defer func() {
		_ = writer.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := couponRegex.FindAllString(line, -1)

		for _, match := range matches {
			couponCode := strings.TrimSpace(match)
			fmt.Printf("Identified coupon code: %s, publishing to topic: %s\n", couponCode, topic)

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

	if err = scanner.Err(); err != nil && err != io.EOF {
		fmt.Printf("Error reading decompressed content from %s: %v\n", filePath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully processed GZ file: %s\n", filePath)
}

func logf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
