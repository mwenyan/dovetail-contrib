package history

import (
	impl "github.com/TIBCOSoftware/dovetail-contrib/smartcontract-go/activity/history"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return impl.NewActivity(metadata)
}