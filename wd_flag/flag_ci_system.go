package wd_flag

const (
	// NameCliWoodpeckerBackend
	//  Provides the backend flag. This value is type when the is run in woodpecker, default is empty.
	//  Possible values are auto-detect, docker, local or kubernetes
	NameCliWoodpeckerBackend = "build.woodpecker.backend"

	// EnvKeyWoodpeckerBackend
	//  Provides the backend flag. This value is type when the is run in woodpecker, default is empty.
	//  Possible values are auto-detect, docker, local or kubernetes
	EnvKeyWoodpeckerBackend = "WOODPECKER_BACKEND"

	// NameCliWoodpeckerHostName
	//  Provides the configures the agent hostname. default is empty.
	NameCliWoodpeckerHostName = "build.woodpecker.hostname"

	// EnvKeyWoodpeckerHostName
	//  Provides the configures the agent hostname. default is empty.
	EnvKeyWoodpeckerHostName = "WOODPECKER_HOSTNAME"

	// NameCliWoodpeckerFilterLabels
	//  Configures labels to filter pipeline pick up.
	//  Use a list of key-value pairs like key=value,second-key=*. * can be used as a wildcard.
	//  By default, agents provide three additional labels platform=os/arch, hostname=my-agent and repo=*
	//  most this will be empty
	NameCliWoodpeckerFilterLabels = "build.woodpecker.filter.labels"

	// EnvKeyWoodpeckerFilterLabels
	//  Configures labels to filter pipeline pick up.
	//  Use a list of key-value pairs like key=value,second-key=*. * can be used as a wildcard.
	//  By default, agents provide three additional labels platform=os/arch, hostname=my-agent and repo=* .
	//  most this will be empty
	EnvKeyWoodpeckerFilterLabels = "WOODPECKER_FILTER_LABELS"

	// NameCliCiMachine
	//  Provides name of the CI machine
	NameCliCiMachine = "ci_machine.name"

	// EnvKeyCiMachine
	//  Provides name of the CI machine
	EnvKeyCiMachine = "CI_MACHINE"

	// NameCliCiSystemPlatform
	//  Provides ci system platform as: linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64
	//  The syntax of the platform is GOOS/GOARCH from https://go.dev/doc/install/source
	NameCliCiSystemPlatform = "ci_system.platform"

	// EnvKeyCiSystemPlatform
	//  Provides ci system platform as: linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64
	//  The syntax of the platform is GOOS/GOARCH from https://go.dev/doc/install/source
	EnvKeyCiSystemPlatform = "CI_SYSTEM_PLATFORM"

	// NameCliCiSystemName
	//  Provides name of the CI system: woodpecker
	NameCliCiSystemName = "ci_system.name"

	// EnvKeyCiSystemName
	//  Provides name of the CI system: woodpecker
	EnvKeyCiSystemName = "CI_SYSTEM_NAME"

	// NameCliCiSystemUrl
	//  Provides link to CI system
	NameCliCiSystemUrl = "ci_system.url"

	// EnvKeyCiSystemUrl
	//  Provides link to CI system
	EnvKeyCiSystemUrl = "CI_SYSTEM_URL"

	// NameCliCiSystemHost
	//  Provides hostname of CI server
	NameCliCiSystemHost = "ci_system.host"

	// EnvKeyCiSystemHost
	//  Provides hostname of CI server
	EnvKeyCiSystemHost = "CI_SYSTEM_HOST"

	// NameCliCiSystemVersion
	//  Provides version of CI server
	NameCliCiSystemVersion = "ci_system.version"

	// EnvKeyCiSystemVersion
	//  Provides version of CI server
	EnvKeyCiSystemVersion = "CI_SYSTEM_VERSION"
)
