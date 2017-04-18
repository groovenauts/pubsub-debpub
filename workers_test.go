package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkersDone(t *testing.T) {
	w1 := &Worker{}
	w2 := &Worker{}
	w3 := &Worker{}

	ws := &Workers{w1, w2, w3}
	assert.Equal(t, false, ws.done())

	w1.done = true
	assert.Equal(t, false, ws.done())
	w2.done = true
	assert.Equal(t, false, ws.done())
	w3.done = true
	assert.Equal(t, true, ws.done())
}

func TestWorkersError(t *testing.T) {
	w1 := &Worker{}
	w2 := &Worker{}
	w3 := &Worker{}

	ws := &Workers{w1, w2, w3}
	assert.NoError(t, ws.error())

	w1.error = fmt.Errorf("foo")
	assert.Error(t, ws.error())
	assert.Equal(t, "foo", ws.error().Error())

	w2.error = fmt.Errorf("bar")
	assert.Error(t, ws.error())
	assert.Equal(t, "foo\nbar", ws.error().Error())
}
