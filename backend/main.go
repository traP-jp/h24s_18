package main

import (
	"encoding/gob"
	"net/http"

	"golang.org/x/oauth2"

	//"github.com/google/generative-ai-go/genai"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/h24s_18/model"

	"github.com/traP-jp/h24s_18/handler"
	//"google.golang.org/api/option"
)

func main() {

	//ctx := context.Background()
	//emb, err := gemini.GetEmbedding("せいかつ", ctx)
	//if err != nil {
	//	panic(err)
	//}
	//emb2, err := gemini.GetEmbedding("あそび", ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(emb)
	//fmt.Println(emb2)
	//
	//return

	gob.Register(&oauth2.Token{})

	user := os.Getenv("MARIADB_USERNAME")
	pass := os.Getenv("MARIADB_PASSWORD")
	host := os.Getenv("MARIADB_HOSTNAME")
	dbname := os.Getenv("MARIADB_DATABASE")

	// ctx := context.Background()

	//client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer client.Close()
	//
	//em := client.EmbeddingModel("text-embedding-004")
	//
	//bat := em.NewBatch().AddContent(genai.Text("日本人")).AddContent(genai.Text("勤勉"))
	//
	//res, err := em.BatchEmbedContents(ctx, bat)
	//if err != nil {
	//	panic(err)
	//}
	//
	//val := make([]float32, 0, 768)
	//
	//for i, v := range res.Embeddings[0].Values {
	//	val = append(val, v+res.Embeddings[1].Values[i])
	//}

	err := model.Init(user, pass, host, dbname)
	if err != nil {
		panic(err)
	}

	// Echoの新しいインスタンスを作成
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"http://localhost:5173", "https://encounter.trap.show"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderAccessControlAllowOrigin, echo.HeaderOrigin, echo.HeaderXHTTPMethodOverride, echo.HeaderContentType},
	}))

	e.Use(middleware.Logger())

	// 「/hello」というエンドポイントを設定する
	e.GET("/hello", handler.GetHello)

	e.GET("/api/users/:userId", handler.GetUser)
	e.GET("/api/oauth2/authorize", handler.AuthorizeHandler)
	e.GET("/api/oauth2/callback", handler.CallbackHandler)
	e.GET("/api/me", handler.GetMeHandler)
	e.PATCH("/api/me", handler.PatchMe)
	e.GET("/api/tags", handler.GetTags)
	e.GET("/api/users", handler.FindUserByTag)
	e.POST("/api/me/tags", handler.PostUserTag)
	e.POST("/api/me/tags/bulk", handler.BulkInsertUserTags)
	e.PATCH("api/me/tags", handler.UpdateUserTag)
	e.DELETE("/api/me/tags/:tagName", handler.DeleteUserTag)

	// Webサーバーをポート番号8080で起動し、エラーが発生した場合はログにエラーメッセージを出力する
	e.Logger.Fatal(e.Start(":8080"))
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
