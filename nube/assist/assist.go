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

func (inst *Client) GetHost(uuid string) (data *model.Host, response *rest.Reply) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) AddHost(body *model.Host) (data *model.Host, response *rest.Reply) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) UpdateHost(uuid string, body *model.Host) (data *model.Host, response *rest.Reply) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

type Message struct {
	Message string `json:"message"`
}

func (inst *Client) DeleteHost(uuid string) (data Message, response *rest.Reply) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.DELETE).
		SetPath(path).
		DoRequest()
	data = Message{
		Message: "delete error",
	}
	if req.GetStatus() == 200 {
		data.Message = "delete ok"
	}
	response = inst.Rest.RestResponse(req, nil)
	return
}
