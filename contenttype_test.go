package contenttype

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// getTestHandler returns a http.HandlerFunc for testing http middleware
func getTestHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {}

	return http.HandlerFunc(fn)
}

func TestMiddleware(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "mytest/testval"

	rr := httptest.NewRecorder()
	p := Middleware(s, expected)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("Middleware failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestJSON(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "application/json"

	rr := httptest.NewRecorder()
	p := JSON(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("JSON failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestHTML(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "text/html"

	rr := httptest.NewRecorder()
	p := HTML(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("HTML failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestXML(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "application/xml"

	rr := httptest.NewRecorder()
	p := XML(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("XML failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestPDF(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "application/pdf"

	rr := httptest.NewRecorder()
	p := PDF(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("PDF failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestText(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "text/plain"

	rr := httptest.NewRecorder()
	p := Text(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("Text failed: Content-Type expected %s, got %s.", expected, value)
	}
}

func TestOctetStream(t *testing.T) {
	s := http.NewServeMux()
	s.HandleFunc("/", getTestHandler())

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "application/octet-stream"

	rr := httptest.NewRecorder()
	p := OctetStream(s)
	p.ServeHTTP(rr, req)

	value := rr.Header().Get("Content-Type")
	if value != expected {
		t.Fatalf("OctetStream failed: Content-Type expected %s, got %s.", expected, value)
	}
}
