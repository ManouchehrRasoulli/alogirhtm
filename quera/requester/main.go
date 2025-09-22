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

	url := "https://local-api.stage-express.ir/salar-backoffice/order"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZXlKMWMyVnlTV1FpT2pFd01ESTVMQ0oxYzJWeWJtRnRaU0k2SW0wdWNtRnpiM1ZzYVNKOSIsImV4cCI6MTc1ODUzMTUxNn0.YXpWB-Ds-KTd5gMcmOOmLyncxhtO4QMPhNwAD8ShyymFb_snU2fLMmne7ID6UHQ9rLT5C3-BubrRgAXkX8U6BA")

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
						defer res.Body.Close()
						body, err := io.ReadAll(res.Body)
						if err != nil {
							fmt.Println(i, err)
							continue
						}

						if res.StatusCode != http.StatusOK {
							fmt.Println(i, "not ok", string(body))
							continue
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
