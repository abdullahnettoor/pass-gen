/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	e "github.com/abdullahnettoor/pass-gen/app/models/errors"
	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/usecase"
	"github.com/abdullahnettoor/pass-gen/app/utils"
	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List all stored keys",
	Long:  `Retrieve and display all keys stored in the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			req req.GetKey
			err error
		)

		//validate user token
		req.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		// fetch all keys
		result, err := usecase.AllKey(req)
		if errors.Is(err, e.ErrIsEmpty) {
			fmt.Println("No Secrets are stored")
			return
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, val := range result.Name {
			fmt.Printf("key :%s\n", val)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}
