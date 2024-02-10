package itests

import (
	clarumcore "github.com/go-clarum/clarum-core"
	"github.com/go-clarum/clarum-core/logging"
	"github.com/go-clarum/clarum-core/orchestration/command"
	clarumhttp "github.com/go-clarum/clarum-http"
	"os"
	"testing"
	"time"
)

var appInstance = command.Command().
	Components("go", "run", "../src/main.go").
	Warmup(1 * time.Second).
	Build()

var apiClient = clarumhttp.Http().Client().
	Name("apiClient").
	BaseUrl("http://localhost:3001/").
	Timeout(2000 * time.Millisecond).
	Build()

func TestMain(m *testing.M) {
	clarumcore.Setup()

	if err := appInstance.Run(); err != nil {
		logging.Errorf("Test suite did not start because of startup error - %s", err)
		return
	}

	result := m.Run()

	if err := appInstance.Stop(); err != nil {
		logging.Errorf("Test suite ended with shutdown error  - %s", err)
	}
	clarumcore.Finish()

	os.Exit(result)
}
