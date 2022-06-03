package main

import (
	"fmt"
	"hw_m4_03/gometr"
	"time"
)

func main() {
	// client := new(gometr.GoMetrClient)
	// slice := client.GetHealth()
	// for _, val := range slice {
	// 	if val.Status == gometr.PassStatus {
	// 		fmt.Println(val.ServiceID)
	// 	}
	// }

	che := new(gometr.Checker)

	fmt.Print(che)
	fmt.Println()

	che.Add(gometr.NewGoMetrClient("aurl/lalal/client", 255))
	che.Add(gometr.NewGoMetrClient("burl/lalal/client", 246))
	che.Add(gometr.NewGoMetrClient("durl/lalal/client", 256))
	che.Add(gometr.NewGoMetrClient("eurl/lalal/client", 253))

	go che.Run()

	time.Sleep(5 * time.Second)

	che.Add(gometr.NewGoMetrClient("curl/lalal/client", 259))

	time.Sleep(45 * time.Second)

}
