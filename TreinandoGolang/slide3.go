package main

import (
	"errors"
	"fmt"
)

type Videoo struct {
	Duracao        int
	NivelEntretido int
}

func main() {
	videos := []Videoo{
		{Duracao: 120, NivelEntretido: 8},
		{Duracao: 90, NivelEntretido: 9},
		{Duracao: 150, NivelEntretido: 7},
		{Duracao: 80, NivelEntretido: 10},
	}

	tempoDisponivel := 200

	videoEscolhido, err := escolherVideo(videos, tempoDisponivel)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Vídeo escolhido:", videoEscolhido)
	}
}

func escolherVideoo(videos []Videoo, tempoDisponivel int) (Videoo, error) {
	var videoEscolhido Videoo
	maiorNivelEntretido := 0

	for _, video := range videos {
		if video.Duracao <= tempoDisponivel && video.NivelEntretido > maiorNivelEntretido {
			videoEscolhido = video
			maiorNivelEntretido = video.NivelEntretido
		}
	}

	if maiorNivelEntretido == 0 {
		return Video{}, errors.New("Nenhum vídeo adequado encontrado")
	}

	return videoEscolhido, nil
}
