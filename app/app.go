package app

import (
	"butler{ .Vars.repoPath }/butler{ .Project.Name }/config"
)

// Sample application
type Application struct {
	cfg *config.Config
	quit chan struct{}
}

func NewApplication(cfg *config.Config) *Application {
	return &Application{
		cfg: cfg,
		quit: make(chan struct{}),
	}
}

// Sample application will wait until the `Shutdown` method will be called
func (a *Application) Run() error {
	<-a.quit
	return nil
}

func (a *Application) Shutdown() {
	close(a.quit)
}
