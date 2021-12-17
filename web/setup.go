package web

import (
	"context"
	"net/http"
	"sync"

	"github.com/Tackem-org/Global/remoteWebSystem"
	"github.com/Tackem-org/User/static"
)

var (
	httpServer *http.Server
)

func Setup() {
	remoteWebSystem.Setup(&static.FS)
	remoteWebSystem.AddPath("/", RootPage)
	remoteWebSystem.AddAdminPath("/", AdminRootPage)
	remoteWebSystem.AddPath("{{number:userid}}", UserIDPage)
	remoteWebSystem.AddPath("{{string:username}}", UserNamePage)
}

func Shutdown(wg *sync.WaitGroup) {
	if err := httpServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	httpServer = nil
}
