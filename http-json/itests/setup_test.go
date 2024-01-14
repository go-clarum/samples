package itests

import (
	"fmt"
	clarumcore "github.com/goclarum/clarum/core"
	"github.com/goclarum/clarum/core/orchestration/command"
	clarumhttp "github.com/goclarum/clarum/http"
	"log/slog"
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
		slog.Error(fmt.Sprintf("Test suite did not start because of startup error - %s", err))
		return
	}

	result := m.Run()

	if err := appInstance.Stop(); err != nil {
		slog.Error(fmt.Sprintf("Test suite ended with shutdown error  - %s", err))
	}
	clarumcore.Finish()

	os.Exit(result)
}
