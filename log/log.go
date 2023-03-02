package log

import (
	"fmt"
	"os"
	"time"
)

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// WriteLn baroi navishtani malumot ba log
func WriteLn(text string) {
	if !exists("logs/log" + time.Now().Format("20060102") + ".txt") {
		os.Create("logs/log" + time.Now().Format("20060102") + ".txt")
	}
	f, _ := os.OpenFile("logs/log"+time.Now().Format("20060102")+".txt", os.O_APPEND|os.O_WRONLY, 0600)
	text = getTime() + " " + text
	f.WriteString(text + "\r\n")
	f.Close()
	fmt.Println(text)
}

func getTime() string {
	return time.Now().Format("2006.01.02 15:04:05 ->")
}
