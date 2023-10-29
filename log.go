package main

import (
	"fmt"
	"net/http"
	"bytes"
)

func printToConsole(w http.ResponseWriter, r *http.Request) {
	// Create a buffer to capture the response
	var buffer bytes.Buffer

	// Wrap the original response writer with a custom one that captures the response
	customResponseWriter := NewResponseCaptureWriter(w, &buffer)

	// Write some data to the custom response writer
	fmt.Fprint(customResponseWriter, "Hello, World!")

	// You can access the captured response in the buffer
	capturedResponse := buffer.String()

	// Now you can use the capturedResponse as needed
	fmt.Println("Captured Response:", capturedResponse)
}

// CustomResponseCaptureWriter is a custom response writer that captures the response
type CustomResponseCaptureWriter struct {
	originalWriter http.ResponseWriter
	buffer         *bytes.Buffer
}

func (w *CustomResponseCaptureWriter) Write(p []byte) (int, error) {
	// Write to the original response writer and the buffer
	n, err1 := w.originalWriter.Write(p)
	n, err2 := w.buffer.Write(p)
	if err1 != nil {
		return n, err1
	}
	return n, err2
}

func (w *CustomResponseCaptureWriter) Header() http.Header {
	return w.originalWriter.Header()
}

func (w *CustomResponseCaptureWriter) WriteHeader(statusCode int) {
	w.originalWriter.WriteHeader(statusCode)
}

func NewResponseCaptureWriter(w http.ResponseWriter, buffer *bytes.Buffer) *CustomResponseCaptureWriter {
	return &CustomResponseCaptureWriter{
		originalWriter: w,
		buffer:         buffer,
	}
}

func main() {
	http.HandleFunc("/", printToConsole)
	http.ListenAndServe(":8080", nil)
}