package commands

import (
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get GoLang depedencies",
	Run:   getDepedencies,
}

func getDepedencies(cmd *cobra.Command, args []string) {
	out, err := exec.Command("env", "GOOS=js", "GOARCH=wasm", "go", "get").Output()
	if err != nil {
		color.Red("err: ", err)
	}
	color.Green(string(out[:]))

}
