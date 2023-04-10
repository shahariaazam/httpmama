package gopkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnSomeString(t *testing.T) {
	assert.NotEmpty(t, ReturnSomeString())
}
