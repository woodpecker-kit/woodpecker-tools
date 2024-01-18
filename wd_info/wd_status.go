package wd_info

const (
	// form https://github.com/woodpecker-ci/woodpecker/blob/main/server/model/const.go

	BuildStatusSuccess  = "success"  // successfully finished
	BuildStatusFailure  = "failure"  // failed to finish (exit code != 0)
	BuildStatusError    = "error"    // error with the config / while parsing / some other system problem
	BuildStatusKilled   = "killed"   // killed by user
	BuildStatusSkipped  = "skipped"  // skipped as another step failed
	BuildStatusPending  = "pending"  // pending to be executed
	BuildStatusRunning  = "running"  // currently running
	BuildStatusBlocked  = "blocked"  // waiting for approval
	BuildStatusDeclined = "declined" // blocked and declined
	BuildStatusCreated  = "created"  // created / internal use only
)
