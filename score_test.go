package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScore(t *testing.T) {
	assert.Equal(t, 8, Score("bling"))
	assert.Equal(t, 23, Score("squiz"))
	assert.Equal(t, 10, Score("money"))
}
