// Making a Self made Shell with GO 
//
//

package main 
import "fmt"

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read the keyboard input 
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

