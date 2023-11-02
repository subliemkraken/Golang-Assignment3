package main

import (
	"Golang-Assignment3/files"
	"time"
)

func main() {
	go func() {
		for {
			files.UpdateResp()
			files.PrintStatus()
			files.GetResp()
			time.Sleep(15 * time.Second)
		}
	}()
	select {}
}
