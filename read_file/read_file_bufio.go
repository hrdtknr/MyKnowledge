package main

import (
  "fmt"
  "os"
  "bufio"
)

// 1行ずつ読み込むパターン
func main(){
  data, _ :=  os.Open("sample.txt")
  defer data.Close()
  scanner := bufio.NewScanner(data)
  for scanner.Scan(){
		fmt.Println(scanner.Text())
		fmt.Println("read")
  }
}