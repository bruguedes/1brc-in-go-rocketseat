package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {

	start := time.Now()
	// os.Open abre um arquivo para leitura, tem como resposta um ponteiro para o arquivo e um erro.
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	// defer fecha o arquivo após a execução, liberando o recurso.
	defer measurements.Close()

	// Cria um mapa para armazenar as medições.
	data := make(map[string]Measurement)

	// bufio.NewScanner cria um scanner para leitura do arquivo.
	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		// scanner.Text() retorna a linha lida.
		rawData := scanner.Text()
		semicolons := strings.Index(rawData, ";")
		location := rawData[:semicolons]
		rawTemp := rawData[semicolons+1:]
		// strconv.ParseFloat converte uma string para um float64.
		temp, err := strconv.ParseFloat(rawTemp, 64)

		if err != nil {
			panic(err)
		}

		measurement, ok := data[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}

		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}
		data[location] = measurement

	}

	locations := make([]string, 0, len(data))
	for locationName := range data {
		locations = append(locations, locationName)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, locationName := range locations {
		measurement := data[locationName]
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
			locationName,
			measurement.Min,
			measurement.Sum/float64(measurement.Count),
			measurement.Max,
		)
	}
	fmt.Printf("}\n")

	fmt.Println("Tempo de execução: ", time.Since(start))
}
