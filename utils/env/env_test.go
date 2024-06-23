package env_test

import (
	"itmx-test/utils/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	env.CheckDotEnv()
	ret := env.MustGet("PORT")
	assert.Equal(t, "", ret)
}
