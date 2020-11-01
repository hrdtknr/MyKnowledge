package main

import(
	"fmt"
	"crypto/md5" // hash の種類
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex" // hex.Encode... に使う
)

func main(){
	s := "password1"
	b := []byte(s) // string => byte

	md5 := md5.Sum(b)
	fmt.Println("----- md5 -----")
	fmt.Printf("as %T: %x\n", md5, md5)
	fmt.Printf("as %T: %s\n", hex.EncodeToString(md5[:]), hex.EncodeToString(md5[:]))

	sha1 := sha1.Sum(b)
	fmt.Println("----- sha1 -----")
	fmt.Printf("as %T: %x\n", sha1, sha1)
	fmt.Printf("as %T: %s\n", hex.EncodeToString(sha1[:]), hex.EncodeToString(sha1[:]))

	sha256 := sha256.Sum256(b) // 他と違う Sum256と書く
	fmt.Println("----- sha256 -----")
	fmt.Printf("as %T: %x\n", sha256, sha256)
	fmt.Printf("as %T: %x\n", hex.EncodeToString(sha256[:]), hex.EncodeToString(sha256[:]))

	sha512 := sha512.Sum512(b) // 256と同様
	fmt.Println("----- sha512 -----")
	fmt.Printf("as %T: %x\n", sha512, sha512)
	fmt.Printf("as %T: %s\n", hex.EncodeToString(sha512[:]),hex.EncodeToString(sha512[:]))
}