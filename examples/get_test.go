package examples

import (
	"fmt"
	"testing"
	"net/http"
	"github.com/justzeros/go-httpclient/gohttp"
	"errors"
	"strings"
	"os"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package examples")

	gohttp.StartMockServer()
	defer gohttp.StopMockServer()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	//Tell the HTTP library to mock any further requests from here

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//		Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock {
			Method:	http.MethodGet,
			Url: "https://api.github.com",
			Error: errors.New("timeout getting github endpoints"),
		})

		//Execution
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("An error was expected")
		}

		if err == nil {
			t.Error("An error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message recieved")
		}

	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock {
			Method:	http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url": 123}`,
			//			Error: `"json unmarshal error"`,
		})

		//Execution
		endpoints, err := GetEndpoints()

		//Validation
		if endpoints != nil {
			t.Error("An error was expected")
		}

		if err == nil {
			t.Error("An error was expected")
		}

		if !strings.Contains(err.Error(), "json: cannot unmarshal number into Go") {
			t.Error("invalid error message recieved")
		}
	})


	t.Run("TestNoError", func(t *testing.T) {

		// Initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock {
			Method:	http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody: `{"current_user_url":"https://api.github.com/user"}`,
		})

		//Execution
		endpoints, err := GetEndpoints()

		//Validation
		if err != nil {
			t.Error(fmt.Sprintf("no error was expected, and we got '%s'",err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected, and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
