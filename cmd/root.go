package cmd

import (
	"fmt"
	"os"

	"github.com/bigbird023/fortify-json-parser-to-excel/converter"
	"github.com/bigbird023/fortify-json-parser-to-excel/parser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cfgInput string
var cfgOutput string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fortifyjsonparsertoexcel",
	Short: "A command line application to parse and convert a fortify json document's issues to excel",
	Long: `A command line application to parse and convert a fortify json document's issues to excel. Example Usage:

	fortifyjsonparsertoexcel {inputfile} {outputfile}
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fxp := parser.NewFortifyJSONParser()
		c := converter.NewConverter(cfgInput, cfgOutput, fxp)
		err := c.Convert()
		if err != nil {
			fmt.Print(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/.fortifyjsonparsertoexcel.yaml", "config file (default is $HOME/.fortifyjsonparsertoexcel.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgInput, "input", "", "config file (default is blank)")
	rootCmd.PersistentFlags().StringVar(&cfgOutput, "output", "", "config file (default is blank)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".fortifyjsonparsertoexcel" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fortifyjsonparsertoexcel")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
