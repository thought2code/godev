package osutil

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/thought2code/godev/internal/strconst"
)

func RunCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	output, err := command.CombinedOutput()

	fmt.Printf("%s %s %s\n", strconst.EmojiRunning, cmd, strings.Join(args, strconst.Space))

	if len(output) > 0 {
		fmt.Print(string(output))
	}

	if err != nil {
		return err
	}

	fmt.Printf("%s %s %s\n", strconst.EmojiSuccess, cmd, strings.Join(args, strconst.Space))
	return nil
}
