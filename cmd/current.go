package cmd

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var city string

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "get the current weather",
	Run: func(cmd *cobra.Command, args []string) {
		const API_KEY = "07604ddd4e0764404025c9e293c62868"
		log.Print("getting weather for the location: ", city)
		client := resty.New()
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, API_KEY)
		resp, err := client.R().Get(url)
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}
		fmt.Println("Response:", resp.String())
	},
}

func init() {
	currentCmd.Flags().StringVarP(&city, "city", "c", "", "City to fetch weather for")
	currentCmd.MarkFlagRequired("city")
}
