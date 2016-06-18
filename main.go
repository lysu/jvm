package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"github.com/lysu/jvm/classpath"
	"github.com/lysu/jvm/classfile"
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
	cf := loadClass(className, cp)
	fmt.Println(cmdConfig.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _ , err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
