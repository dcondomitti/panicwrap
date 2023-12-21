// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris js

package panicwrap

import (
	"os"
	"syscall"
)

var signalsToIgnore = []os.Signal{os.Interrupt, syscall.SIGQUIT}
