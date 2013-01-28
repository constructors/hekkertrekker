package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commands = []string{"start", "deliver", "done"}

func bye(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	os.Exit(1)
}

func start() {
	var selection int
	var tryBranchName, branchName string

	fmt.Printf("Fetching story list...")
	stories := pt.AcceptableStories()
	fmt.Printf("done\n")

	if len(stories) == 0 {
		bye("No stories to accept!\n")
	}

	colors := map[string]int{
		"bug":     91,
		"feature": 92,
	}
	for i := range stories {
		fmt.Printf("% 3d \033[%dm%s\033[0m (%d pts)\n",
			i, colors[stories[i].Type],
			stories[i].Name, stories[i].Estimate)
	}

	fmt.Printf("Please choose your story number [0]: ")
	fmt.Scanf("%d", &selection)

	if selection >= len(stories) || selection < 0 {
		start()
		return
	}

	fmt.Printf("Accepting story...")
	pt.AcceptStory(stories[selection].ID)
	fmt.Printf("done\n")

	tryBranchName = strings.ToLower(stories[selection].Name)
	tryBranchName = strings.TrimSpace(tryBranchName)
	tryBranchName = strings.Replace(tryBranchName, " ", "_", -1)

	fmt.Printf("New branch name [%s]: ", tryBranchName)
	in := bufio.NewReader(os.Stdin)
	branchName, _ = in.ReadString('\n')
	branchName = strings.TrimSpace(branchName)

	if branchName == "" {
		branchName = tryBranchName
	} else {
		branchName = strings.ToLower(branchName)
		branchName = strings.Replace(branchName, " ", "_", -1)
	}

	branchName = fmt.Sprintf("%s_%d", branchName, stories[selection].ID)

	fmt.Printf("Creating branch %s from default...\n", branchName)
	hgUpdate("default")
	hgNewBranch(branchName)
	hgCommit(fmt.Sprintf(generalConfig.NewBranchCommitMsg,
		stories[selection].ID))
	fmt.Printf("Pushing upstream...\n%s\n", hgPushNewBranch())
}

func deliver() {
	currentBranch := hgBranch()
	currentBranchItems := strings.Split(currentBranch, "_")
	currentStory, err := strconv.Atoi(currentBranchItems[len(currentBranchItems)-1])

	if err != nil || currentStory == 0 {
		bye("This doesn't seem to be a story branch\n")
	}

	fmt.Printf("Delivering story %d...", currentStory)
	pt.DeliverStory(currentStory)
	fmt.Printf("done\n")

	fmt.Printf("Merging to branch %s...\n", repositoryConfig.StagingBranch)
	hgUpdate(repositoryConfig.StagingBranch)
	hgMerge(currentBranch)
	hgCommit(fmt.Sprintf(generalConfig.DeliverCommitMsg,
		currentStory))
	fmt.Printf("Pushing upstream...\n%s\n", hgPush())
	hgUpdate(currentBranch)
}

func done() {
	currentBranch := hgBranch()
	currentBranchItems := strings.Split(currentBranch, "_")
	currentStory, err := strconv.Atoi(currentBranchItems[len(currentBranchItems)-1])

	if err != nil || currentStory == 0 {
		bye("This doesn't seem to be a story branch\n")
	}

	fmt.Printf("Adding label...\n")
	if pt.DoneStory(currentStory) {
		fmt.Printf("Closing branch...\n")
		hgCloseBranch(generalConfig.CloseCommitMsg)
		fmt.Printf("Merging to branch default...\n")
		hgUpdate("default")
		hgMerge(currentBranch)
		hgCommit(fmt.Sprintf(generalConfig.DoneCommitMsg,
			currentStory))
		fmt.Printf("Pushing upstream...\n%s\n", hgPush())
	} else {
		bye("This story isn't accepted yet!\n")
	}
}

func usage() {
	bye("Usage: %s [%s]\n", os.Args[0], strings.Join(commands, "|"))
}

func main() {
	var matchedCommands []string

	initConfig()
	initPivotalTracker()

	if len(os.Args) < 2 {
		usage()
	}

	for i := range commands {
		if strings.HasPrefix(commands[i], os.Args[1]) {
			matchedCommands = append(matchedCommands, commands[i])
		}
	}

	if len(matchedCommands) != 1 {
		usage()
	}

	switch matchedCommands[0] {
	case "start":
		start()
	case "deliver":
		deliver()
	case "done":
		done()
	default:
		usage()
	}
}
