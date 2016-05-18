package main

import (
	"fmt"

	"github.com/hudl/fargo"
)

func main() {
	c := fargo.NewConn("http://192.168.99.100:8080/eureka/v2")

	i := fargo.Instance{
		HostName:         "i-6543",
		Port:             9090,
		App:              "TESTAPP",
		IPAddr:           "127.0.0.10",
		VipAddress:       "127.0.0.10",
		SecureVipAddress: "127.0.0.10",
		DataCenterInfo:   fargo.DataCenterInfo{Name: fargo.MyOwn},
		Status:           fargo.UP,
	}

	c.RegisterInstance(&i)
	f, _ := c.GetApps()

	for key, theApp := range f {
		fmt.Println("App:", key, " First Host Name:", theApp.Instances[0].HostName)
	}

	app, _ := c.GetApp("TESTAPP")
	fmt.Printf("%v\n", app)
}
