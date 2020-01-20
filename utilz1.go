

package utilz

import (
	"fmt"
	"os"
	//"errors"
)
// Check - check err return; print err msg; and exit if necessary
func Check(e error ,es string,ex bool) {
	if e !=nil {
		fmt.Printf ("\nerror: %s  %v\n\n",es,e)
		if ex {
			os.Exit(1)
		}
	}
}