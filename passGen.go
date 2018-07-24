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
	"flag"
	"fmt"
	"math"
	"os"

	"golang.org/x/crypto/sha3"
)

const (
	postgres  = 30
	mysql     = 41
	sqlServer = 128
	mariaDB   = 41
	aurora    = 38
	oracle    = 30
)

// sourced from github.com/pivotal-cf/aws-services-broker/brokers/rds/internal/sql/generators.go

func generatePassword(salt, id string, maxIdentifierLength float64) string {
	bytes := []byte(id + salt)
	digest := sha3.Sum224(bytes)
	result := base64.RawURLEncoding.EncodeToString(digest[:])
	length := int(math.Min(maxIdentifierLength, float64(len(result))))
	return result[:length]
}

func printHelp() {
	helpDoc := `
This tool regenerates the original masterusername password for an AWS Relational
Database Service PostgreSQL Instance originally created by the PCF Service
Broker for AWS provided a service instance guid and master salt key.

Please see Pivotal knowledge base on this tool here:
https://discuss.pivotal.io/hc/en-us/articles/360001356494

Usage: passGen -i [identity] -s [salt]

	-m
		Construct a MySQL instance password.

Standard security recommendations apply to distribution of the generated
password.

This tool is provided as a general service and is not under any official
supported capacity. There is no implied or guaranteed warranty or statement of
support.

Released under MIT license,	copyright 2018 Tyler Ramer
	`
	fmt.Println(helpDoc)

}

func main() {

	// sourced from github.com/pivotal-cf/aws-services-broker/brokers/rds/internal/postgres.go
	var maxIdentifierLength = 30

	helpFlag := flag.Bool("h", false, "help")
	idFlag := flag.String("i", "", "ID")
	saltFlag := flag.String("s", "", "Salt")
	mysqlFlag := flag.Bool("m", false, "mysql")
	flag.Parse()

	switch {
	case *helpFlag:
		printHelp()
		os.Exit(1)
	case *idFlag == "":
		printHelp()
		os.Exit(1)
	case *saltFlag == "":
		printHelp()
		os.Exit(1)
	}

	id := *idFlag
	salt := *saltFlag

	if *mysqlFlag {
		maxIdentifierLength = 41
	}

	passwd := generatePassword(salt, id, float64(maxIdentifierLength))
	fmt.Println("Generated password:")
	fmt.Println(passwd)

}
