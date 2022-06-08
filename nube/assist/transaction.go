package assist

import (
	"fmt"

	"github.com/NubeIO/rubix-assist-model/model"
)

func (inst *Client) GetTransactions() (data []model.Transaction, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetResult(&[]model.Transaction{}).
		Get(path)
	return *resp.Result().(*[]model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) AddTransaction(body *model.Transaction) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf(Paths.Transactions.Path)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Transaction{}).
		Post(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) UpdateTransaction(uuid string, body *model.Transaction) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Transaction{}).
		Patch(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}

func (inst *Client) DeleteTransaction(uuid string, body *model.Transaction) (data *model.Transaction, response *Response) {
	path := fmt.Sprintf("%s/%s", Paths.Transactions.Path, uuid)
	response = &Response{}
	resp, err := inst.Rest.R().
		SetBody(body).
		SetResult(&model.Transaction{}).
		Patch(path)
	return resp.Result().(*model.Transaction), response.buildResponse(resp, err)
}
