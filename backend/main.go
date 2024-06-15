package main

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_18/model"
	"os"

	"github.com/traP-jp/h24s_18/handler"
	"google.golang.org/api/option"
)

func main() {
	user := os.Getenv("MARIADB_USERNAME")
	pass := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOSTNAME")
	dbname := os.Getenv("MARIADB_DATABASE")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAow8apSTGtxzmskl7PtCXxNTJ-3I472Wg"))

	if err != nil {
		panic(err)
	}
	defer client.Close()

	em := client.EmbeddingModel("text-embedding-004")

	bat := em.NewBatch().AddContent(genai.Text("日本人")).AddContent(genai.Text("勤勉"))

	res, err := em.BatchEmbedContents(ctx, bat)
	if err != nil {
		panic(err)
	}

	val := make([]float32, 0, 768)

	for i, v := range res.Embeddings[0].Values {
		val = append(val, v+res.Embeddings[1].Values[i])
	}

	err = model.Init(user, pass, host, dbname)
	if err != nil {
		panic(err)
	}

	// Echoの新しいインスタンスを作成
	e := echo.New()

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", handler.GetHello)

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}