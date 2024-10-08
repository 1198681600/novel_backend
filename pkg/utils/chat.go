package utils

import (
	"context"
	"gitea.peekaboo.tech/peekaboo/crushon-backend/internal/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
	"time"
	"unicode/utf8"
)

func CheckUserContentLength(content string) bool {
	return !(utf8.RuneCountInString(content) > 3000)
}

func MustCalToken(strs ...string) int {
	str := strings.Join(strs, " ")
	if viper.GetBool("service.tokenizer.disable") {
		return len(str) / 4
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 200*time.Millisecond)
	defer cancel()
	type CalculateRequest struct {
		Prompt string `json:"prompt"`
		Model  string `json:"model"`
	}
	type CalculateResponse struct {
		Count int `json:"count"`
	}
	result := new(CalculateResponse)
	resp, err := global.ReqClient.R().
		SetContext(ctx).
		SetBody(CalculateRequest{Prompt: str, Model: "text-davinci-003"}).
		SetSuccessResult(result).
		Post(viper.GetString("service.tokenizer.url"))
	if err != nil || !resp.IsSuccessState() {
		return min(1, len(str)/4)
	}
	return result.Count
}

func TruncateTokenByMaxToken(str string, maxToken int) (string, int, error) {
	if viper.GetBool("service.tokenizer.disable") {
		s := str[:maxToken]
		return s, MustCalToken(s), nil
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 300*time.Millisecond)
	defer cancel()
	type TruncateRequest struct {
		Prompt    string `json:"prompt"`
		Model     string `json:"model"`
		MaxTokens int    `json:"max_tokens"`
	}
	type TruncateResponse struct {
		TruncatedPrompt string `json:"truncated_prompt"`
		TokenCount      int    `json:"token_count"`
	}
	result := new(TruncateResponse)
	resp, err := global.ReqClient.R().
		SetContext(ctx).
		SetBody(
			TruncateRequest{
				Prompt:    str,
				Model:     "text-davinci-003",
				MaxTokens: maxToken,
			},
		).SetSuccessResult(result).
		Post(viper.GetString("service.truncate.url"))
	if err != nil {
		global.Logger.Error("TruncateTokenByMaxToken failed", zap.Error(err))
		return str, 0, err
	}
	if !resp.IsSuccessState() {
		global.Logger.Error("TruncateTokenByMaxToken failed", zap.Any("resp", resp))
		return str, 0, err
	}
	return result.TruncatedPrompt, result.TokenCount, nil
}
