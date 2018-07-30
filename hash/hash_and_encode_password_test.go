package hashencode

import "testing"

// func TestSha512HashPassword(t *testing.T) {
//   cases := []struct {
//     in, want string
//   }{
//     { "angryMonkey", "6441e1581eb9814973755c2d0d002b132c7e2952f3a7f69369168f941cd8448163eaf8c576a11bd10e41f3354a099d2f29b64f664949cf415deecbb603e81fed" },
//   }
//   for _, c:= range cases {
//     got := Sha512HashPassword(c.in)
//     if got != c.want {
//       t.Errorf("HashPassword(%q) == %q, want %q", c.in, got, c.want)
//     }
//   }
// }
//
// func TestBase64EncodeHash(t *testing.T) {
//   cases := []struct {
//     in, want string
//   }{
//     { "angryMonkey",
//       "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" },
//   }
//
//   for _, c := range cases {
//     got := Base64EncodeHash(c.in)
//     if got != c.want {
//       t.Errorf("Base64EncodeHash(%q) == %q, want %q", c.in, got, c.want)
//     }
//   }
// }
//

func TestHasher(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    { "angryMonkey",
      "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" },
    { "password",
      "sQnzu7wkTrgkQZF+0G1hi5AI3Qmzvv0bXgc5THBqi7mAsdd4Xll27ASbRt9fEyavWi6m0QP9B8lThf+rDKy8hg==" },
    { "jumpCloud",
      "m9yBN6qTUXwZprB/UR+Gj7qNJFykiNsDlFaExX235EuMfM3BZAoFgEhIG3IRjSv0HmsLRtbceJz5AMWcwfCnDA==" },
  }

  for _, c := range cases {
    got := Sha512HasherImp.Hash([]byte(c.in))
    if got != c.want {
      t.Errorf("Base64EncodeHash(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}
