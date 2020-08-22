package hangman

import "fmt"

// You can generate ASCII art on http://patorjk.com/software/taag/

func DrawWelcome() {
	fmt.Println(`
	_______  _______  _______           ______                      
	(  ____ )(  ____ \(       )|\     /|(  __  \                     
	| (    )|| (    \/| () () || )   ( || (  \  )                    
	| (____)|| (__    | || || || (___) || |   ) |                    
	|     __)|  __)   | |(_)| ||  ___  || |   | |                    
	| (\ (   | (      | |   | || (   ) || |   ) |                    
	| ) \ \__| (____/\| )   ( || )   ( || (__/  )                    
	|/   \__/(_______/|/     \||/     \|(______/                     
																	 
			  _______  _        _______    _______  _______  _       
	|\     /|(  ___  )( (    /|(  ____ \  (       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/  | () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |        | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____   | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )  | |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) |  | )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)  |/     \||/     \||/    )_)															 
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Print("You lost :( ! The word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("YOU WON MATE! The word was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
