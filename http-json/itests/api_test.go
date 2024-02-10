package itests

import (
	"github.com/go-clarum/clarum-http/constants"
	"github.com/go-clarum/clarum-http/message"
	"net/http"
	"testing"
)

func TestJsonOKValidation(t *testing.T) {
	apiClient.In(t).Send().
		Message(message.Get())

	apiClient.In(t).Receive().
		Json().
		Message(message.Response(http.StatusOK).
			ContentType(constants.ContentTypeJsonHeader).
			Payload("{" +
				"\"name\": \"Bruce Wayne\"," +
				"\"aliases\": [" +
				"\"Batman\"," +
				"\"The Dark Knight\"" +
				"]," +
				" \"age\": 38," +
				" \"height\": 1.879," +
				"\"location\": {" +
				"\"street\": \"Mountain Drive\"," +
				"\"number\": 1007," +
				"\"hidden\": false" +
				"}" +
				"}"))
}
