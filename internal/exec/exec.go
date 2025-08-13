package exec

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

var ErrEmptyGitDiff = errors.New("[Executor] - Empty git diff")

type AI interface {
	Execute(diff string) (string, error)
}

type Executor struct {
	ai AI
}

func NewExecutor(ai AI) *Executor {
	return &Executor{
		ai: ai,
	}
}

func (e *Executor) Start() (string, error) {
	gitDiff, err := e.getGitDiff()
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(gitDiff) == "" {
		return "", ErrEmptyGitDiff
	}

	msg, err := e.ai.Execute(gitDiff)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (e *Executor) Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	return err
}

func (e *Executor) getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}
