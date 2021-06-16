package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Pokemon struct {
	tableName      struct{} `pg:"pokemon"`
	Number         int      `json:"number"`
	Name           string   `json:"name"`
	Type1          string   `json:"type1"`
	Type2          string   `json:"type2"`
	Color          string   `json:"color"`
	Ability1       string   `json:"ability1"`
	Ability2       string   `json:"ability2"`
	HiddenAbility  string   `json:"hiddenAbility" pg:"hiddenability"`
	Generation     int      `json:"generation"`
	Height         string   `json:"height"`
	Weigth         string   `json:"weigth"`
	Hp             int      `json:"hp"`
	Attack         int      `json:"attack"`
	Defense        int      `json:"defense"`
	SpecialAttack  int      `json:"specialAttack" pg:"specialattack"`
	SpecialDefense int      `json:"specialDefense" pg:"specialdefense"`
	Speed          int      `json:"speed"`
}

func getPokemonInfo(c echo.Context) error {

	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "",
		Database: "pokedex",
	})

	id, _ := strconv.Atoi(c.Param("id"))

	info := &Pokemon{}
	var err = db.Model(info).Where("number = ?", id).Select()
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, "Can't find an entry with the number #"+c.Param("id"))
	}

	return c.JSONPretty(http.StatusOK, info, " ")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/pokedex/:id", getPokemonInfo)
	e.Logger.Fatal(e.Start(":1323"))
}
