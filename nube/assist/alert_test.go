package assist

import (
	"fmt"
	"testing"
	"time"

	"github.com/NubeIO/rubix-assist-model/model"
)

func TestAlert(*testing.T) {
	client := New("0.0.0.0", 8080)

	alertModel := &model.Alert{From: "test_TEST", HostUUID: "hos_B0376FCF5F11", Host: "", AlertType: "system_ping", Count: 34, Date: time.Now()}
	alertModelChanged := &model.Alert{From: "test_CHANGED", HostUUID: "hos_B0376FCF5F11", Host: "", AlertType: "system_ping", Count: 34, Date: time.Now()}

	fmt.Println("\nTesting AddAlert function:")
	aa, r := client.AddAlert(alertModel)
	fmt.Println(aa)
	fmt.Println(r)

	fmt.Println("\nTesting UpdateAlert function:")
	ua, r := client.UpdateAlert(aa.UUID, alertModelChanged)
	fmt.Println(ua)

	fmt.Println("\nTesting GetAlert function:")
	ga, r := client.GetAlert(aa.UUID)
	fmt.Println(ga)

	fmt.Println("\nTesting getAlerts function:")
	gaa, r := client.GetAlerts()
	fmt.Println(gaa)

	fmt.Println("\nTesting DeleteAlert function:")
	da := client.DeleteAlert(aa.UUID)
	fmt.Println(da)

}
