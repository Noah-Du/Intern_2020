package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	//	log "github.com/sirupsen/logrus"
)

func main() {
	// Create variable 'name' to store image's name, 'accelerator' to store registry mirror's url
	var name string
	var accelerator string

	flag.StringVar(&name, "i", "", "image，set to empty")
	flag.StringVar(&accelerator, "a", "hub-mirror.c.163.com/library/", "registry mirror url，set to 163yun")
	flag.Parse()

	if strings.HasPrefix(accelerator, "https://") {
		fmt.Println("Registry Mirror address should not contain 'https://', please eliminate it.")
		log.Fatal("Please fix the error and retry.")
	}

	if !strings.HasSuffix(accelerator, "/library/") {
		fmt.Println("Registry Mirror address should end with '/library/', please fix it.")
		log.Fatal("Please fix the error and retry.")
	}

	download := accelerator + name
	fmt.Printf("Download Address Is: %v \n", download)
	cmd := exec.Command("vctl", "pull", download)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
}
