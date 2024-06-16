package gemini

import (
	"context"
	"fmt"
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

func GetEmbedding(str string, ctx context.Context) (Embedding, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	em := client.EmbeddingModel("text-embedding-004")

	ctx = context.Background()

	res, err := em.EmbedContent(ctx, genai.Text(str))
	if err != nil {
		panic(err)
	}

	val := make([]float32, EmbeddingDimension)

	in := copy(val, res.Embedding.Values)

	fmt.Println(res.Embedding.Values)

	fmt.Println(in)

	emb := Embedding(val)

	fmt.Println(emb[0], emb[1], str)

	return emb, nil
}

func (emb *Embedding) String() string {
	var s string
	for _, v := range *emb {
		s += strconv.FormatFloat(float64(v), 'f', -1, 32) + " "

	}
	return s[:len(s)-1]
}

func (emb *Embedding) len() int {
	return len(*emb)

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
