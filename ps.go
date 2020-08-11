package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("vctl","ps")
    out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
    }
    fmt.Printf("%s\n", string(out))
  //fmt.Println(cmd)
  //out, err := cmd.CombinedOutput()
  //outString := string(out)
  //fmt.Println(outString)
}
