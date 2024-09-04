package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"encoding/base64"
	"github.com/joho/godotenv"
)

type MpesaService struct {
	ConsumerKey    string
	ConsumerSecret string
	ShortCode      string
	PassKey        string
	Environment    string
	BaseURL        string
}

func LoadEnv() error {
	return godotenv.Load(".env")
}

func NewMpesaService() (*MpesaService, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("MPESA_BASE_URL")
	if baseURL == "" {
		baseURL = "https://sandbox.safaricom.co.ke" // Default to sandbox if not set
	}

	return &MpesaService{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		ShortCode:      os.Getenv("MPESA_SHORTCODE"),
		PassKey:        os.Getenv("MPESA_PASSKEY"),
		Environment:    os.Getenv("MPESA_ENV"),
		BaseURL:        baseURL,
	}, nil
}

func (service *MpesaService) GenerateAccessToken() (string, error) {
	url := fmt.Sprintf("%s/oauth/v1/generate?grant_type=client_credentials", service.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(service.ConsumerKey, service.ConsumerSecret)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("unable to fetch access token")
	}

	return token, nil
}

func (service *MpesaService) STKPush(phoneNumber, amount, accountReference, transactionDesc string) (string, error) {
	token, err := service.GenerateAccessToken()
	fmt.Printf(token);
	if err != nil {
		return "", err
	}
	now := time.Now()
	// Format the time to YYYYMMDDHHMMSS
	timestamp := now.Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(service.ShortCode + service.PassKey + timestamp))

	// password := fmt.Sprintf("%s%s%s", service.ShortCode, service.PassKey, timestamp)

	payload := map[string]string{
		"BusinessShortCode": service.ShortCode,
		"Password":         password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            service.ShortCode,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       os.Getenv("CALLBACK_URL"),
		"AccountReference":  accountReference,
		"TransactionDesc":   transactionDesc,
	}

	jsonPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/mpesa/stkpush/v1/processrequest", service.BaseURL), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
