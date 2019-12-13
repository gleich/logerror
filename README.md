# logerror

This package allows you to log an error to the console in one line.

## Without this package

```go
package main

import (
    "os/exec"
    "github.com/fatih/color"
    log "github.com/sirupsen/logrus"
)

func main() {
    out, err := exec.Command("git", "--version").Output()
    if err != nil {
        color.Red("Printing some information about the command error")
        log.Error(err)
    }
}
```

## With this package

```go
package main

import (
    "os/exec"
    "github.com/Matt-Gleich/logerror"
)

func main() {
    out, err := exec.Command("git", "--version").Output()
    logerror.Log(err, "Printing some information about the command error")
}
```
