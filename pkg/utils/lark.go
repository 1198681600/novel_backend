package utils

import (
	"encoding/json"
	"fmt"
	"gitea.peekaboo.tech/peekaboo/crushon-backend/internal/global"
	"gitea.peekaboo.tech/peekaboo/crushon-backend/pkg"
	"go.uber.org/zap"
)

func LarkNotifyBody(larkUrl string, title string, content string, metadata []map[string]string) {
	md := make([]map[string]string, 0, len(metadata)+1)
	md = append(md, map[string]string{
		"tag":  "text",
		"text": content + "\n",
	})

	for _, v := range metadata {
		v["tag"] = "text"

	}
	md = append(md, metadata...)
	data := map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": []any{
						md,
					},
				},
			},
		},
	}
	//str, _ := json.Marshal(data)
	//fmt.Println(string(str))

	var result json.RawMessage
	resp, err := global.ReqClient.R().SetBody(data).
		SetSuccessResult(&result).
		Post(larkUrl)
	if err != nil {
		return
	}
	if !resp.IsSuccessState() {
		return
	}
	//bytes, _ := json.Marshal(result)
	//fmt.Println(string(bytes))
}

type NotifyMetadataEntry struct {
	Key   string
	Value string
}

func LarkNotifyBodyV2(larkURL string, title string, content string, metadata []NotifyMetadataEntry) {
	if !pkg.IsValidURL(larkURL) {
		global.Logger.Error("LarkNotifyBody lark url invalid", zap.String("url", larkURL))
		return
	}
	md := make([]map[string]string, 0, len(metadata)+1)
	md = append(md, map[string]string{
		"tag":  "text",
		"text": content + "\n",
	})
	for _, entry := range metadata {
		md = append(md, map[string]string{
			"tag":  "text",
			"text": fmt.Sprintf("%s: %s\n", entry.Key, entry.Value),
		})
	}
	data := map[string]any{
		"msg_type": "post",
		"content": map[string]any{
			"post": map[string]any{
				"zh_cn": map[string]any{
					"title": title,
					"content": []any{
						md,
					},
				},
			},
		},
	}
	//str, _ := json.Marshal(data)
	//fmt.Println(string(str))

	resp, err := global.ReqClient.R().SetBody(data).
		Post(larkURL)
	//Post("https://open.feishu.cn/open-apis/bot/v2/hook/5f7c679d-24f6-4363-8aad-97894e349029")
	if err != nil {
		global.Logger.Error("LarkNotifyBody error", zap.Error(err))
		return
	}
	if !resp.IsSuccessState() {
		global.Logger.Error("LarkNotifyBody error", zap.Error(err))
		return
	}
}
