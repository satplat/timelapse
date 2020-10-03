package main

import (
	"App/internal"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	data, _ := ioutil.ReadFile("land.txt")
	coordinate := string(data)
	token := "GET TOKEN"
	fmt.Println("Start Registering Land...")
	landID, err := internal.RegisterLand(coordinate, token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Land Registered!")
	fmt.Println("Start Getting Images...")
	err = internal.GetImagesData(landID, token)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Getting Images Done!")
	fmt.Println("Starting Creating Timelapse...")
	err = exec.Command("/bin/bash", ".bash.sh").Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Jobs Done! Enjoy Your timelapse.mp4")
}
