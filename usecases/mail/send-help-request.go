package mail

import (
	"fmt"
	"gps-worker/domain"
)

func SendHelperMail(entrypoint domain.Position, paths <-chan domain.Position) {
	for p := range paths {
		if p.Distance == 63 {
			fmt.Println(p)
			break
		}
	}
}
