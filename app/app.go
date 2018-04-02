package app

import (
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/resources"
	"butler{ .Vars.repoPath }/butler{ toSnakeCase .Project.Name }/interfaces"
	"github.com/THE108/go-skeleton/config"
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
// Two options for config initialisation
// 1. Getter style
func Run(cfg interfaces.Config, logger interfaces.Logger, mon interfaces.Monitoring, res *resources.Resources) error {
	// here user's code goes...
	return a.doStaff()
}

// 2. Struct style
// Parser interface parses toml file into struct like:
// type AppConfig struct {
//	 App struct {
//		 Field string
//	 }
// }
func Run2(cfg *config.CommonConfig, parser interfaces.ConfigParser, logger interfaces.Logger, mon interfaces.Monitoring, res *resources.Resources) error {
	// here user's code goes...
	return a.doStaff()
}

func Shutdown() {
	a.stop()
}
