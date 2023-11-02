package files

import (
	"fmt"
	"math/rand"
)

type Status struct {
	StatusWater int `json:"water"`
	StatusWind  int `json:"wind"`
}

var water, wind int

func UpdateResp() Status {
	water = rand.Intn(100) + 1
	wind = rand.Intn(100) + 1

	return Status{
		StatusWater: water,
		StatusWind:  wind,
	}
}

func GetResp() {
	statusWater := "aman"

	if water < 6 {
		statusWater = "aman"
	} else if water >= 5 && water <= 8 {
		statusWater = "siaga"
	} else {
		statusWater = "bahaya"
	}

	statusWind := "aman"
	if wind < 6 {
		statusWind = "aman"
	} else if wind >= 6 && wind <= 15 {
		statusWind = "siaga"
	} else {
		statusWind = "bahaya"
	}

	fmt.Println("status water: ", statusWater)
	fmt.Println("status wind: ", statusWind)
}
