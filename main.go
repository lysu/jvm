package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"github.com/lysu/jvm/classpath"
)

var javaCmd *cobra.Command
var cmdConfig CmdConfig

type CmdConfig struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	jreOption   string
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
			if len(args) > 0 {
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
	javaCmd.Flags().StringVar(&cmdConfig.jreOption, "Xjre", "", "path to jre")
	javaCmd.Execute()
}

func startVM(cmd *CmdConfig) {
	cp := classpath.Parse(cmd.jreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v jre:%v\n",
		cp, cmd.class, cmd.args, cmd.jreOption)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}
