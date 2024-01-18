package wd_flag

const (
	// TimeFormatDefault
	//  default time format
	TimeFormatDefault = "2006-01-02-03-04-05"

	// NameCliCi
	//  Provides the ci flag. This value is `woodpecker` when the is run in woodpecker
	NameCliCi = "build.name_cli"

	// EnvKeyNameCi
	//  Provides the ci flag. This value is `woodpecker` when the is run in woodpecker
	EnvKeyNameCi = "CI"

	// NameCliCiWorkspace
	//  Provides the Path of the workspace where source code gets cloned to
	NameCliCiWorkspace = "build.ci.workspace"

	// EnvKeyNameCiWorkspace
	//  Provides the Path of the workspace where source code gets cloned to
	EnvKeyNameCiWorkspace = "CI_WORKSPACE"
)
