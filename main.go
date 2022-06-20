package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	Id   string `json:"id"`
	Name string `json:"nombre"`
	Type string `json:"tipo"`
}

type PokemonReponse struct {
	Id   string `json:"id"`
	Name string `json:"nombre"`
}

func ParsePokemonResponse(all []Pokemon) []PokemonReponse {
	var pResponse = []PokemonReponse{}
	for _, p := range all {
		pResponse = append(pResponse, PokemonReponse{Id: p.Id, Name: p.Name})
	}
	return pResponse
}

func FindPokemonById(all []Pokemon, id string) *Pokemon {
	for _, p := range all {
		if id == p.Id {
			return &p
		}
	}
	return nil
}

func FindPokemonLocation(all []Pokemon, id string) int {
	var location = 0
	for _, p := range all {
		if id != p.Id {
			location++
		} else {
			return location
		}
	}
	return -1
}

func main() {
	pokemons := []Pokemon{
		{Id: "1", Name: "pikachu", Type: "electricidad"},
		{Id: "2", Name: "charizard", Type: "fuego"},
		{Id: "3", Name: "bulbasaur", Type: "planta"},
		{Id: "4", Name: "mewto", Type: "gato"},
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pokemon", func(c *gin.Context) {
		pr := ParsePokemonResponse(pokemons)
		c.JSON(http.StatusOK, pr)
	})

	r.GET("/pokemon/:id", func(c *gin.Context) {
		pokemonId := c.Param("id")
		pokemon := FindPokemonById(pokemons, pokemonId)
		if pokemon == nil {
			c.JSON(http.StatusNotFound, gin.H{"mensaje": "El pokemon solicitado no existe"})
			return
		}

		c.JSON(http.StatusOK, pokemon)
	})

	r.POST("/pokemon", func(c *gin.Context) {
		var newPokemon Pokemon
		if err := c.BindJSON(&newPokemon); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"mensaje": "No se encuentran los parámetros necesarios"})
			return
		}
		pokemons = append(pokemons, newPokemon)
		c.IndentedJSON(http.StatusCreated, newPokemon)
	})

	r.PUT("/pokemon/:id", func(c *gin.Context) {
		pokemonId := c.Param("id")
		newPokemonId := FindPokemonLocation(pokemons, pokemonId)

		var updatePokemon Pokemon
		if err := c.BindJSON(&updatePokemon); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"mensaje": "No se encuentran los parámetros necesarios"})
			return
		}

		pokemons[newPokemonId].Id = updatePokemon.Id
		pokemons[newPokemonId].Name = updatePokemon.Name
		pokemons[newPokemonId].Type = updatePokemon.Type

		c.IndentedJSON(http.StatusOK, updatePokemon)
	})

	r.DELETE("/pokemon/:id", func(c *gin.Context) {
		pokemonId := c.Param("id")
		id := FindPokemonLocation(pokemons, pokemonId)

		if id-1 < 0 || id-1 >= len(pokemons) {
			c.JSON(http.StatusNotFound, gin.H{"mensaje": "ID fuera de rango"})
			return
		} else {
			newLength := 0
			for index := range pokemons {
				if id != index {
					pokemons[newLength] = pokemons[index]
					newLength++
				}
			}
			pokemons = pokemons[:newLength]
		}
		c.Status(http.StatusNoContent)
	})

	r.Run(":8765")
}
