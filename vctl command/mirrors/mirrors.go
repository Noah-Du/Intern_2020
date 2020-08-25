package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func foo() {

	log.SetReportCaller(true)

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	log.SetOutput(os.Stdout)

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	foo()

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
