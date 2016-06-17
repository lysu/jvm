package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var javaCmd *cobra.Command
var cmdConfig CmdConfig

type CmdConfig struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func main() {
	javaCmd = &cobra.Command{
		Use:   "java",
		Short: "java to run jvm",
		Run: func(cmd *cobra.Command, args []string) {
			if cmdConfig.versionFlag {
				fmt.Println("version 0.0.1")
				return
			}
			if len(args) > 1 {
				cmdConfig.class = args[0]
				cmdConfig.args = args[1:]
			}
			if cmdConfig.helpFlag || cmdConfig.class == "" {
				fmt.Println(javaCmd.UsageString())
				return
			}
			startVM(&cmdConfig)
		},
	}
	javaCmd.Flags().BoolVar(&cmdConfig.helpFlag, "help", false, "print help message")
	javaCmd.Flags().BoolVar(&cmdConfig.helpFlag, "?", false, "print help message")
	javaCmd.Flags().BoolVar(&cmdConfig.versionFlag, "version", false, "print version and exit")
	javaCmd.Flags().StringVarP(&cmdConfig.cpOption, "classpath", "c", "", "classpath")
	javaCmd.Execute()
}

func startVM(cmd *CmdConfig) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
