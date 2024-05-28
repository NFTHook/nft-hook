package util

import (
	"encoding/base64"
	"net/http"
	"os"

	ipfs "github.com/ipfs/go-ipfs-api"
)

var ipfsclient *http.Client

const INFURA_IPFS_HOST = "https://ipfs.infura.io:5001"

type Configuration struct {
	ProxyHost string `json:"proxy_host"`
	ProxyPort string `json:"proxy_port"`
}

// 自定义的 Transport 用于添加 HTTP 请求头
type headerTransport struct {
	Transport http.RoundTripper
	Header    http.Header
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return t.Transport.RoundTrip(req)
}

func basicAuth(projectId, projectSecret string) string {
	auth := projectId + ":" + projectSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func init() {
	ipfsclient = &http.Client{
		Transport: &headerTransport{
			Transport: http.DefaultTransport,
			// Header:    http.Header{"Authorization": []string{"Basic " + basicAuth(config.AppCfg.Infura.ProjectID, config.AppCfg.Infura.ProjectSecret)}},
		},
	}
}

func UploadFileToIPFS(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	shell := ipfs.NewShellWithClient(INFURA_IPFS_HOST, ipfsclient)
	cid, err := shell.Add(file)
	if err != nil {
		return "", err
	}

	return cid, nil
}
