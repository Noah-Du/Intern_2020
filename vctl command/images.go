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
	CREATIONTIME string
	SIZE         string
	DIGEST       string
	MEASURE      string
}

// Features list
type Features []Feature

func main() {

	var filter string
	var out []byte
	var err error

	flag.StringVar(&filter, "f", "", "filter, set to empty")
	boolPtr := flag.Bool("d", false, "whether see all containers or only running, set to false")
	flag.Parse()

	if *boolPtr == true {
		cmd := exec.Command("vctl", "images", "-d")
		out, err = cmd.Output()
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
				fmt.Sscanf(result[i], "%s %s %s %s %s", &obj.NAME, &obj.CREATIONTIME, &obj.SIZE, &obj.MEASURE, &obj.DIGEST)
				switch req[0] {
				case "NAME":
					if strings.Contains(obj.NAME, req[1]) {
						list = append(list, obj)
					}
				case "CREATIONTIME":
					if strings.Contains(obj.CREATIONTIME, req[1]) {
						list = append(list, obj)
					}
				case "SIZE":
					if strings.Contains(obj.SIZE, req[1]) {
						list = append(list, obj)
					}
				case "DIGEST":
					if strings.Contains(obj.DIGEST, req[1]) {
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
				fmt.Printf("%-16s%-28s%-5s%-6s%s\n", v.NAME, v.CREATIONTIME, v.SIZE, v.MEASURE, v.DIGEST)
			}
		}

	} else {
		cmd := exec.Command("vctl", "images")
		out, err = cmd.Output()
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
				fmt.Sscanf(result[i], "%s %s %s %s", &obj.NAME, &obj.CREATIONTIME, &obj.SIZE, &obj.MEASURE)
				switch req[0] {
				case "NAME":
					if strings.Contains(obj.NAME, req[1]) {
						list = append(list, obj)
					}
				case "CREATIONTIME":
					if strings.Contains(obj.CREATIONTIME, req[1]) {
						list = append(list, obj)
					}
				case "SIZE":
					if strings.Contains(obj.SIZE, req[1]) {
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
				fmt.Printf("%-16s%-28s%-5s%s\n", v.NAME, v.CREATIONTIME, v.SIZE, v.MEASURE)
			}
		}

	}
}
