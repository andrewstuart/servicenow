package servicenow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

// Client helps users interact with a ServiceNow instance.
type Client struct {
	Username, Password, Instance string
}

func sys(param string) string {
	return fmt.Sprintf("sysparm_%s", param)
}

// PerformFor creates and executes an authenticated HTTP request to an
// instance, for the given table, action and optional id, with the passed
// options, and unmarhals the JSON into the passed output interface pointer,
// returning an error.
func (c Client) PerformFor(table, action, id string, opts url.Values, out interface{}) error {
	inst := c.Instance

	if !regexp.MustCompile("^https?://").MatchString(inst) {
		inst = "https://" + inst
	}

	u := fmt.Sprintf("%s/%s.do", inst, table)

	vals := url.Values{
		"JSONv2":      {""},
		sys("action"): {action},
	}

	if id != "" {
		vals.Set(sys("sys_id"), id)
		vals.Set("displayvalue", "true")
	}

	if opts != nil {
		vals.Set(sys("query"), opts.Encode())
	}

	req, err := http.NewRequest(http.MethodGet, u+"?"+vals.Encode(), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(out)
}

// GetFor performs a servicenow get to the specified table, with options, and
// unmarhals JSON into the output parameter.
func (c Client) GetFor(table string, id string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "get", id, opts, out)
}

// GetRecordsFor performs a servicenow getRecords to the specified table, with
// options, and unmarhals JSON into the output parameter.
func (c Client) GetRecordsFor(table string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "getRecords", "", opts, out)
}

// GetRecords performs a servicenow getRecords to the specified table, with
// options, and returns a map for each item
func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error) {
	var out struct {
		Records []map[string]interface{}
	}
	err := c.GetRecordsFor(table, opts, &out)
	return out.Records, err
}
