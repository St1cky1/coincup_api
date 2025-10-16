package coincap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can`t be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil

}

func (c Client) GetAssets() ([]Asset, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("file .env not find")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("Api key not find")
	}

	url := "https://rest.coincap.io/v3/assets/"

	// создание запроса
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+apiKey) // передача обязательных хедеров
	// отправка запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	// Распарсиваем json ответ в структуры
	var r assetsResponse

	err = json.Unmarshal(body, &r) // распарсиваем слайсл байт в структуру
	if err != nil {
		return nil, err
	}
	// fmt.Println(r.Assets.Info()) // для вывода одной криптовалюты по id
	return r.Assets, nil
}

func (c Client) GetAsset(name string) (Asset, error) {

	 err := godotenv.Load()
    if err != nil {
        return Asset{}, errors.New("file .env not find")
    }

    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        return Asset{}, errors.New("API key not found")
    }

	url := fmt.Sprintf("https://rest.coincap.io/v3/assets/%s", name)

	// создание запроса
	req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("Authorization", "Bearer "+apiKey)
	// отправка запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return Asset{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Asset{}, err
	}
	// fmt.Println(string(body))
	// Распарсиваем json ответ в структуры
	var r assetResponse

	err = json.Unmarshal(body, &r) // распарсиваем слайсл байт в структуру
	if err != nil {
		return Asset{}, err
	}
	// fmt.Println(r.Assets.Info()) // для вывода одной криптовалюты по id
	return r.Asset, nil
}
