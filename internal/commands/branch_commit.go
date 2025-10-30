package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ffelipelimao/gia/internal/ai"
	"github.com/ffelipelimao/gia/internal/exec"
	"github.com/spf13/cobra"
)

func NewBranchCommitCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "branch commit",
		Aliases: []string{"bc"},
		Short:   "Create a git branch and commit message with AI",
		Long:    `Generate a git branch and commit message using AI based on git diff.`,
		Args:    cobra.MaximumNArgs(1),
		Run:     branchCommitRunner,
	}
}

func branchCommitRunner(cmd *cobra.Command, args []string) {
	provider := getProvider(args)
	strategy := ai.NewDefaultFactory()
	aiClient, err := strategy.Create(cmd.Context(), provider)
	if err != nil {
		log.Fatalf("Failed to create AI client: %v", err)
	}

	executor := exec.NewExecutor(aiClient)

	shouldGenerate := true
	var branch string
	var commit string

	for {
		if shouldGenerate {
			branch, commit, err = executor.StartBranchCommit()
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("📝 Generated branch:\n%s\n\n", branch)
		fmt.Printf("📝 Generated commit:\n%s\n\n", commit)
		fmt.Print("Options:\n")
		fmt.Print("  [a] Accept and create\n")
		fmt.Print("  [r] Regenerate branch\n")
		fmt.Print("  [q] Quit\n\n")
		fmt.Print("Choose an option: ")

		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Failed to read input:", err)
		}

		choice = strings.TrimSpace(strings.ToLower(choice))

		switch choice {
		case "a", "accept":
			err = executor.Branch(branch)
			if err != nil {
				log.Fatal("Failed to branch:", err)
			}
			err = executor.Commit(branch)
			if err != nil {
				log.Fatal("Failed to commit:", err)
			}
			fmt.Println("✅ commit and branch create successfully!")
			return

		case "r", "regenerate":
			fmt.Println("🔄 Regenerating message...")
			continue

		case "q", "quit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid option. Please choose 'a', 'r' or 'q'.")
			shouldGenerate = false
		}
	}
}
