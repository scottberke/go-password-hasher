package hashencode

import (
	"crypto/sha512"
  "encoding/base64"
	"hash"
)

var Sha512HasherImp = &sha512Hasher{
  hashAlgo: sha512.New,
  hashAlgoName: "Sha 512",
}

type hasher interface {
  Hash(password []byte ) string
  encode([]byte) string
}

type sha512Hasher struct {
  hashAlgo     func() hash.Hash
  hashAlgoName string
}

func (p *sha512Hasher) Hash(password []byte) string {
  sha512Hash := p.hashAlgo()
	sha512Hash.Write(password)

	return p.encode(sha512Hash.Sum(nil))
}

func (p *sha512Hasher) encode(hash []byte) string {
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	return encodedHash
}
