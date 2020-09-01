package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Feature contains containers' features
type Feature struct {
	NAME         string
	IMAGE        string
	COMMAND      string
	IP           string
	PORTS        string
	STATUS       string
	CREATIONTIME string
}

// Features list
type Features []Feature

func main() {

	var filter string

	flag.StringVar(&filter, "f", "", "filter, set to empty")
	flag.Parse()

	cmd := exec.Command("vctl", "ps")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	if filter == "" {
		fmt.Printf("%s\n", string(out))
	} else {

		req := strings.Split(filter, "=")
		req[0] = strings.ToUpper(req[0])

		result := strings.Split(string(out), "\n")

		var list Features

		for i := 3; i < len(result)-1; i++ {
			obj := Feature{}
			fmt.Sscanf(result[i], "%s %s %s %s %s %s %s", &obj.NAME, &obj.IMAGE, &obj.COMMAND,
				&obj.IP, &obj.PORTS, &obj.STATUS, &obj.CREATIONTIME)
			switch req[0] {
			case "NAME":
				if strings.Contains(obj.NAME, req[1]) {
					list = append(list, obj)
				}
			case "IMAGE":
				if strings.Contains(obj.IMAGE, req[1]) {
					list = append(list, obj)
				}
			case "COMMAND":
				if strings.Contains(obj.COMMAND, req[1]) {
					list = append(list, obj)
				}
			case "IP":
				if strings.Contains(obj.IP, req[1]) {
					list = append(list, obj)
				}
			case "PORTS":
				if strings.Contains(obj.PORTS, req[1]) {
					list = append(list, obj)
				}
			case "STATUS":
				if strings.Contains(obj.STATUS, req[1]) {
					list = append(list, obj)
				}
			case "CREATIONTIME":
				if strings.Contains(obj.CREATIONTIME, req[1]) {
					list = append(list, obj)
				}
			default:
				log.Fatal("Input is wrong, please try again")
			}
		}

		for i := 0; i < 3; i++ {
			fmt.Printf("%s\n", result[i])
		}

		for _, v := range list {
			fmt.Printf("%-14s%-16s%-12s%-18s%-8s%-10s%-9s\n", v.NAME, v.IMAGE, v.COMMAND, v.IP, v.PORTS,
				v.STATUS, v.CREATIONTIME)
		}
	}
}
