package cmd

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
	maxIdentifierLength int
)

var (
	mysqlFlag    bool
	postgresFlag bool
)

const (
	postgres = 41
	mysql    = 30
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&id, "identity", "i", "", "Service instance identity")
	rootCmd.PersistentFlags().StringVarP(&salt, "salt", "s", "", "Master salt key")
	rootCmd.PersistentFlags().BoolVar(&mysqlFlag, "mysql", false, "MySQL RDS instance")
	rootCmd.PersistentFlags().BoolVar(&postgresFlag, "postgres", false, "Postgres RDS instance")

}

var rootCmd = &cobra.Command{
	Use:   "passGen [-flags]",
	Short: "passGen is used to display the original masterusername password for an AWS Relational Database Service",
	Long: `This tool displays the original masterusername password for an AWS Relational
	Database Service Instance originally created by the PCF Service
	Broker for AWS provided a service instance guid and master salt key.
	
	Please see Pivotal knowledge base on this tool here:
	https://discuss.pivotal.io/hc/en-us/articles/360001356494
	
	Usage: passGen -i [identity] -s [salt]  --
	
	Standard security recommendations apply to distribution of the generated
	password.
	
	This tool is provided as a general service and is not under any official
	supported capacity. There is no implied or guaranteed warranty or statement of
	support.
	
	Released under MIT license,	copyright 2018 Tyler Ramer`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case id == "":
			cmd.Help()
			os.Exit(1)
		case salt == "":
			cmd.Help()
			os.Exit(1)
		}
		switch {
		case mysqlFlag:
			maxIdentifierLength = mysql
		case postgresFlag:
			maxIdentifierLength = postgres
		}

		passwd := generatePassword(salt, id, float64(maxIdentifierLength))
		fmt.Println("Generated password:")
		fmt.Println(passwd)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generatePassword(salt, id string, maxIdentifierLength float64) string {
	bytes := []byte(id + salt)
	digest := sha3.Sum224(bytes)
	result := base64.RawURLEncoding.EncodeToString(digest[:])
	length := int(math.Min(maxIdentifierLength, float64(len(result))))
	return result[:length]
}
