package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var name string
	fmt.Println("请输入要下载的镜像名称及版本")
	fmt.Scanln(&name)
	cmd := exec.Command("vctl", "pull", "docker.mirrors.ustc.edu.cn/library/ubuntu")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
    }
	fmt.Printf("%s\n", string(out))
}
