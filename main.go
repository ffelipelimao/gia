package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/kljensen/snowball" // Biblioteca ativa para stemming
)

const (
	maxDiffLines = 500
)

func main() {
	fmt.Println("🤖 TinyCommit - Gerador Inteligente de Mensagens")

	if !isGitRepository() {
		fmt.Println("Erro: Diretório não é um repositório Git")
		os.Exit(1)
	}

	diff, err := getStagedDiff()
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
		os.Exit(1)
	}

	if diff == "" {
		fmt.Println("Nenhuma alteração preparada para commit")
		os.Exit(0)
	}

	showSmartDiff(diff)
	message := generateSmartMessage(diff)

	fmt.Printf("\n💡 Sugestão:\n%s\n", message)
	fmt.Print("\nAceitar? (S/n/editar): ")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	switch strings.ToLower(input) {
	case "n":
		fmt.Print("Digite sua mensagem: ")
		message, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	case "e", "editar":
		fmt.Printf("Editar mensagem (%s): ", message)
		message, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	}

	message = strings.TrimSpace(message)
	if message != "" {
		if err := exec.Command("git", "commit", "-m", message).Run(); err != nil {
			fmt.Printf("Erro no commit: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✅ Commit criado!")
	}
}

func isGitRepository() bool {
	return exec.Command("git", "rev-parse", "--is-inside-work-tree").Run() == nil
}

func getStagedDiff() (string, error) {
	out, err := exec.Command("git", "diff", "--cached", "--color=always").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git diff falhou: %v", err)
	}
	return string(out), nil
}

func showSmartDiff(diff string) {
	files := make(map[string]struct{})
	changes := struct{ add, del int }{}

	fmt.Println("\n📋 Alterações detectadas:")
	lines := strings.Split(diff, "\n")
	for i := 0; i < len(lines) && i < maxDiffLines; i++ {
		line := lines[i]

		switch {
		case strings.HasPrefix(line, "+++ b/"):
			file := strings.TrimPrefix(line, "+++ b/")
			files[file] = struct{}{}
			fmt.Printf("\n📄 %s\n", file)
		case strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "++"):
			changes.add++
			fmt.Println(line)
		case strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "--"):
			changes.del++
			fmt.Println(line)
		}
	}

	fmt.Printf("\n📊 Resumo: %d arquivos, +%d -%d linhas\n", len(files), changes.add, changes.del)
}

func generateSmartMessage(diff string) string {
	// Extrai palavras importantes
	words := extractKeywords(diff)

	// Classifica o tipo de commit
	commitType := classifyCommit(diff)

	// Gera a mensagem base
	msg := fmt.Sprintf("%s: ", commitType)

	// Adiciona palavras-chave relevantes
	if len(words) > 0 {
		msg += strings.Join(words[:min(3, len(words))], ", ") + " "
	}

	// Adiciona estatísticas
	stats := countChanges(diff)
	msg += fmt.Sprintf("(+%d/-%d)", stats.add, stats.del)

	return msg
}

func extractKeywords(diff string) []string {
	wordMap := make(map[string]bool)
	r := regexp.MustCompile(`\+[^+].*\b([a-zA-Z]{4,})\b`)

	matches := r.FindAllStringSubmatch(diff, -1)
	for _, m := range matches {
		if len(m) > 1 {
			word := strings.ToLower(m[1])
			stemmed, _ := snowball.Stem(word, "english", true)
			if len(stemmed) > 3 && !isCommonWord(stemmed) {
				wordMap[stemmed] = true
			}
		}
	}

	words := make([]string, 0, len(wordMap))
	for w := range wordMap {
		words = append(words, w)
	}

	return words
}

func isCommonWord(word string) bool {
	common := map[string]bool{
		"func": true, "return": true, "class": true,
		"method": true, "value": true, "object": true,
	}
	return common[word]
}

func classifyCommit(diff string) string {
	stats := countChanges(diff)

	switch {
	case stats.add > 50 && stats.del < 10:
		return "Adicionar"
	case stats.del > stats.add && stats.del > 20:
		return "Remover"
	case strings.Contains(diff, "fix") || strings.Contains(diff, "bug"):
		return "Corrigir"
	case stats.add < 10 && stats.del < 10:
		return "Atualizar"
	default:
		return "Modificar"
	}
}

func countChanges(diff string) (stats struct{ add, del int }) {
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "++"):
			stats.add++
		case strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "--"):
			stats.del++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
