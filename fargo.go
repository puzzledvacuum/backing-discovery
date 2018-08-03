package main

import (
	"fmt"

	"github.com/hudl/fargo"
)

func main() {
	c := fargo.NewConn("http://localhost:8081/eureka/v2")

	// i := fargo.Instance{
	// 	HostName:         "i-6543",
	// 	Port:             9091,
	// 	App:              "TESTAPP",
	// 	IPAddr:           "127.0.0.1",
	// 	VipAddress:       "127.0.0.1",
	// 	SecureVipAddress: "127.0.0.1",
	// 	DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
	// 	Status:           fargo.UP,
	// }
	//
	// c.RegisterInstance(&i)
	f, _ := c.GetApps()

	for key, theApp := range f {
		fmt.Println("App:", key, " First Host Name:", theApp.Instances[0].HostName)
	}

	app, _ := c.GetApp("backing-catalog")
	fmt.Printf("%v\n", app)
	app, _ = c.GetApp("backing-fulfillment")
	fmt.Printf("%v\n", app.HostName)
}
