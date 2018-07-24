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

import "github.com/Tylarb/AWS-RDS-master-user-passgen/cmd"

// sourced from github.com/pivotal-cf/aws-services-broker/brokers/rds/internal/sql/generators.go

func main() {

	cmd.Execute()
	/*
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
		fmt.Println(passwd) */

}
