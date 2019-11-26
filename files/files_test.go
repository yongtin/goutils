package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	var check bool
	var err error
	check, err = IsFile("/tmp")
	assert.False(t, check)
	assert.NotNil(t, err)

	check, err = IsFile("/etc/hosts")
	assert.True(t, check)
	assert.Nil(t, err)

	check, err = IsFile("/tpm")
	assert.NotNil(t, err)
	assert.False(t, check)

}

func TestIsDir(t *testing.T) {
	var check bool
	var err error
	check, err = IsDir("/tmp")
	assert.True(t, check)
	assert.Nil(t, err)

	check, err = IsDir("/etc/hosts")
	assert.False(t, check)
	assert.NotNil(t, err)

	check, err = IsDir("/tpm")
	assert.False(t, check)
	assert.NotNil(t, err)

}
