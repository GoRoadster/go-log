## go-log

logger wrapper and formatter for logrus

### install

`go get github.com/GoRoadster/go-log`

### usage

```go
import (
	"errors"
	"time"
	
	"github.com/GoRoadster/go-log"
)

func main() {
	path := "./"
	logPrefix := "test-log"
	logLevel := log.TRACE
	shouldSave := true

	err := log.InitLogger(path, logPrefix, logLevel, shouldSave)
	if err != nil {
		panic(err)
	}

	// simple message logging
	log.Info("starting to log")

	// log with additional params
	log.Error("some random failure", "err", errors.New("failed here"), "closing", true)
}
```
