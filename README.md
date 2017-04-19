# pubsub-devpub

[![Build Status](https://secure.travis-ci.org/groovenauts/pubsub-devpub.png)](https://travis-ci.org/groovenauts/pubsub-devpub)


`pubsub-devpub` helps you to publish lots of messages to `Google Cloud Pubsub` topics.
`pubsub-devpub` works concurrently with `--number` option.
`pubsub-devpub` publishes messages for each line of given file like this:

```jsonl
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000001\"]"}}
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000002\"]"}}
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000003\"]"}}
```

Now `pubsub-devpub` supports [JSON Lines](http://jsonlines.org/) only.


## Install

To install cli, simply run:
```
go get github.com/groovenauts/pubsub-devsub
```

Make sure your PATH includes the $GOPATH/bin directory so your commands can be easily used:

```
export PATH=$PATH:$GOPATH/bin
```

## Usage

```bash
$ ./pubsub-devpub --help
NAME:
   pubsub-devpub - github.com/groovenauts/pubsub-devpub

USAGE:
   pubsub-devpub [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --filepath value, -f value  Messages filepath (.jsonl JSON Lines format)
   --number value, -n value    Number of go routine to publish (default: 10)
   --loglevel value, -l value  Log level: debug info warn error fatal panic
   --help, -h                  show help
   --version, -v               print the version
```

## Messages file

### Fields

| Name | Type | Required | Description |
|------|------|----------|-------------|
| topic | string | True | Topic name to publish |
| attributes | map[string]string | True | Attributes of message to publish |
| data  | string | False | Data of message to publish |
| command | []string | False | Command to run after publishing |

#### Command Template

You can use template for `command` field. A keyword within `%{...}` is expanded.

| Keyword | Type | Meaning |
|---------|------|---------|
| topic   | string | Topic name to publish |
| data  | string | Data of message to publish |
| attributes | map[string]string | Attributes of message to publish |
| attrs      | map[string]string | Alias for `attributes`
| message_id | string | Message ID of published message |
| msgId      | string | Alias for message_id |

You can specified the key for keyworkd whose type is map[string]string like this:

```
%{attrs.foo}
```


### Example

The file given must be a `.jsonl` file which has `JSON` format lines.

```json
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"key":"must be string","value":"must be string","download_files":"[\"gs://test-bucket1/path/to/file000001\"]"},"data":"message_t_publish","command":["echo","Message ID: %{msgId}"]}
```

The following text is the same one in pretty format.


```json
{
  "topic":"projects/proj-dummy-999/topics/devpub-target-topic",
  "attributes":{
      "key":"must be string",
      "value":"must be string",
      "download_files":"[\"gs://test-bucket1/path/to/file000001\"]"
  },
  "data":"message_t_publish",
  "command":["echo","Message ID: %{msgId}"]
}
```


## How to build

```
$ make setup
$ make check
$ make build
```


## How to release

```
$ make release
```
