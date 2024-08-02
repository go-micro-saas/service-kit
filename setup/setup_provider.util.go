package setuputil

import "sync"

var (
	singletonMutex           sync.Once
	singletonLauncherManager LauncherManager
)

func NewSingletonLauncherManager(configFilePath string) (LauncherManager, error) {
	var err error
	singletonMutex.Do(func() {
		singletonLauncherManager, err = NewLauncherManager(configFilePath)
		if err != nil {
			singletonMutex = sync.Once{}
		}
	})
	return singletonLauncherManager, err
}
