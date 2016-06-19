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
	helpFlag         bool
	versionFlag      bool
	cpOption         string
	jreOption        string
	class            string
	args             []string
	verboseClassFlag bool
	verboseInstFlag  bool
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

	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
