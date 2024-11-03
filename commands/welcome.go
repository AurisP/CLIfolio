package commands

const welcomeText = `    
  _  _                ___ _           _           _    
 | || |___ _  _      |_ _( )_ __     /_\ _  _ _ _(_)___
 | __ / -_) || |  _   | ||/| '  \   / _ \ || | '_| (_-<
 |_||_\___|\_, | ( ) |___| |_|_|_| /_/ \_\_,_|_| |_/__/
           |__/  |/ 

Hello dear User! Welcome to my CLI portfolio! Type 'help' to see the available commands.`

func Welcome() string {
	return welcomeText
}
