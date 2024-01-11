package itests

import (
	clarumcore "github.com/goclarum/clarum/core"
	"github.com/goclarum/clarum/core/orchestration/command"
	clarumhttp "github.com/goclarum/clarum/http"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"testing"
	"time"
)

var appInstance = command.Command().
	Components("go", "run", "../main.go").
	Warmup(3 * time.Second).
	Build()

var apiClient = clarumhttp.Http().Client().
	Name("apiClient").
	BaseUrl("http://localhost:3001/").
	Timeout(2000 * time.Millisecond).
	Build()

func TestMain(m *testing.M) {
	clarumcore.Setup()

	if err := appInstance.Run(); err != nil {
		log.Errorf("Test suite did not start because of startup error - %s", err)
		return
	}

	result := m.Run()

	clarumcore.Finish()
	if err := appInstance.Stop(); err != nil {
		log.Errorf("Test suite ended with shutdown error  - %s", err)

	}

	os.Exit(result)
}
