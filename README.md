# Furdarius\Contenttype
[![Build Status](https://travis-ci.org/Furdarius/contenttype.svg?branch=master)](https://travis-ci.org/Furdarius/contenttype) [![Coverage Status](https://coveralls.io/repos/github/Furdarius/contenttype/badge.svg?branch=master)](https://coveralls.io/github/Furdarius/contenttype?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/furdarius/contenttype)](https://goreportcard.com/report/github.com/furdarius/contenttype)

Library provides set of middlewares for setting `Content-Type` header.

## Installation

With a properly configured Go toolchain:
```sh
go get -u github.com/furdarius/contenttype
```

## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/furdarius/contenttype"
)

func main() {
	// Preferred router initialization
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	handler := contenttype.Middleware(mux, "application/json")
	// Also you can use predefined contenttype middlewares.
	// Availible middlewares: JSON, HTML, XML, PDF, Text, OctetStream
	//
	// You can change middleware above with:
	// handler := contenttype.JSON(mux)

	http.ListenAndServe(":3000", handler)
}
```

#### Availible middlewares with predefined type:

##### JSON

JSON middleware set header `Content-Type: application/json`

```go
h := contenttype.JSON(mux)
```

##### HTML

HTML middleware set header `Content-Type: text/html`

```go
h := contenttype.HTML(mux)
```

##### XML

XML middleware set header `Content-Type: application/xml`

```go
h := contenttype.XML(mux)
```

##### PDF

PDF middleware set header `Content-Type: application/pdf`

```go
h := contenttype.PDF(mux)
```

##### Text

Text middleware set header `Content-Type: text/plain`

```go
h := contenttype.Text(mux)
```

##### OctetStream

OctetStream middleware set header `Content-Type: application/octet-stream`

```go
h := contenttype.OctetStream(mux)
```