package hashencode

import (
	"crypto/sha512"
  "encoding/base64"
	"hash"
)

// Sha512HashEncoder that will be exported from the pkg.
// I set this up in this fashion so that hashencoder could
// be extended for other hashing algorithms in the future
var Sha512HashEncoder = &hashEncoder{
  hashAlgo: 		sha512.New,
  hashAlgoName: "SHA512",
}

// hashEncoder struct for building a sha hash encoder
type hashEncoder struct {
  hashAlgo     func() hash.Hash
  hashAlgoName string
}

// Hash function writes a hash and then returns the base64
// encoding of the hash
func (p *hashEncoder) Hash(password []byte) string {
  shaHash := p.hashAlgo()
	shaHash.Write(password)

	return p.encode(shaHash.Sum(nil))
}

// encode takes byte slice and returns the base64 encoding
func (p *hashEncoder) encode(hash []byte) string {
	encodedHash := base64.StdEncoding.EncodeToString(hash[:])

	return encodedHash
}
