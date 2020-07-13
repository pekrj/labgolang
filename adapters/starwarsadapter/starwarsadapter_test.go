package starwarsadapter

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetPlanets(t *testing.T) {
	planets, err := GetPlanets()
	fmt.Println(planets)
	fmt.Println(err)
}

func TestGetPlanetsWithGocks(t *testing.T) {
	defer gock.Off()

	dat, err := ioutil.ReadFile("../filesmocks/planetmock.json")
	assert.NoError(t, err)
	fmt.Print(string(dat))

	gock.New(url).
		Get("/planets").
		Reply(200).
		BodyString(string(dat))

	planets, err := GetPlanets()
	fmt.Println(err)
	fmt.Println(planets)

	if assert.NoError(t, err) {
		assert.Equal(t, 9, len(planets))
	}

}
