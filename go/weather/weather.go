/*
We need to know the wind speed to determine if we should turn on our wind turbine or not.
But if it's too hot, we don't want to turn it on because heat is bad for the turbine.

What you need to do is fetch the current weather forecast and check the wind and the temperature.
If the average wind speed of the last 5 readings received is more than 5 mph,
and current temperature is less than 15 degrees we want to turn on the wind turbine, otherwise we turn the turbine off.

Because the weather can change quickly, we want to fetch the current forecast every second
and make sure that happens in a deterministic way.

Example response:
{
	"current": {
		"condition": {
			"text": "Overcast"
			},
		"is_day": 1,
		"last_updated": "2022-07-07 10:45",
		"last_updated_epoch": 1657187100,
		"precip_in": 0.0,
		"precip_mm": 0.0,
		"temp_c": 19.0,
		"uv": 5.0,
		"wind_degree": 350,
		"wind_dir": "N",
		"wind_mph": 6.9
	}
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type current struct {
	Temp float32 `json:"temp_c"`
	Wind float32 `json:"wind_mph"`
}

type response struct {
	Cur current `json:"current"`
}

func readFromService(x chan current) {
	url := newMockServer()

	for i := 0; i < 20; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		val, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		var response response
		//fmt.Println(string(val))
		json.Unmarshal(val, &response)
		x <- response.Cur
		//fmt.Println(response)
		time.Sleep(1 * time.Second)
	}
	close(x)
}

//go:noinline
func computeResults(data chan current, done chan struct{}) {
	var speeds []float32
	var total float32
	for {
		val, more := <-data
		if more {
			total = 0
			if len(speeds) > 4 {
				for _, val := range speeds {
					total += val
				}
				if total/5 > 5 && val.Temp > 15 {

					fmt.Println("start the turbine..", " average ", total/5, " temp :", val.Temp)
				} else {
					fmt.Println("Stop the turbine..", " average ", total/5, " temp :", val.Temp)
				}
				speeds = speeds[1:]
			} else {
				speeds = append(speeds, val.Wind)
			}
			fmt.Println(speeds)
		} else {
			done <- struct{}{}
			break
		}
	}
}

// example of using an empty struct in channel
// empty struct are used as size of empty struct is 0 and this is moemory efficient
func main() {
	data := make(chan current)
	done := make(chan struct{})
	go readFromService(data)
	go computeResults(data, done)
	<-done
}
