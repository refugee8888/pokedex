package mypokedex

type StatsPokemon struct {
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Stats      []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Pokedex struct {
	MyPokemons map[string]StatsPokemon
}

func New() *Pokedex {
	return &Pokedex{
		MyPokemons: make(map[string]StatsPokemon),
	}
}

func (p *Pokedex) Add(pk StatsPokemon) {
	p.MyPokemons[pk.Name] = pk
}
