package mail

import (
	"fmt"
	d "gps-worker/domain"
)

func SendHelperMail(paths <-chan *d.Position) {
	for p := range paths {
		fmt.Println(p)
	}
}
