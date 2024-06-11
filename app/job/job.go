package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nfthook/app/model"
	"nfthook/database"
	"nfthook/util"
	"strconv"

	"github.com/robfig/cron"
)

func init() {
	c := cron.New()

	c.AddFunc("@every 900s", UpdateNFTOwnerInfo)
	c.AddFunc("@every 900s", UpdateZoraNFTOwnerInfo)

	c.Start()
}

func UpdateNFTOwnerInfo() {

	//查询出所有非zora的数据
	db := database.DB

	var nftTokens []model.NftToken

	db.Where(" chain_name in ('Op','Base')").Find(&nftTokens)

	for _, v := range nftTokens {

		chain := "optimism"

		if v.ChainName == "Base" {
			chain = "base"
		}

		//查询出对应的 slug
		getRequestPath := "/api/v5/mktplace/nft/asset/detail"
		getParams := map[string]string{
			"chain":           chain,
			"contractAddress": v.ContractAddr,
			"tokenId":         "1",
		}

		result, _ := util.OkxSendGetRequest(getRequestPath, getParams)

		var okxNftAssetDetail util.OkxNftAssetDetail

		// 解析 JSON 数据
		err := json.Unmarshal([]byte(result), &okxNftAssetDetail)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// fmt.Println(v.ContractAddr, "   ", okxNftAssetDetail.Data.Collection.Slug)

		//查询出Holder 数据

		getRequestPath = "/api/v5/mktplace/nft/collection/detail"

		getParams = map[string]string{
			"slug": okxNftAssetDetail.Data.Collection.Slug,
		}

		result, _ = util.OkxSendGetRequest(getRequestPath, getParams)

		var okxNftCollectionDetail util.OkxNftCollectionDetail

		// 解析 JSON 数据
		err = json.Unmarshal([]byte(result), &okxNftCollectionDetail)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		fmt.Println(v.ContractAddr, "   ", okxNftAssetDetail.Data.Collection.Slug, "    原值:", v.HolderNum, "    ", okxNftCollectionDetail.Data.Stats.OwnerCount)

		// 将字符串转换为 int64
		intValue, err := strconv.ParseInt(okxNftCollectionDetail.Data.Stats.OwnerCount, 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int64:", err)
			continue
		}

		if v.HolderNum != intValue {
			db.Model(&model.NftToken{}).Where("id = ?", v.ID).Update("holder_num", okxNftCollectionDetail.Data.Stats.OwnerCount)
		}

	}

}

type ZoraResponse struct {
	TokenHoldersCount string `json:"token_holders_count"`
	TransfersCount    string `json:"transfers_count"`
}

func UpdateZoraNFTOwnerInfo() {

	db := database.DB

	var nftTokens []model.NftToken

	db.Where(" chain_name in ('Zora')").Find(&nftTokens)

	for _, v := range nftTokens {

		baseURL := "https://explorer.zora.energy/api/v2/tokens/"
		contractAddress := v.ContractAddr
		endpoint := "/counters"
		url := baseURL + contractAddress + endpoint

		// 创建一个新的HTTP请求
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		// 设置请求头
		req.Header.Set("Accept", "application/json")

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		defer resp.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		// 解析JSON响应
		var zoraResponse ZoraResponse
		err = json.Unmarshal(body, &zoraResponse)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// 打印解析后的数据
		fmt.Println("Token Holders Count:", v.ContractAddr, v.Name, v.HolderNum, zoraResponse.TokenHoldersCount)

		// 将字符串转换为 int64
		intValue, err := strconv.ParseInt(zoraResponse.TokenHoldersCount, 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int64:", err)
			continue
		}

		if v.HolderNum != intValue {
			db.Model(&model.NftToken{}).Where("id = ?", v.ID).Update("holder_num", zoraResponse.TokenHoldersCount)
		}
	}
}
