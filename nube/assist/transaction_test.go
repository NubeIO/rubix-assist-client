package assist

import (
	"fmt"
	"testing"

	"github.com/NubeIO/rubix-assist-model/model"
)

func TestTransaction(*testing.T) {
	client := New("0.0.0.0", 8080)

	TransactionModel := &model.Message{}
	TransactionModelChanged := &model.Message{}

	fmt.Println("\nTesting AddTransaction function:")
	at, r := client.AddTransaction(TransactionModel)
	fmt.Println(at)
	fmt.Println(r)

	fmt.Println("\nTesting UpdateTransaction function:")
	ut, r := client.UpdateTransaction(at.UUID, TransactionModelChanged)
	fmt.Println(ut)

	fmt.Println("\nTesting GetTransaction function:")
	gt, r := client.GetTransaction(at.UUID)
	fmt.Println(gt)

	fmt.Println("\nTesting GetTransactions function:")
	gaa, r := client.GetTransactions()
	fmt.Println(gaa)

	fmt.Println("\nTesting DeleteAlert function:")
	da := client.DeleteAlert(at.UUID)
	fmt.Println(da)

}
