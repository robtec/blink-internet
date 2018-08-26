# Blink Internet

A simple utility to monitor internet connectivity using [Blink1](https://blink1.thingm.com/)

## Setup

```go
go get github.com/robtec/blink-internet
```

USB Liberies are also required
* libusb
* libusb-compat

Consult your operating systems package manager and Google to install

## Usage

```go
$ blink-internet
# monitoring internet connectivity #
```

## Colours

* Blue - loading (first ping)
* Green - connected to the internet (google.com)
* Amber - connection timeout
* Red - no connection
