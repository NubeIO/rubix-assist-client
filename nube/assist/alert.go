package assist

import (
	"fmt"

	"github.com/NubeIO/rubix-assist-model/model"
)

func (inst *Client) GetAlerts() (data []model.Alert, response *Response) {
	path := fmt.Sprintf(Paths.Alerts.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Alert{}).
		Get(path)
	return *resp.Result().(*[]model.Alert), response.buildResponse(resp, err)
}

func (inst *Client) GetAlert(uuid string) (data *model.Alert, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Alerts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&model.Alert{}).
		Get(path)
	return resp.Result().(*model.Alert), response.buildResponse(resp, err)
}

func (inst *Client) AddAlert(body *model.Alert) (data *model.Alert, response *Response) {
	path := fmt.Sprintf(Paths.Alerts.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Alert{}).
		Post(path)
	return resp.Result().(*model.Alert), response.buildResponse(resp, err)
}

func (inst *Client) UpdateAlert(uuid string, body *model.Alert) (data *model.Alert, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Alerts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Alert{}).
		Patch(path)
	return resp.Result().(*model.Alert), response.buildResponse(resp, err)
}

func (inst *Client) DeleteAlert(uuid string) (response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Alerts.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		Delete(path)
	return response.buildResponse(resp, err)
}
