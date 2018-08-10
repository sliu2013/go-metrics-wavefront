package wavefront

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	"io"
	"net/http"
)

func TestServerMethodOfdirectReporterStructType (t *testing.T) {
	expected := "https://domain.wavefront.com"
	reporter := directReporter{serverURL: expected, token: "12345678"}
	if reporter.Server() != expected {
		t.Error("come on")
	}
}


//func TestReportMethodOfdirectReporterStructType (t *testing.T) {
//	expected := "https://domain.wavefront.com"
//	reporter := directReporter{serverURL: expected, token: "12345678"}
//	reporter.Report("x", "y", mock_new_request)
//	if reporter.Server() != expected {
//		t.Error("come on")
//	}
//}

func TestSuccessResponse(t *testing.T){

}


func TestErrorResponse(t *testing.T){

}


func mock_new_request(method string, url string, body io.Reader) (*http.Request, error){
	return nil, nil
}