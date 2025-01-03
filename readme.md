# Raptor

Raptor is a lightweight configuration manager package for Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/codegshinde/raptor.svg)](https://pkg.go.dev/github.com/codegshinde/raptor)

## Installation

To use Raptor in your Go project, you can install it using the following `go get` command:

```bash
go get -u github.com/codegshinde/raptor

```
## Usage
Here's a simple example demonstrating how to use Raptor to manage configurations in your Go application:

```code
package main

import (
	"fmt"
	"github.com/codegshinde/raptor"
)

func main() {
    // Set the configuration file
    err := raptor.SetConfigFile("path/to/your/config/file.txt")
    if err != nil {
        fmt.Println("Error setting config file:", err)
        return
    }

    // Get a value from the configuration
    value, err := raptor.Get("your_key")
    if err != nil {
        fmt.Println("Error getting value:", err)
        return
    }

    fmt.Println("Value:", value)
}


```

## API Reference
Raptor provides the following functions for configuration management:

SetConfigFile(filePath string) error: Sets the configuration file for Raptor.
Get(key string) (string, error): Retrieves the value associated with a key from the configuration.

Set(key, value string): Sets a key-value pair in the configuration.
Delete(key string): Removes a key from the configuration.

GetAll() map[string]string: Returns the entire configuration map.

## Testing
You can run tests for Raptor using the following commands:
```bash
cd tests
go test
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.
