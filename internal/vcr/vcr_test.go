package vcr

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestHTTPRequest(t *testing.T) {
	// Start recorder
	r, err := recorder.New("fixtures/iana-reserved-domains")
	if err != nil {
		t.Fatal(err)
	}

	defer r.Stop()

	client := &http.Client{
		Transport: r,
	}

	url := "https://www.iana.org/domains/reserved"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("Failed to get URL %s: %s", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	wantHeading := "<h1>IANA-managed Reserved Domains</h1>"
	bodyContent := string(body)

	if !strings.Contains(bodyContent, wantHeading) {
		t.Errorf("Heading %s not found in response", wantHeading)
	}
}
