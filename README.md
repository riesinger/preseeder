# preseeder

preseeder is a simple service serving [Debian Preseed
files](//https://wiki.debian.org/DebianInstaller/Preseed) for use when setting
up your home lab, cloud infrastructure, personal machine or local VM.

It features a simple templating mechanism, so you can serve customized preseeds
via a single HTTP endpoint.

## Installation

Since preseeder is written in Go, you can do a simple `go get
github.com/riesinger/preseeder/cmd/preseeder` to install it locally.
Alternatively, you're able to `git clone` this repository and run `go build
./cmd/preseeder`.

Furthermore, a `Dockerfile` is provided, so you could also build preseeder via
Docker.

## Usage

After starting preseeder, it will listen to HTTP requests on Port `3000`.

You can now request preseeds via a call to `127.0.0.1:3000/preseed/<hostname>`.

The hostname of the machine is the primary means of differentiating between
preseed templates. By default, a preseed template called `_default_` will be
served from `./preseeds`. If a template with the same name as `hostname` is
present, it will be served instead of the default preseed.

Since the templates use the power of [Go
Templates](https://golang.org/pkg/text/template/), you can already configure a
lot of different things using the hostname for differentiation.

### Docker

The Docker container reads preseed templates from the `/data` directory, so you
could use the container something like this:

```sh
docker run --name preseeder -v $(pwd)/preseeds:/data -p 3000:3000 <you docker tag>
```


