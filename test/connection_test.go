package test

import (
	"belajar_golang_dependency_injection/simple"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Databases")
	assert.NotNil(t, connection)

	cleanup()
}
