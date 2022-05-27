package assist

import (
	"fmt"
	"github.com/NubeIO/rubix-assist-client/rest"
	"github.com/NubeIO/rubix-assist-model/model"
)

func (inst *Client) GetHostNetworks() (data []*model.Network, response *rest.Reply) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) GetHostNetwork(uuid string) (data *model.Network, response *rest.Reply) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) AddHostNetwork(body *model.Network) (data *model.Network, response *rest.Reply) {
	path := fmt.Sprintf(Paths.Hosts.Path)
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) UpdateHostNetwork(uuid string, body *model.Network) (data *model.Network, response *rest.Reply) {
	path := fmt.Sprintf("%s/%s", Paths.Hosts.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

func (inst *Client) DeleteHostNetwork(uuid string) (data Message, response *rest.Reply) {
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
