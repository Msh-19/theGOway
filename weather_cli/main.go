package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

const apiKey = "b3d489fd87d31926206675893e28d4a5" 

type WeatherResponse struct {
    Name    string `json:"name"`
    Main    struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
    Timezone int `json:"timezone"` 
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: weather <city>")
        return
    }

    city := strings.Join(os.Args[1:], " ")
    getWeather(city)
}

func getWeather(city string) {
    encodedCity := url.QueryEscape(city)
    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", encodedCity, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == 401 {
        fmt.Println("Error: Unauthorized. Check your API key.")
        return
    }

    if resp.StatusCode != 200 {
        fmt.Printf("Error: received status code %d\n", resp.StatusCode)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }

    var weatherResponse WeatherResponse
    err = json.Unmarshal(body, &weatherResponse)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Get the current time in the destination's timezone
    location := time.FixedZone("destination", weatherResponse.Timezone)
    currentTime := time.Now().In(location).Format("2006-01-02 15:04:05")

    // Include the current time in the output
    fmt.Printf("Weather report for %s at %s:\n", weatherResponse.Name, currentTime)
    fmt.Printf("Temperature: %.2f°C\n", weatherResponse.Main.Temp)
    fmt.Printf("Description: %s\n", weatherResponse.Weather[0].Description)
}
