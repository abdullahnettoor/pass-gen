/*
Copyright © 2024 Abdullah Nettoor abdullahnettoor@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"log"

	"github.com/abdullahnettoor/pass-gen/app/config"
	"github.com/abdullahnettoor/pass-gen/app/db"
	"github.com/abdullahnettoor/pass-gen/app/repo"
	"github.com/abdullahnettoor/pass-gen/app/usecase"
	"github.com/abdullahnettoor/pass-gen/app/utils"
	"github.com/abdullahnettoor/pass-gen/cmd"
)

// main initializes the application by:
// 1. Loading configuration from environment variables
// 2. Establishing database connection
// 3. Setting up dependencies
// 4. Starting the CLI command handler
func main() {
	// Initialize configuration from environment variables
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("error from config ", err)
	}

	// Initialize database connection
	DB, err := db.InitDB(config.DbConnectionString)
	if err != nil {
		log.Fatal("error during connecting to database")
	}

	// Load configurations into different packages
	utils.LoadConfig(config)
	usecase.LoadConfig(config)
	repo.InitRepository(DB)

	// Start handling CLI commands
	cmd.Execute()
}
