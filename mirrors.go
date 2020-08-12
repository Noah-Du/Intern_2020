package main

import (
	"fmt"
	"os/exec"
)

// 1. 如何设置成多镜像线路可自己切换？

func main() {
	var name string
	fmt.Println("请输入要下载的镜像名称及版本")
	fmt.Scanln(&name)
	name = "hub-mirror.c.163.com/library/" + name
	fmt.Printf("%v", name)
	cmd := exec.Command("vctl", "pull", name)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
    }
	fmt.Printf("%s\n", string(out))
}
