package main

import(
	"fmt"
	"crypto/md5" // hash の種類

	"encoding/hex" // hex.Encode... に使う

	"io/ioutil" // for read file
)

func main(){

	data, _ := ioutil.ReadFile("hash.txt")

	fmt.Println("----- 読み込んだファイルの内容 -----")
	fmt.Println(string(data))
	fmt.Println()

	fmt.Println("----- とりあえず data を突っ込んでみる -----")
	b1 := []byte(data)
	md5ForData := md5.Sum(b1)
	fmt.Printf("as %T: %x\n", md5ForData, md5ForData)
	fmt.Printf("as %T: %s\n", hex.EncodeToString(md5ForData[:]), hex.EncodeToString(md5ForData[:]))
	fmt.Println()

	fmt.Println("----- string(data) を突っ込んでみる -----")
	b2 := []byte(string(data))
	md5ForStringData := md5.Sum(b2)
	fmt.Printf("as %T: %x\n", md5ForStringData, md5ForStringData)
	fmt.Printf("as %T: %s\n", hex.EncodeToString(md5ForStringData[:]), hex.EncodeToString(md5ForStringData[:]))
	fmt.Println()
}


/*
----- 読み込んだファイルの内容 -----
なんか適当に書いて
goで読み込む
とりあえず3行くらいで

----- とりあえず data を突っ込んでみる -----
as [16]uint8: 9c1d74e224c0ecb7164955e14c594f7a
as string: 9c1d74e224c0ecb7164955e14c594f7a

----- string(data) を突っ込んでみる -----
as [16]uint8: 9c1d74e224c0ecb7164955e14c594f7a
as string: 9c1d74e224c0ecb7164955e14c594f7a
*/

/*
----- 読み込んだファイルの内容 -----
なんか適当に書いて
goで読み込む
とりあえず3行くらいで
1行増やす

----- とりあえず data を突っ込んでみる -----
as [16]uint8: d5a3c32fe935621e732cd08c87e410f0
as string: d5a3c32fe935621e732cd08c87e410f0

----- string(data) を突っ込んでみる -----
as [16]uint8: d5a3c32fe935621e732cd08c87e410f0
as string: d5a3c32fe935621e732cd08c87e410f0
*/