package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Pokemon struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Image string `json:"img"`
    Type string `json:"type"`
}

func getPokemonImage(urlPokemon string) {
    response, err := http.Get(urlPokemon)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Deu ruim pra ler a resposta da requisição")
        return
    }

    fmt.Println(string(body))
}

func main() {
    data, err := os.ReadFile("data/pokemons.json")
    if err != nil {
        panic(err)
    }

    var pokemons []Pokemon
    json.Unmarshal(data, &pokemons)
    
    for _, p := range pokemons {
        // fmt.Println(p.Name)
        getPokemonImage(p.Image)
    }


}