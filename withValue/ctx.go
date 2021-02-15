package withValue

import (
	"fmt"
	"net/http"
	"context"
)

type WithValueContext struct {}

func foo (res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "A FOO http request")
	// Print out extra context added by fooMiddleware.
	fmt.Println("Added context content by FOO middleware: ",req.Context().Value("DummyKey"))
}

func (v WithValueContext) Demo() {
	http.HandleFunc("/",fooMidleware(foo))
	http.ListenAndServe(":1234", nil)
}

func fooMidleware(hf http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
		// Add extra info into request context
		ctx := context.WithValue(req.Context(), "DummyKey", "DummyValue")
		req = req.WithContext(ctx)
		hf.ServeHTTP(res, req)
	})
}

