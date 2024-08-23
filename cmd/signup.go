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

// signupCmd represents the signup command
var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Register a new user account",
	Long:  `Create a new user account by providing the necessary credentials.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			user = req.User{}
		)
		reader := bufio.NewReader(os.Stdin)

		//read user name
		fmt.Printf("Enter your user name :")
		user.UserName, _ = reader.ReadString('\n')
		if user.UserName = strings.TrimSpace(user.UserName); len(user.UserName) == 0 {
			log.Fatal("user name is empty")
		}

		if user.Password = strings.TrimSpace(string(user.UserName)); len(user.Password) <= 4 {
			log.Fatal("password is less than five digit kindly strong your the password")
		}

		//read password
		fmt.Printf("Enter your password :")
		password, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if user.Password = strings.TrimSpace(string(password)); len(user.Password) <= 4 {
			log.Fatal("password is less than five digit kidly strong your the password")
		}

		//read confirm password
		fmt.Printf("Re-enter your password :")
		password, err = term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("")

		if user.ConfirmPassword = strings.TrimSpace(string(password)); len(user.ConfirmPassword) == 0 {
			log.Fatal("confirm password is emtpy kindly enter the password")
		}

		result, err := usecase.Signup(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("signup succesfully with user id ", result)
	},
}

func init() {
	rootCmd.AddCommand(signupCmd)
}
