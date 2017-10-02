package servicenow

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	asrt := assert.New(t)
	res := `{"error": "foo", "reason": "bar"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		asrt.True(ok, "Basic auth should be used")
		asrt.Equal("foo", u)
		asrt.Equal("bar", p)
		fmt.Fprintln(w, res)
	}))
	defer ts.Close()

	c := Client{
		Instance: ts.URL,
		Username: "foo",
		Password: "bar",
	}

	var o struct{ Records []ChangeRequest }

	// When a ServiceNow pseudo error is sent

	err := c.PerformFor("change", "getRecords", "", nil, nil, &o)

	asrt.Error(err)
	asrt.IsType(Err{}, err)
	e := err.(Err)
	asrt.Equal(e.Err, "foo")
	asrt.Equal(e.Reason, "bar")
	asrt.Contains(e.Error(), "foo: bar")

	// When an actual response is sent

	res = `{"records": []}`
	err = c.PerformFor("change", "getRecords", "", nil, nil, &o)
	asrt.NotNil(o.Records)
	asrt.Nil(err)
}
