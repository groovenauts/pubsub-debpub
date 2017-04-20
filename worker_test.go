package main

import (
	"testing"

	pubsub "google.golang.org/api/pubsub/v1"

	"github.com/stretchr/testify/assert"
)

const (
	topic       = "projects/proj-dummy-999/topics/dummy-topic"
	dummyMsgId1 = "dummy-message-id1"
	dummyMsgId2 = "dummy-message-id2"
)

var (
	attrs = map[string]string{
		"foo": "A",
		"bar": "B",
	}
)

func TestWorkerBuildCommand(t *testing.T) {
	w := &Worker{}

	res1 := &pubsub.PublishResponse{
		MessageIds: []string{dummyMsgId1},
	}

	msg1 := &Message{
		Topic:      topic,
		Attributes: attrs,
		Command:    nil,
	}
	cmd1, err := w.buildCommand(msg1, res1)
	assert.Nil(t, err)
	assert.Nil(t, cmd1)

	msg2 := &Message{
		Topic:      topic,
		Attributes: attrs,
		Command:    []string{},
	}
	cmd2, err := w.buildCommand(msg2, res1)
	assert.Nil(t, err)
	assert.Nil(t, cmd2)

	msg3 := &Message{
		Topic:      topic,
		Attributes: attrs,
		Command:    []string{"echo", "%{msgId}"},
	}
	msg4 := &Message{
		Topic:      topic,
		Attributes: attrs,
		Command:    []string{"echo", "%{message_id}"},
	}
	for _, msg := range []*Message{msg3, msg4} {
		cmd, err := w.buildCommand(msg, res1)
		assert.Nil(t, err)
		assert.NotNil(t, cmd)
		// assert.Equal(t, "echo", cmd3.Path)
		assert.Equal(t, []string{"echo", dummyMsgId1}, cmd.Args)
	}
}
