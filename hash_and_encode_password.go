package hashencode
// package main

import (
	// "fmt"
	"crypto/sha512"
  // "encoding/hex"
  "encoding/base64"
	"hash"
)

//
// func Sha512HashPassword(password string) string {
//   sha512Hash := sha512.New()
//   sha512Hash.Write([]byte(password))
//   passwordHash := hex.EncodeToString(sha512Hash.Sum(nil)[:])
//
//   return passwordHash
// }
//
// func Base64EncodeHash(password string) string {
//   // base64String := base64.URLEncoding.EncodeToString([]byte(hash))
//   sha512Hash := sha512.New()
//   sha512Hash.Write([]byte(password))
//   base64String := base64.StdEncoding.EncodeToString(sha512Hash.Sum(nil)[:])
//
//   return base64String
// }

// Other imlementation testing


var sha512HasherImp = &sha512Hasher{
  hashAlgo: sha512.New,
  hashAlgoName: "Sha 512",
}


type hasher interface {
  hash(password []byte ) string

  encode([]byte) string
}

type sha512Hasher struct {
  hashAlgo     func() hash.Hash
  hashAlgoName string
}

func (p *sha512Hasher) hash(password []byte) string {
  sha512Hash := p.hashAlgo()
	sha512Hash.Write(password)

	return p.encode(sha512Hash.Sum(nil))
}

func (p *sha512Hasher) encode(hash []byte) string {
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	return encodedHash
}

//
// func main() {
//   hash := HashPassword("angryMonkey")
//   fmt.Println(hash)
// }
