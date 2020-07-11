package starwarsadapter

import (
	"fmt"
	"testing"
)

func TestGetPlanets(t *testing.T) {
	planets, err := GetPlanets()
	fmt.Println(planets)
	fmt.Println(err)
}
