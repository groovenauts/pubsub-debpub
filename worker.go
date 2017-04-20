package main

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	pubsub "google.golang.org/api/pubsub/v1"

	log "github.com/Sirupsen/logrus"
	"github.com/groovenauts/blocks-variable"
)

type Message struct {
	Topic      string            `json:"topic"`
	Data       string            `json:"data,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Command    []string          `json:"command,omitempty"`
}

type Worker struct {
	service *pubsub.Service
	lines   chan string
	done    bool
	error   error
}

func (w *Worker) run() {
	for {
		flds := log.Fields{}
		log.Debugln("Getting a target")
		var line string
		select {
		case line = <-w.lines:
		default: // Do nothing to break
		}
		if line == "" {
			log.Debugln("No target found any more")
			w.done = true
			w.error = nil
			break
		}

		flds["line"] = line
		log.WithFields(flds).Debugln("Job Start")

		err := w.process(line)
		flds["error"] = err
		if err != nil {
			w.done = true
			w.error = err
			break
		}
		log.WithFields(flds).Debugln("Job Finished")
	}
}

func (w *Worker) process(line string) error {
	flds := log.Fields{"line": line}
	log.WithFields(flds).Debugln("Processing line")

	var msg Message
	err := json.Unmarshal([]byte(line), &msg)
	if err != nil {
		flds := log.Fields{"error": err, "line": line}
		log.WithFields(flds).Errorln("JSON parse error")
		return err
	}

	flds["message"] = msg
	log.WithFields(flds).Debugln("Publishing message")

	topic := w.service.Projects.Topics
	call := topic.Publish(msg.Topic, &pubsub.PublishRequest{
		Messages: []*pubsub.PubsubMessage{
			&pubsub.PubsubMessage{
				Attributes: msg.Attributes,
				Data:       base64.StdEncoding.EncodeToString([]byte(msg.Data)),
			},
		},
	})

	res, err := call.Do()
	if err != nil {
		flds["attributes"] = msg.Attributes
		flds["data"] = msg.Data
		flds["error"] = err
		log.WithFields(flds).Errorln("Publish error")
		return err
	}

	flds["MessageIds"] = res.MessageIds
	log.WithFields(flds).Infoln("Publish successfully")

	cmd, err := w.buildCommand(&msg, res)
	if cmd != nil {
		flds["command"] = cmd
		log.WithFields(flds).Debugln("Executing")
		err := cmd.Run()
		if err != nil {
			flds["error"] = err
			log.WithFields(flds).Errorln("Command returned error")
			return err
		}
	}
	log.WithFields(flds).Infoln("Executed successfully")
	return nil
}

func (w *Worker) buildCommand(msg *Message, res *pubsub.PublishResponse) (*exec.Cmd, error) {
	if msg.Command == nil {
		return nil, nil
	}
	if len(msg.Command) == 0 {
		return nil, nil
	}

	d := map[string]interface{}{
		"topic":      msg.Topic,
		"data":       msg.Data,
		"attrs":      msg.Attributes,
		"attributes": msg.Attributes,
		"msgId":      res.MessageIds[0],
		"message_id": res.MessageIds[0],
	}

	v := &bvariable.Variable{Data: d}

	expandeds := []string{}
	for _, t := range msg.Command {
		expanded, err := v.Expand(t)
		if err != nil {
			return nil, err
		}
		vals := strings.Split(expanded, v.Separator)
		for _, val := range vals {
			expandeds = append(expandeds, val)
		}
	}
	cmd := exec.Command(expandeds[0], expandeds[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd, nil
}
