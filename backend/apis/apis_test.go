package apis

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	assert.NotNil(t, NewContext(http.StatusOK, "OK", nil))
}
