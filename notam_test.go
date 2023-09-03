package notam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetch(t *testing.T) {
	items, err := Fetch("EPRZ")
	assert.Nil(t, err)
	assert.NotEmpty(t, items)
}

func TestFetch_InvalidIcao(t *testing.T) {
	items, err := Fetch("INVALID_CODE")
	assert.NotNil(t, err)
	assert.Nil(t, items)
}
