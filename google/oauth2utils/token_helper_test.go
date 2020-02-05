package oauth2utils

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOAuth2GetCodeFromLocalServer(t *testing.T) {
	ctx := context.Background()

	code, err := getCodeFromLocalServer(ctx)
	fmt.Println(code)
	assert.NotNil(t, err)
}
