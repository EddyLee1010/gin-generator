package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var updateSelfCmd = &cobra.Command{
	Use:   "update",
	Short: "更新工具",
	Run: func(cmd *cobra.Command, args []string) {
		updateSelf()
	},
}

func init() {
	rootCmd.AddCommand(updateSelfCmd)

}
func updateSelf() {
	cmd := exec.Command("go", "install", "github.com/eddylee1010/gin-generator@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ update failed: %v\n", err)
		return
	}

	fmt.Println("✅ update success!")
}
