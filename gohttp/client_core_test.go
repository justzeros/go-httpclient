package gohttp

import (
	"testing"
	//"net/http"
	//"encoding/xml"
)

func TestGetRequestHeaders(t *testing.T) {
	// // initialization

	// client := NewBuilder().Build()

	// //	client,err := httpClient.Get("https://api.github.com",nil)
	// commonHeaders := make(http.Header)
	// commonHeaders.Set("Content-Type","application/json")
	// commonHeaders.Set("User-Agent","cool-http-client")
	// client.Headers = commonHeaders

	// //execution
	// requestHeaders := make(http.Header)
	// requestHeaders.Set("X-Request-Id","ABC-123")
	// finalHeaders := client.getRequestHeaders(requestHeaders)

	// // valication
	// if len(finalHeaders) != 3 {
	// 	t.Error("we expect 3 headers")
	// }

	// if finalHeaders.Get("X-Request-Id") != "ABC-123" {
	// 	t.Error("Invalid request-id recieved")
	// }

	// if finalHeaders.Get("Content-Type") != "application/json" {
	// 	t.Error("Invalid content-type recieved")
	// }

	// if finalHeaders.Get("User-Agent") != "cool-http-client" {
	// 	t.Error("Invalid user-agent received")
	// }
}

func TestGetRequestBody(t *testing.T) {
// 	//initialization
// 	client := httpClient{}

// 	t.Run("NoBodyNilResponse",func (t* testing.T) {
// 		//execution
// 		body, err := client.getRequestBody("",nil)

// 		//validation
// 		if err != nil {
// 			t.Error("err == nil: no error expected when passing a nil body")
// 		}
// 		if body != nil {
// 			t.Error("body == nil: no body expected when passing a nil body")
// 		}

// 	})
// 	t.Run("BodyWithJson",func (t* testing.T) {

// 		//execution
// 		requestBody := []string{"one","two"}
// 		body, err := client.getRequestBody("application/json",requestBody)

// 		//validation
// 		if err != nil {
// 			t.Error("err == nil: no error expected when marshaling a slice as json")
// 		}
// 		if string(body) != `["one","two"]` {
// 			t.Error("body == nil: invalid json body returned")
// 		}

// 	})
// 	t.Run("BodyWithXML",func (t* testing.T) {
// 		//execution
// 		type Plant struct {
// 			XMLName xml.Name `xml:"plant"`
// 			Id      int      `xml:"id,attr"`
// 			Name    string   `xml:"name"`
// 			Origin  []string `xml:"origin"`
// 		}

// 		coffee := &Plant{Id: 1, Name: "Coffee", Origin:[]string{"Mexico"}}
// 		body, err := client.getRequestBody("application/xml",coffee)


// 		//validation
// 		if err != nil {
// 			t.Error("err == nil: no error expected when passing an XML body")
// 		}

// 		tempXML,_ := xml.Marshal(coffee)
// 		if string(body) != string(tempXML) {
// 			t.Error("body == nil: XML body should match request")
// 		}

// 	})
// 	t.Run("BodyWithJsonAsDefault",func (t* testing.T) {
// 		//execution
// 		requestBody := []string{"one","two"}
// 		body, err := client.getRequestBody("not a valid type",requestBody)

// 		//validation
// 		if err != nil {
// 			t.Error("err == nil: no error expected when marshaling the default as json")
// 		}
// 		if string(body) != `["one","two"]` {
// 			t.Error("body == nil: invalid json body returned from default ")
// 		}
// 		if err != nil {
// 			t.Error("err == nil: no error expected when marshaling the default as json")
// 		}
// 		if string(body) != `["one","two"]` {
// 			t.Error("body == nil: invalid json body returned from default ")
// 		}

// 	})

}
