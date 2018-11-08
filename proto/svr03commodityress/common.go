package svr03commodityress

import (
	"strings"
)

var version string = "1.0"
var servername string = "commodityress"

func join(topic string) string {
	return strings.Join([]string{version, "/", servername, "/", topic}, "")
}
