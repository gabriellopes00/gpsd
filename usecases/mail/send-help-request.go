package mail

import (
	"fmt"
	"gps-worker/domain"
)

func SendHelperMail(entrypoint domain.Position, channel chan domain.Position) {
	for p := range channel {
		fmt.Println(p)
	}
}
