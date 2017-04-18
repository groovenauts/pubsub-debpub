# pubsub-devpub

`pubsub-devpub` helps you to publish lots of messages to `Google Cloud Pubsub` topics.
`pubsub-devpub` works concurrently with `--number` option.
`pubsub-devpub` publishes messages for each line of given file like this:

```jsonl
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000001\"]"}}
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000002\"]"}}
{"topic":"projects/proj-dummy-999/topics/devpub-target-topic","attributes":{"download_files":"[\"gs://test-bucket1/path/to/file000003\"]"}}
```

Now `pubsub-debpub` supports [JSON Lines](http://jsonlines.org/) only.


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
   blocks-gcs-proxy - github.com/groovenauts/blocks-gcs-proxy

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
