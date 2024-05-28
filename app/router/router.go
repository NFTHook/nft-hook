package router

import (
	"nfthook/app/controller/nft"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.RouterGroup) {
	userGroup := r.Group("/nft")

	userGroup.GET("/index_list", nft.IndexList)
	userGroup.GET("/nft_detail", nft.NftDetail)

}
