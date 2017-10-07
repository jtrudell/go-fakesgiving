package viewmodel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Base struct {
	Title    string
	ImageUrl string
}

type GiphyGif struct {
	Id string `json:"id"`
}

type GiphyResponse struct {
	Data []GiphyGif `json:"data"`
}

func NewBase() Base {
	giphy := getGif()
	randIndex := rand.Intn(len(giphy.Data))
	return Base{
		Title:    fmt.Sprintf("Fakesgiving %v", time.Now().Year()),
		ImageUrl: "https://media1.giphy.com/media/" + giphy.Data[randIndex].Id + "/giphy.gif",
	}
}

func getGif() *GiphyResponse {
	giphyApiKey := os.Getenv("GIPHY_API_KEY")
	if giphyApiKey == "" {
		log.Fatalln("No Giphy API key found")
	}

	url := "https://api.giphy.com/v1/gifs/search?api_key=" + giphyApiKey + "&q=thanksgiving&limit=25&offset=0&rating=G&lang=en"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln("Something went wrong:", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Something went wrong:", err)
	}

	var gr = new(GiphyResponse)
	err = json.Unmarshal(body, &gr)
	if err != nil {
		log.Fatalln("Something went wrong:", err)
	}
	return gr
}
