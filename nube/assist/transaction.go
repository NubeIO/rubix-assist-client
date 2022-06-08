package assist

import (
	"fmt"

	"github.com/NubeIO/rubix-assist-model/model"
)

func (inst *Client) GetTransactions() (data []model.Message, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Message{}).
		Get(path)
	return *resp.Result().(*[]model.Message), response.buildResponse(resp, err)
}

func (inst *Client) AddTransaction(body *model.Message) (data *model.Message, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Message{}).
		Post(path)
	return resp.Result().(*model.Message), response.buildResponse(resp, err)
}

func (inst *Client) UpdateTransaction(uuid string, body *model.Message) (data *model.Message, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Message{}).
		Patch(path)
	return resp.Result().(*model.Message), response.buildResponse(resp, err)
}

func (inst *Client) DeleteTransaction(uuid string, body *model.Message) (data *model.Message, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Message{}).
		Patch(path)
	return resp.Result().(*model.Message), response.buildResponse(resp, err)
}
