package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// Устанавливаем количество строк лога и имя файла
	linesCount := 1000000
	fileName := "app_log.log"

	// Создаем файл для записи
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())

	// Генерация и запись псевдологов в файл
	for i := 0; i < linesCount; i++ {
		logLine := fmt.Sprintf("%s - - [%s] \"GET / HTTP/1.1\" 200 %d\n",
			randomIP(), time.Now().Format("02/Jan/2006:15:04:05 -0700"), rand.Intn(5000)+1000)
		_, err := file.WriteString(logLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Generated %d lines of log data in %s\n", linesCount, fileName)
}

func randomIP() net.IP {
	ip := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		ip[i] = byte(rand.Intn(256))
	}
	return ip
}
