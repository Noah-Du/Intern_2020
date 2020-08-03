package main

import "fmt"

func main{
  cmd := exec.Command("vctl","ps")
  fmt.Println(cmd)
  out, err := cmd.CombinedOutput()
  outString := string(out)
  fmt.Println(outString)
}
