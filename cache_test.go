package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(2)
	c.Set("x", 1)
	c.Set("y", 2)
	c.Get("x")
	c.Get("x")
	c.Get("y")
	c.Set("z", 3)

	xValue, err := c.Get("x")
	assert.Nil(t, xValue, "Expect x to be nil")
	assert.Error(t, err, "Expect error")

	yValue, err := c.Get("y")
	assert.Equal(t, 2, yValue, "Expect y to be 2")
	assert.NoError(t, err, "Expect no error")

	zValue, err := c.Get("z")
	assert.Equal(t, 3, zValue, "Expect z to be 3")
	assert.NoError(t, err, "Expect no error")

	c.Set("z", 4)
	zValue, err = c.Get("z")
	assert.Equal(t, 4, zValue, "Expect z to be 4")
	assert.NoError(t, err, "Expect no error")
}

func TestLFUCache(t *testing.T) {
	c := NewLFUCache(2)
	c.Set("x", 1)
	c.Set("y", 2)
	c.Get("x")
	c.Get("x")
	c.Get("y")
	c.Set("z", 3)

	xValue, err := c.Get("x")
	assert.Equal(t, 1, xValue, "Expect x to be 1")
	assert.NoError(t, err, "Expect no error")

	yValue, err := c.Get("y")
	assert.Nil(t, yValue, "Expect y to be nil")
	assert.Error(t, err, "Expect error")

	zValue, err := c.Get("z")
	assert.Equal(t, 3, zValue, "Expect z to be 3")
	assert.NoError(t, err, "Expect no error")

	c.Set("z", 4)
	zValue, err = c.Get("z")
	assert.Equal(t, 4, zValue, "Expect z to be 4")
	assert.NoError(t, err, "Expect no error")
}
func TestFIFOCache(t *testing.T) {
	c := NewFIFOCache(2)
	c.Set("x", 1)
	c.Set("y", 2)
	c.Get("x")
	c.Get("x")
	c.Get("y")
	c.Get("y")
	c.Set("z", 3)

	xValue, err := c.Get("x")
	assert.Nil(t, xValue, "Expect x to be nil")
	assert.Error(t, err, "Expect error")

	yValue, err := c.Get("y")
	assert.Equal(t, 2, yValue, "Expect y to be 2")
	assert.NoError(t, err, "Expect no error")

	zValue, err := c.Get("z")
	assert.Equal(t, 3, zValue, "Expect z to be 3")
	assert.NoError(t, err, "Expect no error")

	c.Set("z", 4)
	zValue, err = c.Get("z")
	assert.Equal(t, 4, zValue, "Expect z to be 4")
	assert.NoError(t, err, "Expect no error")
}

func TestGetOnEmptyCache(t *testing.T) {
	var value interface{}
	var err error

	lfu := NewLFUCache(2)
	value, err = lfu.Get("1")
	assert.Nil(t, value, "Expect value to be nil")
	assert.Error(t, err, "Expect error")

	lru := NewLRUCache(2)
	value, err = lru.Get("1")
	assert.Nil(t, value, "Expect value to be nil")
	assert.Error(t, err, "Expect error")

	fifo := NewFIFOCache(2)
	value, err = fifo.Get("1")
	assert.Nil(t, value, "Expect value to be nil")
	assert.Error(t, err, "Expect error")
}
