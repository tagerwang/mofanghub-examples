// 最小示例：go get github.com/sashabaranov/go-openai
package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	cfg := openai.DefaultConfig("你的魔方Hub令牌")
	cfg.BaseURL = "https://mofanghub.cn/v1"
	client := openai.NewClientWithConfig(cfg)
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    "glm-5.3-flash",
		Messages: []openai.ChatCompletionMessage{{Role: "user", Content: "你好"}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}
