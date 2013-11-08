package shipyard

import (
    "log"
    "fmt"
)

func LogMessage(prefix string, message ...string) {
    log.SetPrefix(fmt.Sprintf("%s: ", prefix))
    log.Println(message)
}
