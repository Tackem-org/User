package web

import (
	"github.com/Tackem-org/Global/remoteWebSystem"
	"github.com/Tackem-org/User/static"
)

func Setup() {
	remoteWebSystem.Setup(&static.FS)
	remoteWebSystem.AddPath("/", RootPage)
	remoteWebSystem.AddAdminPath("/", AdminRootPage)
	remoteWebSystem.AddPath("{{number:userid}}", UserIDPage)
	remoteWebSystem.AddPath("{{string:username}}", UserNamePage)
}
