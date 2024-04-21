package main

func main() {
	//export https_proxy=http://127.0.0.1:15236
	//export http_proxy=http://127.0.0.1:15236

	// 读取指定目录下的文件夹的图片
	// imgPathList, _ := util.ListImageFiles("/Users/wanggaowei/Documents/Fire Girl/fire4/")

	// for i, v := range imgPathList {

	// 	// if i < 5 {
	// 	// 	continue
	// 	// }

	// 	fmt.Println("开始上传：", i, v)
	// 	//上传到ipfsls
	// 	cid, err := util.UploadFileToIPFS(v)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		panic("上传失败")
	// 	}

	// 	//获取到cid
	// 	// nftAttr := model.NFTAttribute{TraitType: "LEVEL", Value: util.GetRandomRarity()}
	// 	newNft := model.NFT{Name: "God's fire.", Image: "https://ipfs.io/ipfs/" + cid}

	// 	//生成一个json 写入到./result/json
	// 	util.WriteToFile(newNft, "./result/json/"+strconv.Itoa(i+1)+".json")
	// }
}
