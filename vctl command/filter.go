package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

// Feature contains containers' features
type Feature struct {
	Name         string
	Image        string
	Command      string
	IP           string
	Ports        string
	Status       string
	CreationTime string
}

// Features list
type Features []Feature

func main() {

	var filter string

	flag.StringVar(&filter, "f", "", "filter, set to empty")

	cmd := exec.Command("vctl", "ps")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	result := strings.Split(string(out), "\n")

	fmt.Println(len(result))

	var list Features

	for i := 3; i < len(result)-1; i++ {
		obj := Feature{}
		fmt.Sscanf(result[i], "%s %s %s %s %s %s %s", &obj.Name, &obj.Image, &obj.Command,
			&obj.IP, &obj.Ports, &obj.Status, &obj.CreationTime)
		list = append(list, obj)
	}

	fmt.Println(list)

}
