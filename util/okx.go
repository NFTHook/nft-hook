package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"nfthook/config"
	"time"
)

func CreateSignature(timestamp, method, requestPath, queryString, body, secretKey string) string {
	message := timestamp + method + requestPath + queryString + body
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GetTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func OkxSendGetRequest(requestPath string, params map[string]string) (string, error) {
	queryString := "?"
	for k, v := range params {
		queryString += url.QueryEscape(k) + "=" + url.QueryEscape(v) + "&"
	}
	queryString = queryString[:len(queryString)-1]

	timestamp := GetTimestamp()
	signature := CreateSignature(timestamp, "GET", requestPath, queryString, "", config.Get().Okx.ApiSecret)

	client := &http.Client{}
	fullURL := config.Get().Okx.BaseURL + requestPath + queryString
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("OK-ACCESS-KEY", config.Get().Okx.ApiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", config.Get().Okx.Passphrase)
	// req.Header.Set("OK-ACCESS-PROJECT", apiProject) // 仅适用于 WaaS APIs

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	return string(respBody), nil
}

func OkxSendPostRequest(requestPath string, params map[string]string) (string, error) {
	timestamp := GetTimestamp()
	body, _ := json.Marshal(params)
	signature := CreateSignature(timestamp, "POST", requestPath, "", string(body), config.Get().Okx.ApiSecret)

	client := &http.Client{}
	fullURL := config.Get().Okx.BaseURL + requestPath
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", nil
	}

	req.Header.Set("OK-ACCESS-KEY", config.Get().Okx.ApiKey)
	req.Header.Set("OK-ACCESS-SIGN", signature)
	req.Header.Set("OK-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("OK-ACCESS-PASSPHRASE", config.Get().Okx.Passphrase)
	// req.Header.Set("OK-ACCESS-PROJECT", apiProject) // 仅适用于 WaaS APIs
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", nil
	}

	fmt.Println("Response:", string(respBody))
	return string(respBody), nil
}

type OkxNftAssetDetail struct {
	Code int `json:"code"`
	Data struct {
		AnimationUrl  string `json:"animationUrl"`
		AssetContract struct {
			Chain           string `json:"chain"`
			ContractAddress string `json:"contractAddress"`
			Erc2981         bool   `json:"erc2981"`
			OwnerAddress    string `json:"ownerAddress"`
			TokenStandard   string `json:"tokenStandard"`
		} `json:"assetContract"`
		Attributes string `json:"attributes"`
		Collection struct {
			AssetContracts []struct {
				Chain           string `json:"chain"`
				ContractAddress string `json:"contractAddress"`
				Erc2981         bool   `json:"erc2981"`
				OwnerAddress    string `json:"ownerAddress"`
				TokenStandard   string `json:"tokenStandard"`
			} `json:"assetContracts"`
			BackgroundImage string        `json:"backgroundImage"`
			CategoryList    []interface{} `json:"categoryList"`
			CertificateFlag bool          `json:"certificateFlag"`
			Des             string        `json:"des"`
			DiscordUrl      string        `json:"discordUrl"`
			Image           string        `json:"image"`
			InstagramUrl    string        `json:"instagramUrl"`
			MediumUrl       string        `json:"mediumUrl"`
			Name            string        `json:"name"`
			OfficialWebsite string        `json:"officialWebsite"`
			Slug            string        `json:"slug"`
			Stats           struct {
				FloorPrice  string `json:"floorPrice"`
				LatestPrice string `json:"latestPrice"`
				OwnerCount  string `json:"ownerCount"`
				TotalCount  string `json:"totalCount"`
				TotalVolume string `json:"totalVolume"`
			} `json:"stats"`
			TwitterUrl string `json:"twitterUrl"`
		} `json:"collection"`
		Image             string `json:"image"`
		ImagePreviewUrl   string `json:"imagePreviewUrl"`
		ImageThumbnailUrl string `json:"imageThumbnailUrl"`
		Name              string `json:"name"`
		TokenId           string `json:"tokenId"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type OkxNftCollectionDetail struct {
	Code int `json:"code"`
	Data struct {
		AssetContracts []struct {
			Chain           string `json:"chain"`
			ContractAddress string `json:"contractAddress"`
			Erc2981         bool   `json:"erc2981"`
			OwnerAddress    string `json:"ownerAddress"`
			TokenStandard   string `json:"tokenStandard"`
		} `json:"assetContracts"`
		BackgroundImage string   `json:"backgroundImage"`
		CategoryList    []string `json:"categoryList"`
		CertificateFlag bool     `json:"certificateFlag"`
		Des             string   `json:"des"`
		DiscordUrl      string   `json:"discordUrl"`
		Image           string   `json:"image"`
		InstagramUrl    string   `json:"instagramUrl"`
		MediumUrl       string   `json:"mediumUrl"`
		Name            string   `json:"name"`
		OfficialWebsite string   `json:"officialWebsite"`
		Slug            string   `json:"slug"`
		Stats           struct {
			FloorPrice  string `json:"floorPrice"`
			LatestPrice string `json:"latestPrice"`
			OwnerCount  string `json:"ownerCount"`
			TotalCount  string `json:"totalCount"`
			TotalVolume string `json:"totalVolume"`
		} `json:"stats"`
		TwitterUrl string `json:"twitterUrl"`
	} `json:"data"`
	Msg string `json:"msg"`
}
