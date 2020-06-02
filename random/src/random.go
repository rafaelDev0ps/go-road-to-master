package random

import "math/rand"

type Generator struct {
	Seed int64
}

type NumberGenerator interface {
	GerarNumero(seed int64) int64
}

func (gen Generator) GerarNumero(seed int64) int64 {
	novoRand := rand.New(rand.NewSource(seed))
	return novoRand.Int63()
}
