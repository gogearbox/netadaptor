<p align="center">
	<a href="https://gogearbox.com">
    	<img src="https://raw.githubusercontent.com/gogearbox/gearbox/master/assets/gearbox-512.png"/>
	</a>
    <br />
    <a href="https://godoc.org/github.com/gogearbox/netadaptor">
      <img src="https://godoc.org/github.com/gogearbox/netadaptor?status.png" />
    </a>
    <img src="https://github.com/gogearbox/netadaptor/workflows/Test%20&%20Build/badge.svg?branch=master" />
    <a href="https://goreportcard.com/report/github.com/gogearbox/gearbox">
      <img src="https://goreportcard.com/badge/github.com/gogearbox/netadaptor" />
    </a>
	<a href="https://discord.com/invite/CT8my4R">
      <img src="https://img.shields.io/discord/716724372642988064?label=Discord&logo=discord">
  	</a>
    <a href="https://deepsource.io/gh/gogearbox/netadaptor/?ref=repository-badge" target="_blank">
      <img alt="DeepSource" title="DeepSource" src="https://static.deepsource.io/deepsource-badge-light-mini.svg">
    </a>
</p>

**netadaptor** converts net/http handlers to gearbox :gear: handlers


### Supported Go versions & installation

:gear: gearbox requires version `1.14` or higher of Go ([Download Go](https://golang.org/dl/))

Just use [go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) to download and install gearbox

```bash
go get -u github.com/gogearbox/gearbox
go get u- github.com/gogearbox/netadaptor
```


### Examples

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gogearbox/gearbox"
    "github.com/gogearbox/netadaptor"
)

func main() {
	// Setup gearbox
	gb := gearbox.New()

	// Define your handlers
	gb.Get("/hello", netadaptor.HTTPHandlerFunc(myHandler))

	// Start service
	gb.Start(":3000")
}

// http handler function
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
```
