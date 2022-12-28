package builtins

import (
	"L2/develop/dev08/internal"
	"log"
	"os"
	"sync"
	"syscall"
)

type Exec struct {
	internal.Command
}

func (e *Exec) Run(wg *sync.WaitGroup) error {
	e.CloseFds()
	if len(e.GetArgs()) > 1 {
		if err := syscall.Exec(e.GetArgs()[1], e.GetArgs()[1:], os.Environ()); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
