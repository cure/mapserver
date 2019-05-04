package camerarenderer

import (
	"io/ioutil"
	"mapserver/colormapping"
	"mapserver/db/sqlite"
	"mapserver/mapblockaccessor"
	"mapserver/testutils"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestRenderScene(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())
	testutils.CreateTestDatabase(tmpfile.Name())

	a, err := sqlite.New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	cache := mapblockaccessor.NewMapBlockAccessor(a, 500*time.Millisecond, 1000*time.Millisecond, 1000)

	m := colormapping.NewColorMapping()
	_, err = m.LoadVFSColors(false, "/colors.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := NewRenderer(cache, m)
	data, err := r.RenderScene(0, 100, 0, SOUTH_EAST, DOWN)

	if err != nil {
		t.Fatal(err)
	}

	_, currentfilename, _, _ := runtime.Caller(0)
	path := filepath.Dir(currentfilename) + "/../output"
	os.MkdirAll(path, 0700)

	err = ioutil.WriteFile(path+"/img.png", data, 0644)

	if err != nil {
		t.Fatal(err)
	}
}
