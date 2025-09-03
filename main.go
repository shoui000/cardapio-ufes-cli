package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var diasDaSemana = []string {
	"Domingo",
	"Segunda-Feira",
	"Terça-Feira",
	"Quarta-Feira",
	"Quinta-Feira",
	"Sexta-Feira",
	"Sábado",
}

func main() {
	var data time.Time 

	if len(os.Args) > 1 {
		var err error

		data, err = time.Parse(time.DateOnly, os.Args[1])

		if err != nil {
			fmt.Println("Erro ao receber o input:", err)
			os.Exit(1)
		}

	} else {
		data = time.Now()
	}

	if data.Weekday() == 0 || data.Weekday() == 6 {
		fmt.Printf("Hoje é %s e o RU da UFES de Goiabeiras não abre.\n", diasDaSemana[data.Weekday()],)

		if data.Weekday() == 0 {
			data = data.AddDate(0, 0, -2)
		} else if data.Weekday() == 6 {
			data = data.AddDate(0, 0, -1)
		}

		fmt.Printf("Tomando como base o dia %s, %s.\n", data.Format(time.DateOnly), diasDaSemana[data.Weekday()])
	}

	criacaoRU := time.Date(1968, time.March, 3, 0, 0, 0, 0, time.UTC)

	if data.Before(criacaoRU) {
		fmt.Println("Data utilizada anterior ao dia da criação do RU.")
		os.Exit(1)
	}

	url := fmt.Sprintf("https://ru.ufes.br/cardapio/rss/%s", data.Format(time.DateOnly))
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro ao fazer a requisição http:", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			fmt.Println("Não foi encontrado cardápio para este dia")
			os.Exit(1)
		}

		fmt.Println("Erro ao fazer a requisição http:", resp.Status)
		os.Exit(1)
	}
	/*

		Requisitar ao servidor
		Verificar a resposta
		Parsear a reposta
		Imprimir a reposta

	*/

}
