package swarm

import (
	"testing"

	appModel "github.com/jysperm/deploying/lib/models/app"
	versionModel "github.com/jysperm/deploying/lib/models/version"
	. "github.com/jysperm/deploying/lib/testing"
)

var seedApp appModel.Application
var imageVersion *versionModel.Version
var shasum string

func init() {
	var err error
	seedApp = SeedApp("https://github.com/mason96112569/docker-test.git")
	imageVersion, err = versionModel.CreateVersion(&seedApp)
	seedApp.Version = imageVersion.Tag
	if err != nil {
		panic(err)
	}
}

func TestUpdateService(t *testing.T) {
	if err := UpdateService(seedApp); err != nil {
		t.Error(err)
	}
}

func TestRemoveService(t *testing.T) {
	if err := RemoveService(seedApp); err != nil {
		t.Error(err)
	}
}
