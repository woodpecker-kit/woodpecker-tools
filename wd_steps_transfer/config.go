package wd_steps_transfer

const (
	PrefixKitStepsTransfer = "WOODPECKER_KIT_STEPS_TRANSFER_"

	DefaultKitStepsFileName = ".woodpecker_kit.steps.transfer"
)

type TransferObject struct {
	Mark             string
	CiWorkflowName   string
	CiWorkflowNumber string
	Data             interface{}
}
