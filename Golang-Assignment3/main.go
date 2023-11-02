package main

import (
	data "Golang-Assignment3/files"
	"time"
)

func main() {
	go func() {
		for {
			data.UpdateResp()
			data.PrintStatus()
			data.GetResp()
			time.Sleep(15 * time.Second)
		}
	}()
	select {}
}
