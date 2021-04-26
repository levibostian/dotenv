package helpers

import (
	"log"
	"os/exec"
	"strings"
)

func Command(binPath string, args string) (output string) {
	// binPath, _ := filepath.Abs("../")

	// cmd := exec.Command(fmt.Sprintf("%s/dotenv", binPath), command)
	cmd := exec.Command(binPath, strings.Split(args, " ")...)
	outBytes, err := cmd.Output()
	output = string(outBytes)

	if err != nil {
		log.Printf(err.Error())
		// os.Exit(1)
	}

	return output
}
