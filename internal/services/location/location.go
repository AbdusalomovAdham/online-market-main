package location

import (
	"fmt"
	"io"
	"net/http"
)

func Location(driverLat, driverLon, clientLat, clientLon float64) (string, error) {
	apiKey := "d0d5115da5e845ffbd366be746e8dfca"

	url := fmt.Sprintf(
		"https://api.geoapify.com/v1/routing?waypoints=%f,%f|%f,%f&mode=drive&apiKey=%s",
		driverLat, driverLon, clientLat, clientLon, apiKey,
	)

	fmt.Println("url", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
