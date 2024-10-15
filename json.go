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

    // var configs []Config
    // err = json.Unmarshal(data, &configs)
    // if err != nil {
    //     panic(err)
    // }

    var pokemons []Pokemon
    json.Unmarshal(data, &pokemons)
    
    for _, p := range pokemons {
        fmt.Println(p.Name)
    }

    // configs = append(configs, Config{"canaistech.com.br", "https", "website"})

    // ndata, err := json.Marshal(configs)
    // if err != nil {
	// 	panic(err)
	// }

    // err = os.WriteFile("data/config.json", ndata, 0755)
	// if err != nil {
	// 	panic(err)
	// }
}