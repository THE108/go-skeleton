package app

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/resources"
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/interfaces"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) doStaff() error {
	return nil
}

func (a *Application) stop() {

}

var a = NewApplication()

// Hook called by skeleton
// Signature of that func is standardised
func Run(cfg interfaces.Config, logger interfaces.Logger, mon interfaces.Monitoring, res *resources.Resources) error {
	// here user's code goes...
	return a.doStaff()
}

func Shutdown() {
	a.stop()
}
