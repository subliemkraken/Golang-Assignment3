package files

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintStatus() {
	status := UpdateResp()
	jsonStatus, err := json.MarshalIndent(status, "", "  ")

	if err != nil {
		log.Panicf("Error while printing status:%s\n", err.Error())
	}

	fmt.Println(string(jsonStatus))
}
