package main

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/kenbyte/Anzar"
)

var (
	ecmwfapiUrl = "https://api.open-meteo.com/v1"
	hrrrapiUrl  = "https://api.open-meteo.com/v1"
	metarapiURL = "https://aviationweather.gov/api"

	ecmwfClient *Anzar.Client
	hrrrClient  *Anzar.Client
	metarClient *Anzar.Client

	ecmwfErr error
	hrrrErr  error
	metarErr error
)

func initECMWFClient() {
	ecmwfClient, ecmwfErr = Anzar.NewClient(ecmwfapiUrl, false)
	if ecmwfErr != nil {
		fmt.Printf("Error creating ECMWF client: %v\n", ecmwfErr)
	}
}

func initHRRRClient() {
	hrrrClient, hrrrErr = Anzar.NewClient(hrrrapiUrl, false)
	if hrrrErr != nil {
		fmt.Printf("Error creating HRRR client: %v\n", hrrrErr)
	}
}

func initMETARClient() {
	metarClient, metarErr = Anzar.NewClient(metarapiURL, false)
	if ecmwfErr != nil {
		fmt.Printf("Error creating METAR client: %v\n", metarErr)
	}
}

func buildForecastQuery(lat, lon float64, unit, days, model string) string {
	q := url.Values{}

	q.Set("latitude", fmt.Sprintf("%f", lat))
	q.Set("longitude", fmt.Sprintf("%f", lon))
	q.Set("daily", "temperature_2m_max")
	q.Set("temperature_unit", unit)
	q.Set("forecast_days", days)
	q.Set("models", model)

	return "/forecast?" + q.Encode()
}
func buildMetarQuery(station string) string {
	q := url.Values{}
	q.Set("ids", station)
	q.Set("format", "json")

	return "/data/metar?" + q.Encode()
}

func (city *City) getECMWF() {
	initECMWFClient()
	if ecmwfErr != nil {
		fmt.Println("ECMWF Client initialization failed.")
		return
	}

	requestPath := buildForecastQuery(
		city.Lat,
		city.Long,
		"celsius",
		"2",
		"ecmwf_ifs025",
	)

	ctx := context.Background()

	resp, err := ecmwfClient.Get(ctx, requestPath)
	if err != nil {
		fmt.Printf("Error making request to ECMWF API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}

func (city *City) getHRRR() {
	initHRRRClient()
	if hrrrErr != nil {
		fmt.Println("HRRR Client initialization failed.")
		return
	}

	requestPath := buildForecastQuery(
		city.Lat,
		city.Long,
		"fahrenheit",
		"1",
		"gfs_seamless",
	)

	ctx := context.Background()

	resp, err := hrrrClient.Get(ctx, requestPath)
	if err != nil {
		fmt.Printf("Error making request to HRRR API: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	dallas.getHRRR()
}
