/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/usecase"
	"github.com/abdullahnettoor/pass-gen/app/utils"
	"github.com/spf13/cobra"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Retrieve passwords with key",
	Long:  `Fetch passwords by passing key as command arguments with the flay --key`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			req req.GetSecretPassword
			err error
		)

		req.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		req.Key, err = cmd.Flags().GetString("key")
		if err != nil {
			log.Fatal(err)
		}
		if req.Key = strings.TrimSpace(req.Key); len(req.Key) == 0 {
			log.Fatal("password key is empty")
		}

		result, err := usecase.GetSecretPasswords(&req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("value : %s\n", result.SecretPlainText)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.PersistentFlags().String("key", "", "Retrieve a single password")
}
