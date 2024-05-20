# colog

`colog` is a simple Go logging library that supports multiple log levels, colored console output, and concurrent-safe file writing.

## Features

- Supports three log levels: Error, Info, and Warn
- Colored console output for easy differentiation of log levels
- Concurrent-safe file writing to prevent log corruption

## Installation

Install the package using `go get`:

```sh
go get github.com/HuXin0817/colog
```

## Usage

### Import the package

```go
import "github.com/HuXin0817/colog"
```

### Initialize the logger

Call `Open` to initialize log files in a specified directory:

```go
err := colog.Open("path/to/log/directory")
if err != nil {
    fmt.Printf("Error opening log files: %v\n", err)
}
```

### Logging messages

Log messages with different levels:

```go
colog.Error("This is an error message")
colog.Errorf("This is an error message with a variable: %v", variable)

colog.Info("This is an info message")
colog.Infof("This is an info message with a variable: %v", variable)

colog.Warn("This is a warning message")
colog.Warnf("This is a warning message with a variable: %v", variable)
```

### Example

```go
package main

import (
    "github.com/HuXin0817/colog"
    "fmt"
)

func main() {
    err := colog.Open("./logs")
    if err != nil {
        fmt.Printf("Error opening log files: %v\n", err)
        return
    }

    colog.Error("This is an error message")
    colog.Errorf("This is an error message with a variable: %d", 42)

    colog.Info("This is an info message")
    colog.Infof("This is an info message with a variable: %s", "example")

    colog.Warn("This is a warning message")
    colog.Warnf("This is a warning message with a variable: %f", 3.14)
}
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### Contact

For any questions or suggestions, feel free to open an issue or contact me directly at <202219120810@stu.cdut.edu.cn>.
