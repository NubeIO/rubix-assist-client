package assist

import (
	"fmt"
	"testing"
	"time"
)

func TestHostLocation(*testing.T) {

	client := New("0.0.0.0", 8080)

	hosts, res := client.GetHostSchema()
	fmt.Println(hosts)
	uuid := ""
	fmt.Println(hosts)
	fmt.Println(res.GetStatus())
	host, res := client.GetLocation(uuid)
	fmt.Println(res.StatusCode)
	if res.StatusCode != 200 {
		//return
	}
	fmt.Println(host)
	host.Name = fmt.Sprintf("name_%d", time.Now().Unix())
	host, res = client.AddLocation(host)
	host.Name = "get fucked_" + fmt.Sprintf("name_%d", time.Now().Unix())
	fmt.Println(res.GetStatus())
	if res.GetStatus() != 200 {

	}

	fmt.Println("NEW host", host.Name)
	host, res = client.UpdateLocation(host.UUID, host)
	if res.GetStatus() != 200 {
		//return
	}
}
