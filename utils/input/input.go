package input

import (
	"bufio"
	"fmt"
	"os"

	strutil "github.com/dterbah/gopm/utils/string"
)

// ReadUserInput lit l'entrée utilisateur avec un message.
func ReadUserInput(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (%s) ", prompt, defaultValue)
	input, _ := reader.ReadString('\n')
	return strutil.GetStringIfEmpty(input, defaultValue)
}
