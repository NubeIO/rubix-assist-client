package assist

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"rubix-assist-client/rest"
	"testing"
)

func TestRubix(*testing.T) {

	restService := &rest.Service{}
	restService.Url = "0.0.0.0"
	restService.Port = 8080
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = false
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = nube.Services.RubixService.Proxy
	restService.NubeProxy = nubeProxy

	client := New(restService)

	_, res := client.GetHosts()
	fmt.Println(res.AsString())
}
