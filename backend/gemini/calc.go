package gemini

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	EmbeddingDimension = 768
)

type Embedding []float32

func GetEmbedding(str string, ctx context.Context) (*Embedding, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	em := client.EmbeddingModel("text-embedding-004")

	res, err := em.EmbedContent(ctx, genai.Text(str))
	if err != nil {
		panic(err)
	}

	emb := Embedding(res.Embedding.Values)

	return &emb, nil
}

func (emb *Embedding) String() string {
	var s string
	for _, v := range *emb {
		s += strconv.FormatFloat(float64(v), 'f', -1, 32) + " "

	}
	return s[:len(s)-1]
}

func ScanEmb(val string) (*Embedding, error) {
	emb := Embedding{}
	err := emb.Scan(val)
	return &emb, err
}

func (emb *Embedding) Scan(src string) error {
	var _emb Embedding
	for _, s := range strings.Split(src, " ") {
		v, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		_emb = append(_emb, float32(v))
	}

	*emb = _emb

	return nil
}