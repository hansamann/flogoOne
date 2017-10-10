package HelloSven

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	firstName := context.GetInput("firstName").(string)
	lastName := context.GetInput("lastName").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo activity input: [%s] [%s]", firstName, lastName)

	// Set the result as part of the context
	context.SetOutput("fullName", "Thee fullname is  "+firstName+" "+lastName)

	return true, nil
}
