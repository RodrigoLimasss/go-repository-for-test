package concurrents

import (
	"fmt"
	"runtime"
	"os"
	"time"
)

func Run() {

	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh //channel output
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			}

		}
	}()

	for i := 1; i < 10; i++ {
		go func(value int) {
			fmt.Printf("%d\n", value)
			logCh <- fmt.Sprintf("Message: %d \n", value) //channel Input
		}(i)
	}
	fmt.Scanln()
}
