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
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "User login",
	Long:  `Log in to an existing user account by providing your username and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			user = req.User{}
		)
		reader := bufio.NewReader(os.Stdin)

		//read user name
		fmt.Printf("Enter your username: ")
		user.UserName, _ = reader.ReadString('\n')
		if user.UserName = strings.TrimSpace(user.UserName); len(user.UserName) == 0 {
			log.Fatal("user name is empty")
		}

		if user.UserName = strings.TrimSpace(user.UserName); len(user.UserName) < 4 {
			log.Fatal("username should not be lesser that 4 charachters")
		}

		//read password
		fmt.Printf("Enter your password: ")
		password, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if user.Password = strings.TrimSpace(string(password)); len(user.Password) <= 4 {
			log.Fatal("Password must be more than 4 characters")
		}

		err = usecase.Login(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("login successfully %s \n", user.UserName)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

}
