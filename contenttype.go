package contenttype

import "net/http"

// ContentType is type for "Content-Type" header value
type ContentType string

const (
	// TextType is "text/plain" Content-Type
	TextType ContentType = "text/plain"
	// JSONType is "application/json" Content-Type
	JSONType ContentType = "application/json"
	// HTMLType is "text/html" Content-Type
	HTMLType ContentType = "text/html"
	// XMLType is "application/xml" Content-Type
	XMLType ContentType = "application/xml"
	// OctetStreamType is "application/octet-stream" Content-Type
	OctetStreamType ContentType = "application/octet-stream"
	// PDFType is "application/pdf" Content-Type
	PDFType ContentType = "application/pdf"
)

// Middleware set "Content-Type" with defined value
func Middleware(h http.Handler, contentType ContentType) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", string(contentType))

		h.ServeHTTP(w, r)
	})
}

// JSON middleware set header Content-Type: application/json
func JSON(h http.Handler) http.Handler {
	return Middleware(h, JSONType)
}

// HTML middleware set header Content-Type: text/html
func HTML(h http.Handler) http.Handler {
	return Middleware(h, HTMLType)
}

// XML middleware set header Content-Type: application/xml
func XML(h http.Handler) http.Handler {
	return Middleware(h, XMLType)
}

// PDF middleware set header Content-Type: application/pdf
func PDF(h http.Handler) http.Handler {
	return Middleware(h, PDFType)
}

// Text middleware set header Content-Type: text/plain
func Text(h http.Handler) http.Handler {
	return Middleware(h, TextType)
}

// OctetStream middleware set header Content-Type: application/octet-stream
func OctetStream(h http.Handler) http.Handler {
	return Middleware(h, OctetStreamType)
}
