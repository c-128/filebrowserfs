# FilebrowserFS
A simple Go library to interact with [File Browser's](https://filebrowser.org/) HTTP API.

## Installation
```sh
go get github.com/c-128/filebrowserfs
```

## Examples
```go
package main

import (
	"github.com/c-128/filebrowserfs"
)

func main() {
  client, err := filebrowserfs.Login("http://localhost/", "admin", "securepassword")
	if err != nil {
		return err
	}
  // Create a new file
  err = client.CreateFile("hello_world.txt", false)
  if err != nil {
	  panic(err)
  }

  // Write to the new file
  err = client.WriteFile(
	  "hello_world.txt",
	  bytes.NewBufferString("Hello world!"),
  )
  if err != nil {
	  panic(err)
  }

  // Create a new directory
  err = client.CreateDirectory("hello_directory", false)
  if err != nil {
	  panic(err)
  }

  // Delete the directory
  err = client.Remove("hello_directory")
  if err != nil {
	  panic(err)
  }
}
```
