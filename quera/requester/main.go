package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
)

func main() {

	numcpu := runtime.NumCPU()

	url := "https://salar-backoffice.stage-express.ir/cs/backoffice/order"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZXlKMWMyVnlTV1FpT2pFd01ESTVMQ0oxYzJWeWJtRnRaU0k2SW0wdWNtRnpiM1ZzYVNKOSIsImV4cCI6MTc0OTY1MjExOH0.4p904IJ-_bn7qyXOMYiFMs25mvuZQOuzJBYzd4TuKBrgd4NPjt-veARdYBWimv7k1d-DPJjCtDrjFgzzxj3rAQ")

	ctx, cf := context.WithCancel(context.Background())
	for i := 0; i < numcpu; i++ {
		go func(i int) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					{
						res, err := client.Do(req)
						if err != nil {
							fmt.Println(i, err)
							return
						}
						if res.StatusCode != http.StatusOK {
							fmt.Println(i, "not ok")
							continue
						}

						defer res.Body.Close()
						body, err := io.ReadAll(res.Body)
						if err != nil {
							fmt.Println(err)
							return
						}
						fmt.Println(string(body))
						cf()
						return
					}
				}
			}
		}(i)
	}

	<-ctx.Done()
}
