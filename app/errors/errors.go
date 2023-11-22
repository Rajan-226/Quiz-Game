package errors

import (
	"fmt"
	"os"
)

func ExitWithMessage(error string) {
	fmt.Println(error)
	os.Exit(1)
}
