// Package utilz - package for utilities

package utilz

import (
	"fmt"
	"log"
	"runtime"
	"bytes"
)

var (
	buf      bytes.Buffer
	logger   = log.New(&buf, "logger: ", log.Lshortfile)
	stackBuf []byte
)

// Check - check err return; print err msg; and exit if necessary
func  Check(e error ,es string,ex bool) {
	if e != nil && ex {
		//log.Fatalf("%s  %s", err, errStr)
		// os.Exit(1) will be called here
		panic(es)
	} else if e != nil && !ex {
		runtime.Stack(stackBuf, true)
		fmt.Printf("stack: %s\n", stackBuf)
		log.Printf("non fatal %s %s", e, es)
	}
}

