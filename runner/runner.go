package runner

import (
	"context"
	"log"
	"os/exec"
	"time"
)

type Runner struct {
	ch <-chan string
}

func NewRunner(ch <-chan string) Runner {
	return Runner{
		ch: ch,
	}
}

func (r *Runner) Start() {
	for scriptPath := range r.ch {
		r.scriptRuntime(scriptPath)
	}
}

func (r *Runner) scriptRuntime(scriptPath string) {
	// TODO: make the timeout variable
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("script failed: %v\noutput: %s", err, output)
		return
	}
	log.Printf("script executed successfully:\n%s", output)
}
