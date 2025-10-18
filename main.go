package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/jtizdev/womba-go/client"
)

const version = "1.0.0"

func main() {
	// Define subcommands
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	storyKey := generateCmd.String("story", "", "Jira story key (required)")
	upload := generateCmd.Bool("upload", false, "Upload tests to Zephyr")

	healthCmd := flag.NewFlagSet("health", flag.ExitOnError)

	versionCmd := flag.NewFlagSet("version", flag.ExitOnError)

	// Check for subcommand
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Get environment variables
	apiURL := os.Getenv("WOMBA_API_URL")
	apiKey := os.Getenv("WOMBA_API_KEY")

	if apiURL == "" {
		color.Red("❌ Error: WOMBA_API_URL environment variable not set")
		fmt.Println("\nSet it with: export WOMBA_API_URL=https://womba-api.up.railway.app")
		os.Exit(1)
	}

	if apiKey == "" && os.Args[1] != "version" {
		color.Red("❌ Error: WOMBA_API_KEY environment variable not set")
		fmt.Println("\nSet it with: export WOMBA_API_KEY=your-api-key")
		os.Exit(1)
	}

	// Parse subcommand
	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])

		if *storyKey == "" {
			color.Red("❌ Error: -story flag is required")
			generateCmd.PrintDefaults()
			os.Exit(1)
		}

		generateTests(apiURL, apiKey, *storyKey, *upload)

	case "health":
		healthCmd.Parse(os.Args[2:])
		checkHealth(apiURL, apiKey)

	case "version":
		versionCmd.Parse(os.Args[2:])
		printVersion()

	default:
		color.Red("❌ Unknown command: %s", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Womba CLI - Go Client")
	fmt.Println("\nUsage:")
	fmt.Println("  womba <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  generate    Generate test cases for a Jira story")
	fmt.Println("  health      Check API health")
	fmt.Println("  version     Show version")
	fmt.Println("\nExamples:")
	fmt.Println("  womba generate -story PLAT-12991")
	fmt.Println("  womba generate -story PLAT-12991 -upload")
	fmt.Println("  womba health")
	fmt.Println("\nEnvironment Variables:")
	fmt.Println("  WOMBA_API_URL    Womba API base URL")
	fmt.Println("  WOMBA_API_KEY    Womba API authentication key")
}

func printVersion() {
	color.Cyan("Womba Go CLI v%s", version)
}

func generateTests(apiURL, apiKey, storyKey string, upload bool) {
	color.Cyan("🚀 Generating tests for %s...\n", storyKey)

	wombaClient := client.NewWombaClient(apiURL, apiKey)

	result, err := wombaClient.GenerateTests(storyKey, upload)
	if err != nil {
		color.Red("❌ Generation failed: %v", err)
		os.Exit(1)
	}

	// Print results
	color.Green("\n✅ Successfully generated %d test cases!", len(result.TestCases))
	color.Cyan("📊 Quality Score: %.1f/100", result.QualityScore)
	color.Cyan("📁 Suggested Folder: %s", result.SuggestedFolder)
	color.Cyan("⏱️  Execution Time: %.2fs", result.ExecutionTimeSeconds)

	if result.Metadata != nil {
		if aiModel, ok := result.Metadata["ai_model"].(string); ok {
			color.Cyan("🤖 AI Model: %s", aiModel)
		}
		if testCount, ok := result.Metadata["test_count"].(float64); ok {
			color.Cyan("📝 Test Count: %.0f", testCount)
		}
	}

	// Print test cases
	fmt.Println("\n" + color.YellowString("Generated Test Cases:"))
	fmt.Println(color.YellowString(strings.Repeat("=", 80)))

	for i, testCase := range result.TestCases {
		fmt.Printf("\n%s\n", color.CyanString("%d. %s", i+1, testCase.Title))
		fmt.Printf("   Priority: %s | Type: %s\n", testCase.Priority, testCase.TestType)
		fmt.Printf("   Description: %s\n", testCase.Description)
		fmt.Printf("   Steps: %d\n", len(testCase.Steps))
	}

	// Print Zephyr IDs if uploaded
	if upload && len(result.ZephyrIDs) > 0 {
		fmt.Println("\n" + color.GreenString("✅ Uploaded to Zephyr:"))
		for i, zephyrID := range result.ZephyrIDs {
			fmt.Printf("   %d. %s\n", i+1, zephyrID)
		}
	}

	fmt.Println("\n" + color.GreenString("🎉 Done!"))
}

func checkHealth(apiURL, apiKey string) {
	color.Cyan("🔍 Checking API health...\n")

	wombaClient := client.NewWombaClient(apiURL, apiKey)

	result, err := wombaClient.HealthCheck()
	if err != nil {
		color.Red("❌ Health check failed: %v", err)
		os.Exit(1)
	}

	color.Green("✅ API is healthy!")

	if status, ok := result["status"].(string); ok {
		fmt.Printf("Status: %s\n", status)
	}

	if version, ok := result["version"].(string); ok {
		fmt.Printf("Version: %s\n", version)
	}

	if deps, ok := result["dependencies"].(map[string]interface{}); ok {
		fmt.Println("\nDependencies:")
		for name, status := range deps {
			statusStr := status.(string)
			if statusStr == "connected" {
				color.Green("  ✅ %s: %s", name, statusStr)
			} else {
				color.Yellow("  ⚠️  %s: %s", name, statusStr)
			}
		}
	}
}

