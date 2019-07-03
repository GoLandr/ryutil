package ryutil



import (
	"fmt"
	"os"
	"syscall"
)

type SignalHandler func(s os.Signal, arg interface{})

type SignalSet struct {
	m map[os.Signal]SignalHandler
}

func NewSignalSet() *SignalSet {
	ss := new(SignalSet)
	ss.m = make(map[os.Signal]SignalHandler)
	return ss
}

func (set *SignalSet) Register(s os.Signal, handler SignalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *SignalSet) Handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}
	panic("won't reach here")
}


var stopFlag_ bool =false


func SigHandlerFunc(s os.Signal, arg interface{}) {
	switch s {
	case syscall.SIGUSR1: // check
		fmt.Printf("stopping Status : %v\n", stopFlag_)
	case syscall.SIGUSR2: // run
		formerFlag := stopFlag_
		stopFlag_ = false
		fmt.Printf("stopping Status changed from %v to %v\n", formerFlag, stopFlag_)
	case syscall.SIGQUIT: // stop
		formerFlag := stopFlag_
		stopFlag_ = true
		fmt.Printf("stopping Status changed from %v to %v\n", formerFlag, stopFlag_)

	case syscall.SIGSEGV:
		fmt.Printf("SIGSEGV Status : %v\n", stopFlag_)
	}
}