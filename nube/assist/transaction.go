package assist

import (
	"fmt"

	"github.com/NubeIO/rubix-assist-model/model"
)

func (inst *Client) GetTransactions() (data []model.User, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.User{}).
		Get(path)
	return *resp.Result().(*[]model.User), response.buildResponse(resp, err)
}

func (inst *Client) AddTransaction(body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Post(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) UpdateTransaction(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}

func (inst *Client) DeleteTransaction(uuid string, body *model.User) (data *model.User, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.User{}).
		Patch(path)
	return resp.Result().(*model.User), response.buildResponse(resp, err)
}
