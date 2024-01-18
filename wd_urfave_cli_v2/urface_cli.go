package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
)

// WoodpeckerUrfaveCliFlags
// woodpecker urfave cli flag will append all cli.Flag for Built-in environment variables when woodpecker ci running
func WoodpeckerUrfaveCliFlags() []cli.Flag {
	return UrfaveCliAppendCliFlags(
		basicFlag(),
		repositoryFlag(),
		currentCommitFlag(),
		currentPipelineFlag(),
		currentWorkflowFlag(),
		currentStepFlag(),
		previousCommitFlag(),
		previousPipelineFlag(),
		ciSystemFlag(),
		ciForgeFlag(),
	)
}

// UrfaveCliAppendCliFlag
//
//	append cli.Flag
func UrfaveCliAppendCliFlag(target []cli.Flag, elem []cli.Flag) []cli.Flag {
	if len(elem) == 0 {
		return target
	}

	return append(target, elem...)
}

// UrfaveCliAppendCliFlags
//
//	append cli.Flag
func UrfaveCliAppendCliFlags(target []cli.Flag, elems ...[]cli.Flag) []cli.Flag {
	if len(elems) == 0 {
		return target
	}

	for _, elem := range elems {
		if len(elem) > 0 {
			target = UrfaveCliAppendCliFlag(target, elem)
		}
	}
	return target
}
