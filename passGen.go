/*
Password generator for AWS RDS master user

Command execution:
passGen -i [identity] -s [salt]

Please see relevant Pivotal KB here:



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

// sourced from github.com/pivotal-cf/aws-services-broker/brokers/rds/internal/postgres.go
const maxIdentifierLength = 30

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
This tool generates a password for AWS Relational Database Service master user,
provided an id and master salt key. Please see Pivotal knowledge base on this
tool here:

Usage: passGen -i [identity] -s [salt]

Standard security recommadations apply to distribution of the generated
password.

This tool is provided as a general service and is not under any official
supported capacity. There is no implied or guarenteed warranty or statement of
support.

Released under MIT license,	copyright 2018 Tyler Ramer
	`
	fmt.Println(helpDoc)

}

func main() {

	helpFlag := flag.Bool("h", false, "help")
	idFlag := flag.String("i", "", "ID")
	saltFlag := flag.String("s", "", "Salt")
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

	passwd := generatePassword(salt, id, maxIdentifierLength)
	fmt.Println("Generated password:")
	fmt.Println(passwd)

}
