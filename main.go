package main

import (
	"fmt"
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
			fmt.Println("Error: ", err)
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

	/*

		Requisitar ao servidor
		Verificar a resposta
		Parsear a reposta
		Imprimir a reposta

	*/

}
