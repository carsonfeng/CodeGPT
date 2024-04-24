package git

import (
	"embed"
	"log"

	"github.com/carsonfeng/ZCode/util"
)

//go:embed templates/*
var files embed.FS

const (
	HookPrepareCommitMessageTemplate = "prepare-commit-msg"
	CommitMessageTemplate            = "commit-msg.tmpl"
)

func init() {
	if err := util.LoadTemplates(files); err != nil {
		log.Fatal(err)
	}
}
