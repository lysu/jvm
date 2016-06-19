package main

import (
	"fmt"
	"github.com/lysu/jvm/classpath"
	"github.com/lysu/jvm/instructions"
	"github.com/lysu/jvm/instructions/base"
	"github.com/lysu/jvm/rtda"
	"github.com/lysu/jvm/rtda/heap"
	"github.com/spf13/cobra"
	"strings"
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

	classLoader := heap.NewClassLoader(cp)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
