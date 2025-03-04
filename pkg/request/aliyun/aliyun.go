package aliyun

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/ddliu/go-httpclient"
	"github.com/m1nule/aliyun-m3u8-downloader/pkg/tool"
)

func init() {
	httpclient.Defaults(httpclient.Map{
		"Accept":                 "application/json, text/plain, */*",
		"Accept-Encoding":        "gzip, deflate, br",
		"Accept-Language":        "zh-CN,zh;q=0.9,en;q=0.8",
		httpclient.OPT_USERAGENT: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36",
		// httpclient.OPT_PROXY:      "http://127.0.0.1:8888",
		httpclient.OPT_UNSAFE_TLS: true,
	})
}

func GetVodPlayerInfo(rand, playAuth string, opts ...OptionFunc) (*simplejson.Json, error) {
	rand, _ = tool.EncryptRand([]byte(rand))
	return getVodPlayerInfo(rand, playAuth, opts...)
}

func getVodPlayerInfo(rand, playAuth string, opts ...OptionFunc) (*simplejson.Json, error) {
	playInfoRequestUrl, err := GetPlayInfoRequestUrl(rand, playAuth, opts...)
	if err != nil {
		return nil, err
	}
	resp, err := httpclient.Get(playInfoRequestUrl)
	if err != nil {
		return nil, fmt.Errorf("getVodPlayerInfo: http get url: %s, err: %w", playInfoRequestUrl, err)
	}
	data, err := resp.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("getVodPlayerInfo: read err: %w", err)
	}
	sj, err := simplejson.NewJson(data)
	if err != nil {
		return nil, fmt.Errorf("getVodPlayerInfo: json decode: %s, err: %w", data, err)
	}
	return sj, nil
}
