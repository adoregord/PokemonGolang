package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"main.go/internal/domain"
	"main.go/internal/handler"
	"main.go/internal/repository"
	"main.go/internal/usecase"
)

func catchProbability(rate float32) string {
	if rate < 0 || rate > 100 {
		return "Invalid rate. Please provide a rate between 0 and 100."
	}

	// rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	// chance := rand.Intn(100) + 1     // Generate a random number between 1 and 100
	chance := rand.Float32() * (100 - 0) // Generate a random number between 1 and 100

	if chance <= rate {
		return "SUCCESS"
	}
	return "FAIL"
}

func main() {
	pokemonRepo := repository.NewPokemonRepo()
	pokemonUc := usecase.NewPokemonUsecase(pokemonRepo)
	pokemonH := handler.NewPokemonHandler(pokemonUc)

	playerRepo := repository.NewPlayerRepo()
	playerUc := usecase.NewPlayerUsecase(playerRepo)
	playerH := handler.NewPlayerHandler(playerUc)

	Scanner := bufio.NewScanner(os.Stdin)

	var player2 domain.Player
	var loggedIn bool
	var isAdmin bool

	// make pokemon list
	pokemons := []domain.Pokemon{
		{Name: "Pikachu", Type: "Lightning", CatchRate: 100, IsRare: false, RegisteredDate: time.Now().Format(time.RFC822)},
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
		if err := pokemonH.PokemonAdd(value); err != nil {
			fmt.Println(err)
			return
		}
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
				lowerUsername := strings.ToLower(usernameStr)
				player2, err := playerH.PlayerLogin(lowerUsername)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					fmt.Println()
					continue
				}
				isAdmin = player2.IsAdmin
				loggedIn = true
			case "R", "r", "register", "Register", "REGISTER":
				fmt.Print("Username: ")
				Scanner.Scan()
				usernameStr := Scanner.Text()
				lowerUsername := strings.ToLower(usernameStr)
				player2.UserName = lowerUsername
				player2, err := playerH.PlayerAdd(player2)
				if err != nil {
					fmt.Println(err)
					time.Sleep(2 * time.Second)
					fmt.Println()
					continue
				}
				isAdmin = player2.IsAdmin
				loggedIn = true
			case "e", "E", "exit", "Exit", "EXIT":
				return
			default:
				fmt.Println(fmt.Errorf("please input L/R only"))
				time.Sleep(2 * time.Second)
				fmt.Println()
			}
		}
		fmt.Printf("\nHi %s What do you wanna do?\n1. Catch pokemon!\n2. View Your Pokemons!\n3. View All players(ADMIN ONLY)\n4. Log out\nSelect: ", player2.UserName)
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
				fmt.Print("Pokemon ID you wanna catch: ")
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
				catch := catchProbability(pokemons[pokemon-1].CatchRate)
				for i := 0; i < 3; i++ {
					fmt.Print("• ")
					time.Sleep(1300 * time.Millisecond)
				}
				if catch == "SUCCESS" {
					fmt.Printf("You've CAUGHT %s!!\n", pokemons[pokemon-1].Name)
					playerH.PlayerAddPokemon(player2.UserName, pokemons[pokemon-1])
					playerH.PlayerViewTheirPokemon(player2.UserName)
					time.Sleep(2 * time.Second)
					break
				}
				fmt.Printf("Wild pokemon %s ran away!", pokemons[pokemon-1].Name)
				time.Sleep(2 * time.Second)
				break
			}
		case 2:
			err := playerH.PlayerViewTheirPokemon(player2.UserName)
			if err != nil {
				fmt.Println(err)
				time.Sleep(2 * time.Second)
				continue
			}
			time.Sleep(2 * time.Second)
		case 3:
			if !isAdmin {
				fmt.Printf("Only Admins can access this feature!\n\n")
				time.Sleep(2 * time.Second)
				continue
			}
			playerH.PlayerView()
			time.Sleep(2 * time.Second)
		case 4:
			loggedIn = false
		default:
			err := fmt.Errorf("you have to input number from 1 to 3")
			fmt.Println(err)
		}
		fmt.Println()
	}
}
