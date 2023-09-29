package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jonsen/chatglm"
)

var (
	model  = flag.String("m", "", "chat model path")
	prompt = flag.String("p", "", "prompt")
	ia     = flag.Bool("i", false, "interactive mode")
)

func main() {
	flag.Parse()

	if *model == "" {
		fmt.Println("model path is required")
		flag.Usage()
		return
	}
	if _, err := os.Stat(*model); err != nil {
		fmt.Printf("model %q is not exists.\n", *model)
		return
	}

	if *ia {
		interactive()
	} else {
		do(*model, *prompt)
	}

	fmt.Println("Bye!")
}

func do(model, prompt string) {
	fmt.Println("model:", model, prompt)
	p := chatglm.New(model)
	defer p.Delete()

	if prompt == "" {
		fmt.Println("prompt is required")
		flag.Usage()
		return
	}

	out := p.Generate(prompt)
	for out != "" {
		fmt.Println("wait for respone.")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(out)
}

func interactive() {
	p := chatglm.New(*model)
	defer p.Delete()

	history := []*chatglm.Turn{}

	fmt.Println("欢迎使用 ChatGLM-GO 个人助手！")

	for {

		var prompt, answer string
		fmt.Print("\nPrompt: ")
		fmt.Scanln(&prompt)

		if prompt == "" {
			continue
		}
		if prompt == "stop" || prompt == "exit" {
			return
		}

		out := p.StreamGenerate(chatglm.BuildPrompt(prompt, history))
		fmt.Println("Answer:")
		for txt := range out {
			fmt.Print(txt)
			answer += txt
		}
		fmt.Println("")

		// 保留3个历史
		if len(history) >= 3 {
			history = append(history[:0], history[1:]...)
		}

		history = append(history, &chatglm.Turn{Question: prompt, Answer: answer})

	}
}
