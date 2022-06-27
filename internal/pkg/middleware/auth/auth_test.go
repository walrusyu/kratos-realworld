package auth

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("secret", "ywf")
	spew.Dump(token)
	if err == nil {
		panic(1)
	}
	fmt.Println(token)
}
