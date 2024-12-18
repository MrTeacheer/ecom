package errs

import (
	"log"
)

func Check(err error) {
	switch {
	case err != nil:
		log.Fatal(err)
	}
}
