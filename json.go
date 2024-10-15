package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
    Addr string `json:"addr"`
    Protocol string `json:"protocol"`
    Type string  `json:"type"`
}

type Pokemon struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Image string `json:"img"`
    Type string `json:"type"`
}

func main() {
    data, err := os.ReadFile("data/pokemons.json")
    if err != nil {
        panic(err)
    }

    var pokemons []Pokemon
    json.Unmarshal(data, &pokemons)
    
    for _, p := range pokemons {
        fmt.Println(p.Name)
    }


}