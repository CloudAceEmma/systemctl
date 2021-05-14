# systemctl

This library aims at providing idiomatic `systemctl` bindings for go developers, in order to make it easier to write system tooling using golang.

## What is systemctl

systemctl  is a command-line program which grants the user control over the systemd system and service manager.

systemctl may be used to introspect and control the state of the "systemd" system and service manager. Please refer to systemd(1) for an introduction into the basic concepts and functionality this tool manages.

## Supported functions

- [ ] `systemctl is-failed`
- [ ] `systemctl is-active`
- [ ] `systemctl status`
- [ ] `systemctl restart`
- [ ] `systemctl start`
- [ ] `systemctl stop`
- [ ] `systemctl enable`
- [ ] `systemctl disable`
- [ ] `systemctl is-enabled`
- [ ] `systemctl daemon-reload`
- [ ] `systemctl show`

## TODO

- [ ] Add additional bindings for systemctl options I (the author) don't use frequently (or ever) for others to use.
- [ ] Set up `go test` testing

## Simple example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/taigrr/systemctl/v1"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Equivalent to `systemctl enable dhcpd` with a 10 second timeout
    err := systemctl.Enable("dhcpd", ctx)
    if err != nil {
        log.Fatalf("unable to enable unit %s: %v", "dhcpd", err)
    }
}
```

## External resources

- [Official systemctl documentation](https://www.man7.org/linux/man-pages/man1/systemctl.1.html)
- [Inspiration for this repo](https://github.com/Ullaakut/nmap/)
