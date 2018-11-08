package svr01cfgs

import (
	"strings"
)

var version string = "1.0"
var servername string = "cfgs"

func join(topic string) string {
	return strings.Join([]string{version, "/", servername, "/", topic}, "")
}
