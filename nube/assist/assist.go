package assist

import (
	"fmt"
	"github.com/NubeIO/rubix-assist-client/rest"
	"github.com/NubeIO/rubix-assist-model/model"
)

type Client struct {
	Rest *rest.Service
}

type Path struct {
	Path string
}

var Paths = struct {
	Hosts       Path
	Ping        Path
	Outputs     Path
	OutputsBulk Path
	Inputs      Path
}{
	Hosts:       Path{Path: "/api/hosts"},
	Ping:        Path{Path: "/api/system/ping"},
	Outputs:     Path{Path: "/api/outputs"},
	OutputsBulk: Path{Path: "/api/outputs/bulk"},
	Inputs:      Path{Path: "/api/inputs/all"},
}

// New returns a new instance of the nube common apis
func New(rest *rest.Service) *Client {
	bc := &Client{
		Rest: rest,
	}
	return bc
}

func (inst *Client) GetHosts() (data []*model.Host, response *rest.Reply) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}
