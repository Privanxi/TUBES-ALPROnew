package main

import "fmt"

func main() {
	var restart bool = false
	if nPlayer < 26 {
		printErrorScreen("Sesuatu Tidak Beres")
	}

	entryBanner()
	filePlayer := loadInitialPlayers()

	for !restart {
		user := menuAwal(&filePlayer)
		if user.name == "" {
			fmt.Println("Program selesai, Terima kasih sudah bermain BATTLESHIP!!")
			restart = true
		} else {
			data := initGame(user.difficulty)
		
			// Reset statistik pemain sebelum bermain
			user.hits = 0
			user.numOfTurns = 0
			user.numOfFail = 0

			playGame(user, &data)

			fmt.Println("\nMau main lagi? (y/n): ")
			var again string
			fmt.Scan(&again)
			if again != "y" && again != "Y" {
				fmt.Println("Terima kasih sudah bermain!")
				restart = true
			}
		}
		
	}
}