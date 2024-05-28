package nft

import (
	"encoding/json"
	"log"
	"net/http"

	"nfthook/app/model"
	"nfthook/database"
	ml "nfthook/middleware"

	"github.com/gin-gonic/gin"
)

func IndexList(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	chainId := c.Query("chain_id")

	db := database.DB
	var tokens []model.NftToken

	if len(chainId) > 0 {
		db = db.Where(" chain_id = ? ", chainId)
	}

	db.Find(&tokens)

	c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"list": tokens}))

}

func NftDetail(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	addr := c.Query("addr")

	db := database.DB

	var token model.NftToken
	db.Where("contract_addr = ? ", addr).First(&token)

	// abi := `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"msgSender","type":"address"},{"indexed":true,"internalType":"uint256","name":"mintQuantity","type":"uint256"}],"name":"NewMint","type":"event"},{"inputs":[{"internalType":"uint256","name":"quantity","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"payable","type":"function"}]`

	var abis []interface{}

	err := json.Unmarshal([]byte(token.Abi), &abis)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// btn := `[{"mint_cnt":1,"value":"7777000000000000"},{"mint_cnt":3,"value":"7777000000000000"}]`

	var bnts []interface{}

	// Unmarshal JSON to the interface{} slice
	err = json.Unmarshal([]byte(token.Btns), &bnts)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"info": token, "abi": abis, "btns": bnts}))

}
