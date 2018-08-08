package hashencode

import "testing"

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
    got := Sha512HashEncoder.Hash([]byte(c.in))
    if got != c.want {
      t.Errorf("Base64EncodeHash(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}
