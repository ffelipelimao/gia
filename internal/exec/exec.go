package exec

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var ErrEmptyGitDiff = errors.New("[Executor] - Empty git diff")

type AI interface {
	Execute(diff, operation string) (string, error)
}

type Executor struct {
	ai AI
}

func NewExecutor(ai AI) *Executor {
	return &Executor{
		ai: ai,
	}
}

func (e *Executor) StartCommit() (string, error) {
	gitDiff, err := e.getGitDiff()
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(gitDiff) == "" {
		return "", ErrEmptyGitDiff
	}

	msg, err := e.ai.Execute(gitDiff, "commit")
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

func (e *Executor) StartBranch() (string, error) {
	gitDiff, err := e.getGitDiff()
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(gitDiff) == "" {
		return "", ErrEmptyGitDiff
	}

	msg, err := e.ai.Execute(gitDiff, "branch")
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (e *Executor) Branch(msg string) error {
	log.Printf("Tentando criar branch: %s", msg)

	cmd := exec.Command("git", "checkout", "-b", msg)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Stdout: %s", out.String())
		log.Printf("Stderr: %s", stderr.String())
		return fmt.Errorf("erro ao criar branch: %v", err)
	}

	log.Printf("Branch criada com sucesso: %s", msg)
	return nil
}

func (e *Executor) getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}
