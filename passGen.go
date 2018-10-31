/*
This tool regenerates the original masterusername password for an AWS Relational
Database Service PostgreSQL Instance originally created by the PCF Service
Broker. Required ID and salt key.

Command execution:
passGen -i [identity] -s [salt]

	-m
		Construct a MySQL instance password.

Please see relevant Pivotal KB here:
https://discuss.pivotal.io/hc/en-us/articles/360001356494


Copyright 2018 Tyler Ramer

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

var (
	id                  string
	salt                string
	database			string
	maxIdentifierLength int
)

type rds struct {
	name       string
	passLength int
}

var (
	postgres  = rds{"Postgres", 30}
	mysql     = rds{"MySQL", 41}
	sqlServer = rds{"SQL-Server", 128}
	mariaDB   = rds{"MariaDB", 41}
	aurora    = rds{"Aurora DB", 38}
	oracle    = rds{"Oracle DB", 30}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&id, "identity", "i", "", "Service instance identity")
	rootCmd.PersistentFlags().StringVarP(&salt, "salt", "s", "", "Master salt key")
	rootCmd.PersistentFlags().StringVarP(&database, "database-type", "d", "", "Database Type")
}

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "passGen [-flags]",
	Short: "passGen is used to display the original masterusername password for an AWS Relational Database Service",
	Long: `This tool displays the original masterusername password for an AWS Relational
Database Service Instance originally created by the PCF Service
Broker for AWS provided a service instance guid and master salt key.

Please see Pivotal knowledge base on this tool here:
https://discuss.pivotal.io/hc/en-us/articles/360001356494

Usage: passGen -i [--identity] -s [--salt]  -d [--database-type]

Database Types Available: mysql, postgres, sqlServer, mariadb, aurora, oracle

Only one service name can be provided.

Standard security recommendations apply to distribution of the generated
password.
	
This tool is provided as a general service and is not under any official
supported capacity. There is no implied or guaranteed warranty or statement of
support.
	
Released under MIT license,	copyright 2018 Tyler Ramer`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case id == "" || salt == "":
			cmd.Help()
			os.Exit(1)
		case database == "":
			cmd.Help()
			os.Exit(1)
		}
		switch {
		case database == "mysql":
			displayPassword(mysql)
		case database == "postgres":
			displayPassword(postgres)
		case database == "sqlServer":
			displayPassword(sqlServer)
		case database == "mariadb":
			displayPassword(mariaDB)
		case database == "aurora":
			displayPassword(aurora)
		case database == "oracle":
			displayPassword(oracle)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func displayPassword(service rds) {
	pass := generatePassword(salt, id, float64(service.passLength))
	fmt.Printf("Generated %v password:\n%v\n", service.name, pass)
}

// sourced from github.com/pivotal-cf/aws-services-broker/brokers/rds/internal/sql/generators.go
func generatePassword(salt, id string, maxIdentifierLength float64) string {
	bytes := []byte(id + salt)
	digest := sha3.Sum224(bytes)
	result := base64.RawURLEncoding.EncodeToString(digest[:])
	length := int(math.Min(maxIdentifierLength, float64(len(result))))
	return result[:length]
}

func xOr(flags ...bool) bool {
	var t bool
	for _, flag := range flags {
		if flag && t {
			return false
		}
		if flag {
			t = true
		}
	}
	return t
}
