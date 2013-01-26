package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func hgCmd(arg ...string) string {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command("hg", arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		bye("%s%v\n", stderr.String(), err)
	}

	return strings.TrimSpace(stdout.String())
}

func hgRoot() string {
	return hgCmd("root")
}

func hgUpdate(branch string) {
	hgCmd("update", branch)
}

func hgBranch() string {
	return hgCmd("branch")
}

func hgNewBranch(branch string) {
	hgCmd("branch", branch)
}

func hgCommit(msg string) {
	hgCmd("commit", "-m", msg)
}

func hgPush() string {
	return hgCmd("push")
}

func hgPushNewBranch() string {
	return hgCmd("push", "--new-branch")
}

func hgMerge(branch string) string {
	return hgCmd("merge", branch)
}
