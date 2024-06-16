package gemini

func Dot(a Embedding, b Embedding) float32 {
	var sum float32
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

func Norm(a Embedding) float32 {
	var sum float32
	for _, v := range a {
		sum += v * v
	}
	return sum
}

func CosineSimilarity(a Embedding, b Embedding) float32 {
	return Dot(a, b) / (Norm(a) * Norm(b))
}
