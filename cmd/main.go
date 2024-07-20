package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"main.go/internal/domain"
	"main.go/internal/handler"
	"main.go/internal/repository"
	"main.go/internal/usecase"
)

func main() {
	pokemonRepo := repository.NewPokemonRepo()
	pokemonUc := usecase.NewPokemonUsecase(pokemonRepo)
	pokemonH := handler.NewPokemonHandler(pokemonUc)

	playerRepo := repository.NewPlayerRepo()
	playerUc := usecase.NewPlayerUsecase(playerRepo)
	playerH := handler.NewPlayerHandler(playerUc)

	Scanner := bufio.NewScanner(os.Stdin)

	var player2 domain.Player
	loggedIn := false

	// make pokemon list
	pokemons := []domain.Pokemon{
		{Name: "Pikachu", Type: "Lightning", CatchRate: 70, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Bulbasaur", Type: "Plant", CatchRate: 50, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Charmander", Type: "Fire", CatchRate: 50, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Rattata", Type: "Normal", CatchRate: 80, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Ditto", Type: "Normal", CatchRate: 30, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Mew Two", Type: "Psychic", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Dialga", Type: "Steel/Dragon", CatchRate: 0.001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
		{Name: "Arceus", Type: "Normal", CatchRate: 0.000001, IsRare: true, RegisteredDate: time.Now().Format(time.RFC822)},
	}
	// add to the databse through handler
	for _, value := range pokemons {
		pokemonH.PokemonAdd(value)
	}
	for {
		for !loggedIn {
			fmt.Printf("Please login/register first\nPress E for exit\nLogin/Register [L/R]: ")
			Scanner.Scan()
			optionStr := Scanner.Text()
			switch optionStr {
			case "L", "l", "login", "Login", "LOGIN":
				fmt.Print("Username: ")
				Scanner.Scan()
				usernameStr := Scanner.Text()
				player2.UserName = usernameStr
				_, err := playerH.PlayerLogin(usernameStr)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}
				loggedIn = true
			case "R", "r", "register", "Register", "REGISTER":
				fmt.Print("Username: ")
				Scanner.Scan()
				usernameStr := Scanner.Text()
				player2.UserName = usernameStr
				err := playerH.PlayerAdd(player2)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					continue
				}
				loggedIn = true
			case "e", "E", "exit", "Exit", "EXIT":
				return
			default:
				fmt.Println(fmt.Errorf("please input L/R only"))
				time.Sleep(2 * time.Second)
				fmt.Println()
			}
		}
		fmt.Printf("\nHi %s What do you wanna do?\n1. Catch pokemon!\n2. View Your Pokemons!\n3. Log out\nSelect: ", player2.UserName)
		Scanner.Scan()
		todoStr := Scanner.Text()
		todo, err := strconv.Atoi(todoStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch todo {
		case 1:
			fmt.Printf("\nWhich Pokemon you wanna catch %s?\n", player2.UserName)
			pokemonH.PokemonView()
			for {
				fmt.Print("Pokemon ID you wanna catch:")

				Scanner.Scan()
				pokemonStr := Scanner.Text()
				pokemon, err := strconv.Atoi(pokemonStr)
				if err != nil {
					fmt.Print(err.Error())
				}
				if pokemon < 1 || pokemon > len(pokemons) {
					err := fmt.Errorf("pokemon ID you must input is from 1 to %d", len(pokemons))
					fmt.Println(err)
					time.Sleep(3 * time.Second)
					continue
				}
				playerH.PlayerAddPokemon(player2.UserName, pokemons[pokemon-1])
				playerH.PlayerViewTheirPokemon(player2.UserName)
				time.Sleep(2 * time.Second)
				break
			}
		case 2:
			playerH.PlayerViewTheirPokemon(player2.UserName)
			time.Sleep(2 * time.Second)
		case 3:
			loggedIn = false
		default:
			err := fmt.Errorf("you have to input number from 1 to 3")
			fmt.Println(err)
		}
		fmt.Println()
	}
}
