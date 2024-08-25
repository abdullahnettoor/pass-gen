/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/abdullahnettoor/pass-gen/app/models/req"
	"github.com/abdullahnettoor/pass-gen/app/usecase"
	"github.com/abdullahnettoor/pass-gen/app/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save key and password",
	Long:  `Save password by encrypting based on a key provided`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			credential req.Credential
			err        error
		)

		credential.UserID, err = utils.ValidateToken()
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Enter key:")
		credential.Key, _ = reader.ReadString('\n')
		if credential.Key = strings.TrimSpace(credential.Key); len(credential.Key) == 0 {
			log.Fatal("key is empty")
		}

		//read password
		fmt.Printf("Enter password to save:")
		password, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if credential.Secret = strings.TrimSpace(string(password)); len(credential.Secret) <= 1 {
			log.Fatal("secret is empty")
		}
		fmt.Println(credential)

		err = usecase.StoreSecret(&credential)
		if err != nil {
			log.Fatal("err ", err)
		}
		fmt.Printf("%s stored succesfully ", credential.Key)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
