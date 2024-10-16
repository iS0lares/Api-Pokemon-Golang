package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Pokemon struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Image string `json:"img"`
    Type string `json:"type"`
}

func getPokemonImage(urlPokemon string) {
    
    split := strings.Split(urlPokemon, "/")
    filename := split[len(split)-1]
    response, err := http.Get(urlPokemon)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    

    f, err := os.OpenFile("images/"+filename, os.O_CREATE|os.O_RDWR, 0640)

    if err != nil {
        panic(err)
    }

    body, err := io.ReadAll(response.Body)


    if err != nil {
        fmt.Println("Deu ruim pra ler a resposta da requisição")
        return
    }

    bufio.NewWriter(f).Write(body)


    // fmt.Println(reflect.TypeOf(body))
}

func main() {
    data, err := os.ReadFile("data/pokemons.json")
    if err != nil {
        panic(err)
    }

    var pokemons []Pokemon
    json.Unmarshal(data, &pokemons)

    var wg sync.WaitGroup
    
    for _, p := range pokemons {
        wg.Add(1)

        go func () {
            defer wg.Done()
            getPokemonImage(p.Image)
        }()
            // getPokemonImage(p.Image)

        
    }

    wg.Wait()


}