package server

import (
	"encoding/json"
	"fmt"
	myerrors "kratos-realworld/internal/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpStruct(t *testing.T) {
	a := &myerrors.HTTPError{
		Errors: struct{ Field map[string][]string }{
			Field: make(map[string][]string),
		},
	}
	b, err := json.Marshal(a)
	assert.NoError(t, err)
	fmt.Printf("%s", string(b))
	panic("ok")
}
