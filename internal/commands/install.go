package commands

import (
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Install = &cobra.Command{
	Use:   "install",
	Short: "build all .go files",
	Run:   install,
}

func install(cmd *cobra.Command, args []string) {
	out, err := exec.Command("env", "GOOS=js", "GOARCH=wasm", "go", "install").Output()
	if err != nil {
		color.Red("err: ", err)
	}
	color.Green(string(out[:]))

}
