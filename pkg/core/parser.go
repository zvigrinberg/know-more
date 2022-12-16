package core
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func parseCommandLine(line string) {
   const delimiter = "%*@&"
	components:= strings.Split(line,delimiter)
	date:= components[0]
   fmt.Println("fds")
}

func parseEnvironmentVariables(line string) {
  )
}
