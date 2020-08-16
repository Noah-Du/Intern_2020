package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	//	log "github.com/sirupsen/logrus"
)

func main() {
	var name string
	var accelerator string

	flag.StringVar(&name, "i", "", "镜像，默认为空")
	flag.StringVar(&accelerator, "a", "hub-mirror.c.163.com/library/", "加速器url，默认为网易云加速器")
	flag.Parse()

	download := accelerator + name
	fmt.Printf("下载路径为%v", download)
	cmd := exec.Command("vctl", "pull", download)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
}
