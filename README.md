# Transdev Client

An API client for the Transdev API, written in Go.

## What is this?

Transdev is the organization that owns the following Dutch bus carriers:

- Connexxion
- Breng
- Hermes
- Bravo

On the websites of these carriers you can look up bus departures. This package
implements a client that uses the same API as these websites, and thus allows
you to retrieve bus departure information from your Go project.

As far as I can tell, all websites make use of the same underlying API and
dataset, e.g. you can look up stops served by Connexxion from the Breng website.
This client uses Breng's domain name for the API calls, but any of the others
should work as well.

## Usage

First, install the package in your project:

```shell
go get github.com/jasperbok/transdev-client
```

Then you can instantiate a new client and call methods on it:

```go
package main

import (
	"fmt"
	"strings"
	transdev "github.com/jasperbok/transdevclient"
)

func main() {
	c := transdev.NewClient()

	// Find information for a stop.
	// At the time of writing 'swartbroek' will yield a single stop.
	results, err := c.Search("swartbroek")
	if err != nil {
		panic(err)
    }

	stopID := results.Stops[0].ID
	// As far as my emperical research can tell, stop and quay IDs can
	// be translated into each other by replacing the letter in their ID.
	quayID := strings.Replace(stopID, ":S:", ":Q:", 1)
	name := results.Stops[0].Name
	town := results.Stops[0].Town

	// With the information retrieved above, we can retrieve upcoming
	// departures for the stop.
	departures, err := c.Departures(quayID, name, town)
	if err != nil {
		panic(err)
    }

	fmt.Printf("%v\n", departures)
}
```

## Known shortcomings

- The `/timetable` endpoint is not implemented.
- The `/getavailablecaptions` endpoint is not implemented.
- The `Search` method only supports looking for stops and routes, but you can
  do other search queries via the same endpoint. These are not implemented.
