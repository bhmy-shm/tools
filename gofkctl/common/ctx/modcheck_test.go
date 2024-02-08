package ctx

import (
	"github.com/bhmy-shm/tools/gofkctl/common/utils"
	"go/build"
	"os"
	"path/filepath"
	"testing"

	"github.com/bhmy-shm/tools/gofkctl/common/pathx"
	"github.com/bhmy-shm/tools/gofkctl/rpc/execx"
	"github.com/stretchr/testify/assert"
)

func TestIsGoMod(t *testing.T) {
	// create mod project
	dft := build.Default
	gp := dft.GOPATH
	if len(gp) == 0 {
		return
	}
	projectName := utils.Rand()
	dir := filepath.Join(gp, "src", projectName)
	err := pathx.CreateDirectoryIfNotExist(dir)
	if err != nil {
		return
	}

	_, err = execx.Run("go mod init "+projectName, dir)
	assert.Nil(t, err)
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	isGoMod, err := IsGoMod(dir)
	assert.Nil(t, err)
	assert.Equal(t, true, isGoMod)
}

func TestIsGoModNot(t *testing.T) {
	dft := build.Default
	gp := dft.GOPATH
	if len(gp) == 0 {
		return
	}
	projectName := utils.Rand()
	dir := filepath.Join(gp, "src", projectName)
	err := pathx.CreateDirectoryIfNotExist(dir)
	if err != nil {
		return
	}

	defer func() {
		_ = os.RemoveAll(dir)
	}()

	isGoMod, err := IsGoMod(dir)
	assert.Nil(t, err)
	assert.Equal(t, false, isGoMod)
}

func TestIsGoModWorkDirIsNil(t *testing.T) {
	_, err := IsGoMod("")
	assert.Equal(t, err.Error(), func() string {
		return "the work directory is not found"
	}())
}
