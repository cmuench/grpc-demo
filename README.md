# gRPC Demo

Simple Stock-Update gRPC Server.

The demo contains the sources of a gRPC Server written in golang.
There is a PHP client example available in directory "examples/php".

There is a presentation about that example project:

[![](/doc/images/slide-deck.png)][1]

---

## Installation / Build

With GNU make

```bash
make all
```

---

## Server

```
make run-server
```

The server runs on localhost:9000

## Golang Example Client

The golang based client can be started by running:

```bash
make run-client
```

## Requirements

- golang
- composer (for PHP Client or ddev)

---

## PHP Client Project (ddev)

The PHP example project contains a ddev configuration.

To run the project:

```bash
cd examples/php
ddev start
```

The startup process could take some minutes. The required PHP modules (protobuf, grpc) are compiled during startup 
inside the Docker container.

Open the project in your browser.
URL: <https://php-grpc-demo.ddev.site>

---

[1]: https://speakerdeck.com/cmuench/grpc
