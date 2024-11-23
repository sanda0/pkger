package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{Use: "pack"}
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initializing pack and create pack.congifg.json",
		Run: func(cmd *cobra.Command, args []string) {
			createConfig()
		},
	}

	newPkgCmd := &cobra.Command{
		Use:   "new",
		Short: "Create new package",
		Run: func(cmd *cobra.Command, args []string) {
			pkgName, _ := cmd.Flags().GetString("package-name")
			authorName, _ := cmd.Flags().GetString("author-name")
			authorEmail, _ := cmd.Flags().GetString("autor-email")
			param := NewPkgFlags{
				AuthorName:  authorName,
				AuthorEmail: authorEmail,
				PkgName:     pkgName,
			}
			newPkgForm(param)
		},
	}
	newPkgCmd.Flags().StringP("package-name", "n", "", "Package name")
	newPkgCmd.Flags().StringP("author-email", "e", "", "Author email")
	newPkgCmd.Flags().StringP("author-name", "N", "", "Author name")

	makeCmd := &cobra.Command{
		Use:   "make",
		Short: "Make something",
		Run: func(cmd *cobra.Command, args []string) {
			migration, _ := cmd.Flags().GetBool("migration")
			controller, _ := cmd.Flags().GetBool("controller")
			model, _ := cmd.Flags().GetBool("model")
			mcr, _ := cmd.Flags().GetBool("mcr")

			fmt.Printf("Making with flags - migration: %v, controller: %v, model: %v, mcr: %v\n",
				migration, controller, model, mcr)
		},
	}
	makeCmd.Flags().StringP("migration", "m", "", "Create migration")
	makeCmd.Flags().StringP("controller", "c", "", "Create controller")
	makeCmd.Flags().StringP("model", "M", "", "Create model")
	makeCmd.Flags().StringP("mcr", "a", "", "Create migration,model,controller")
	makeCmd.Flags().StringP("package", "p", "", "Package name")

	rootCmd.AddCommand(initCmd, newPkgCmd, makeCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}

}
