package sendkeys

import (
	"runtime"
	"strings"
	"time"
)

func (kb *KBWrap) linuxDelay() {
	if kb.nodelay {
		return
	}
	// For linux, it is very important to wait 2 seconds
	// kayos note: idfk why tho, this is according to keybd_event author
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
}

func compoundErr(errs []error) string {
	var es []string
	for _, e := range errs {
		if e == nil {
			continue
		}
		es = append(es, e.Error())
	}
	return strings.Join(es, ",")
}

func abs(n int) int {
	// ayyee smash 6ros
	n64 := int64(n)
	y := n64 >> 63
	n64 = (n64 ^ y) - y
	return int(n64)
}
