package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/google/generative-ai-go/genai"

	"google.golang.org/genai"
)

const ModelName string = "gemini-2.0-flash-exp"

func GenerateContent2(prompt, tool string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGoogleAI,
	})

	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{}

	if tool == "search" {
		config.Tools = []*genai.Tool{
			{GoogleSearchRetrieval: &genai.GoogleSearchRetrieval{}}}
	} else if tool == "code" {
		config.Tools = []*genai.Tool{
			{CodeExecution: &genai.ToolCodeExecution{}}}
	} else if tool == "default" {
		config = nil
	} else {
		config = nil
	}
	result, err := client.Models.GenerateContent(ctx,
		ModelName,
		genai.Text(prompt),
		config,
	)

	if err != nil {
		log.Fatal(err)
	}

	return result, err
}

func DescribeFile(fileUrl, prompt string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGoogleAI,
	})

	if err != nil {
		fmt.Println("Init client fail:", err)
		return nil, err
	}

	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("Error fetching image:", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read image fail:", err)
		return nil, err
	}

	ext := filepath.Ext(fileUrl)

	mimeType := mime.TypeByExtension(ext)

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: mimeType}},
	}
	contents := []*genai.Content{{Parts: parts}}

	result, err := client.Models.GenerateContent(ctx, ModelName, contents, nil)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}
