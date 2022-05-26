package assist

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"time"

	"github.com/NubeIO/rubix-assist-client/rest"
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

	hosts, res := client.GetHosts()
	uuid := ""
	for _, host := range hosts {
		uuid = host.UUID
	}
	if uuid == "" {
		return
	}
	host, res := client.GetHost(uuid)
	fmt.Println(res.GetStatus())
	if res.GetStatus() != 200 {
		return
	}
	host.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	host, res = client.AddHost(host)
	host.Name = "get fucked_" + fmt.Sprintf("name_%d", time.Now().Unix())
	if res.GetStatus() != 200 {
		return
	}
	host, res = client.UpdateHost(host.UUID, host)
	if res.GetStatus() != 200 {
		return
	}
	fmt.Println(host.Name)
	fmt.Println(res.GetStatus())
	_, res = client.DeleteHost(host.UUID)
	if res.GetStatus() != 200 {
		return
	}
	fmt.Println(res.GetStatus())
}
