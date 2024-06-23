package router_test

import (
	"itmx-test/mocks"
	router "itmx-test/router/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	service := mocks.CustomerServiceMock{}
	router.NewHTTPHandler(&service)
	assert.Equal(t, "", "")

}
