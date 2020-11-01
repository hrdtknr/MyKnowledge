package main
import (
	"fmt"
	"io/ioutil"
)

func main(){
  data, _ := ioutil.ReadFile("sample.txt")
	fmt.Println(string(data))
}

/*
ちなみにfmt.Println(data)だとこんな結果(stringしない)
$ go run read_file.go
[114 101 97 100 32 102 105 108 101 32 116 101 115 116 13 10 233 129 169 229 189 147 227 129 171 230 155 184 227 129 143 13 10 103 104 106 97 100 102 111 103 117 105 106 97 115 13 10 13 10 102 101 114 103 111 105 97 110 111 101 
103 105 110 97 13 10 13 10 102 97 101 111 103 110 97 101 117 105 114 103 110 97 101 117 114 104 103 97 13 10 13 10 103 102 103 110 100 102 111 103 106 110 101 13 10 13 10 114 111 101 116 121 117 114 101 111 104 103 110 98 111 115 105 107 114 13 10 13 10 13 10 98 109 98 107 106 110 111 116 105 110 114 111 116 121 117 105 104 106 111 114 116 115 105 106 13 10 13 10 102 103 106 107 13 10 13 10 100 13 10 13 10 13 10 117 117 116 117 116 117 101 114 116 105 114 116 117 101 114 116 117]
*/