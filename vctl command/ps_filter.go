package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
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
	var format string
	var out []byte
	var err error
	var t *template.Template
	t = template.New("kawaii")

	var list Features
	obj := Feature{}

	flag.StringVar(&filter, "f", "", "filter, set to empty")
	flag.StringVar(&format, "format", "", "format, set to empty")
	boolPtr := flag.Bool("a", false, "whether see all containers or only running, set to false")
	flag.Parse()

	if *boolPtr == true {
		cmd := exec.Command("vctl", "ps", "-a")
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		cmd := exec.Command("vctl", "ps")
		out, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
	}

	if filter == "" {
		if format == "" {
			fmt.Printf("%s\n", string(out))
		} else {

		}
	} else {

		req := strings.Split(filter, "=")
		req[0] = strings.ToUpper(req[0])

		result := strings.Split(string(out), "\n")

		for i := 3; i < len(result)-1; i++ {

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
			case "EXITED":
				if strings.Contains(obj.STATUS, req[1]) {
					list = append(list, obj)
				}
			case "ANCESTOR": //need revise
				if strings.Contains(obj.IMAGE, req[1]) {
					list = append(list, obj)
				}
			case "PUBLISH":
				if strings.Contains(obj.PORTS, req[1]) {
					list = append(list, obj)
				}
			default:
				log.Fatal("Input is wrong, please try again")
			}
		}

		if format == "" {
			for i := 0; i < 3; i++ {
				fmt.Printf("%s\n", result[i])
			}

			for _, v := range list {
				fmt.Printf("%-14s%-16s%-12s%-6s%-8s%-10s%-9s\n", v.NAME, v.IMAGE, v.COMMAND, v.IP, v.PORTS,
					v.STATUS, v.CREATIONTIME)
			}
		} else {
			format = strings.ToUpper(format)
			format = format + "\n"
			t, _ = t.Parse(format)
			for _, c := range list {
				t.Execute(os.Stdout, c)
			}
		}
	}
}
