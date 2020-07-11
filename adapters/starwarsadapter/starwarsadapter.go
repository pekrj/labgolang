package starwarsadapter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pekrj/labgolang/domain"
)

var JsonEncodingError = errors.New("json enconding error")
var IoReadError = errors.New("io read error")
var ClientRequestError = errors.New("client request error")

const url = "https://swapi.dev/api/"

type PlanetsTO struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

type resultTO struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []PlanetsTO `json:"results"`
}

func GetPlanets() ([]domain.Planet, error) {
	resp, err := http.Get(url + "planets")
	if err != nil {
		return nil, ClientRequestError
	}
	defer resp.Body.Close()
	var planetlist resultTO

	err = json.NewDecoder(resp.Body).Decode(&planetlist)
	if err != nil {
		return nil, JsonEncodingError
	}
	var aaa []domain.Planet
	aaa = make([]domain.Planet, len(planetlist.Results), len(planetlist.Results))

	for i, planetTo := range planetlist.Results {
		aaa[i] = domain.Planet{Name: planetTo.Name, Climate: planetTo.Climate, Terrain: planetTo.Terrain}
	}

	return aaa, nil
}
