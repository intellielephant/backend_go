package service

import (
	"context"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GenerateContent(prompt string) (*genai.GenerateContentResponse, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	return resp, err
}

func UploadFile(filePath, filename string) (*genai.File, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		// 如果未找到 MIME 类型，返回默认类型
		mimeType = "application/pdf"
	}
	options := genai.UploadFileOptions{
		DisplayName: filename,
		MIMEType:    mimeType,
	}
	file, err := client.UploadFileFromPath(ctx, filePath, &options)
	if err != nil {
		log.Fatal(err)
	}

	return file, err
}

func ListFile() ([]*genai.File, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fileList := []*genai.File{}
	iter := client.ListFiles(ctx)
	for {
		ifile, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ifile)
		fileList = append(fileList, ifile)
	}

	return fileList, err
}

func DeleteFile(filename string) error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	err = client.DeleteFile(ctx, filename)

	return err
}
