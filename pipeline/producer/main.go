package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	fmt.Println("process file: ", filePath)
	topic := strings.TrimSuffix(fileName, filepath.Ext(fileName))

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

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Ticker ticked, processing file: %s\n", filePath)
			processFileAndSend(filePath, writer, topic)
		}
	}
}

func processFileAndSend(filePath string, writer *kafka.Writer, topic string) {
	fmt.Printf("Attempting to open file: %s\n", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	//couponRegex := regexp.MustCompile(`\b[a-zA-Z0-9]{8,10}\b`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//matches := couponRegex.FindAllString(line, -1)
		couponCode := strings.TrimSpace(line)
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

	if err = scanner.Err(); err != nil && err != io.EOF {
		fmt.Printf("Error reading content from %s: %v\n", filePath, err)
	}

	fmt.Printf("Successfully processed file: %s\n", filePath)
}

func logf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
