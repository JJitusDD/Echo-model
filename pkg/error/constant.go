package error

import (
	"fmt"
)

// Activity logs error definition
var (
	ErrActivityStatusIsNull = fmt.Errorf("activity status is null")
)

// Workflow logs error definition
var (
	ErrWorkflowIDIsNull = fmt.Errorf("workflow_id is null")
)

var (
	ErrReconcileRecieveDetailNotExist = fmt.Errorf("reconciliation don't exist")
)
