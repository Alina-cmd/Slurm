package main

import (
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

	// fmt.Print(che)
	// fmt.Println()

	ch := make(chan int)

	go che.Run(ch)

	che.Add(gometr.NewGoMetrClient("aurl/lalal/client", 255), ch)
	che.Add(gometr.NewGoMetrClient("burl/lalal/client", 246), ch)
	che.Add(gometr.NewGoMetrClient("durl/lalal/client", 256), ch)
	che.Add(gometr.NewGoMetrClient("eurl/lalal/client", 253), ch)

	time.Sleep(5 * time.Second)

	che.Add(gometr.NewGoMetrClient("curl/lalal/client", 259), ch)

	time.Sleep(45 * time.Second)

}
