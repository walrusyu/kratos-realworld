package biz

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pwd, err := hashPassword("123")
	fmt.Println(pwd)
	if err == nil {
		spew.Dump(pwd)
	}
}

func TestVerifyPassword(t *testing.T) {
	pwd := "123"
	hashed := "$2a$10$KL7XfuCEMUrAuzL06kzecuxzmc/AAj1aH97y/fy1A0QgIp3tfmVOu"
	success := verifyPassword(hashed, pwd)
	a := assert.New(t)
	a.True(success)
}
