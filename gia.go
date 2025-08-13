package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ffelipelimao/gia/internal/ai"
	"github.com/ffelipelimao/gia/internal/exec"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gia",
		Short: "Gia - A CLI tool for AI-driven task execution",
		Long:  `Gia is a CLI tool that uses AI to help with git commits `,
	}

	commitCmd := &cobra.Command{
		Use:     "commit",
		Aliases: []string{"c"},
		Short:   "Generate and execute a git commit with AI",
		Long:    `Generate a commit message using AI based on git diff and execute the commit.`,
		Run: func(cmd *cobra.Command, args []string) {
			httpClient := &http.Client{
				Timeout: 15 * time.Second,
			}
			ai, err := ai.NewIA(httpClient)
			if err != nil {
				log.Fatal(err)
			}

			executor := exec.NewExecutor(ai)

			// Interactive loop for message generation
			for {
				msg, err := executor.Start()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("\n📝 Generated commit message:\n%s\n\n", msg)
				fmt.Print("Options:\n")
				fmt.Print("  [a] Accept and commit\n")
				fmt.Print("  [r] Regenerate message\n")
				fmt.Print("  [e] Edit manually\n")
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
					// Execute the commit
					err = executor.Commit(msg)
					if err != nil {
						log.Fatal("Failed to commit:", err)
					}
					fmt.Println("✅ Commit executed successfully!")
					return

				case "r", "regenerate":
					fmt.Println("🔄 Regenerating message...")
					continue

				case "e", "edit":
					fmt.Print("Enter your commit message: ")
					customMsg, err := reader.ReadString('\n')
					if err != nil {
						log.Fatal("Failed to read custom message:", err)
					}
					customMsg = strings.TrimSpace(customMsg)

					if customMsg == "" {
						fmt.Println("❌ Message cannot be empty. Please try again.")
						continue
					}

					err = executor.Commit(customMsg)
					if err != nil {
						log.Fatal("Failed to commit:", err)
					}
					fmt.Println("✅ Commit executed successfully!")
					return

				case "q", "quit":
					fmt.Println("👋 Goodbye!")
					return

				default:
					fmt.Println("❌ Invalid option. Please choose 'a', 'r', 'e', or 'q'.")
				}
			}
		},
	}

	rootCmd.AddCommand(commitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
