package longrunning

import (
	"context"
	"log"
	"time"
)

func Run(ctx context.Context, d time.Duration) {
	isCanceled := false
	log.Print("Long running job starts")
	defer func() {
		if isCanceled {
			log.Print("Long running job is cancelled")
			return
		}

		log.Print("Long running job ends")
	}()

	select {
	case <-ctx.Done():
		isCanceled = true
	case <-time.After(d):
	}
}
