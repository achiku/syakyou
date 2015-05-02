package version

import "fmt"

var VERSION string

var GITCOMMIT string

func UserAgent() string {
	return fmt.Sprintf("mackerel-agent/%s (revision %s)", VERSION, GITCOMMIT)
}
