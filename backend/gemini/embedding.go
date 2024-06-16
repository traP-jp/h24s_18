package gemini

func Dot(a Embedding, b Embedding) float64 {
	var sum float64
	for i := range a {
		sum += float64(a[i] * b[i])
	}
	return sum
}

func Norm(a Embedding) float64 {
	var sum float64
	for _, v := range a {
		sum += float64(v * v)
	}
	return sum
}

func CosineSimilarity(a Embedding, b Embedding) float64 {
	return Dot(a, b) / (Norm(a) * Norm(b))
}
