package main

import (
	"fmt"

	utils "github.com/zakiego/activity-notification/utils"
)

func main() {
	Run()
}

func Run() {
	if utils.IsOnline() {
		utils.SendNotif("Zaki menyalakan laptop 🟢")
		fmt.Println("Success")
	} else {
		fmt.Println("Try again, no internet connection")
		Run()
	}
}
