package common

import (
	"fmt"

	"github.com/chanfun-ren/executor/api"
)

func GenProjectKey(proj *api.Project) string {
	return fmt.Sprintf("%s|%s", proj.NinjaHost, proj.RootDir)
}

func GenCmdKey(proj *api.Project, cmdId string) string {
	return fmt.Sprintf("%s:%s", GenProjectKey(proj), cmdId)
}
