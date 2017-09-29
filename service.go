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

func (c Client) PerformFor(table, action string, opts url.Values, out interface{}) error {
	inst := c.Instance

	if !regexp.MustCompile("^https?://").MatchString(inst) {
		inst = "https://" + inst
	}

	u := fmt.Sprintf("%s/%s.do", inst, table)

	vals := url.Values{
		"JSONv2":         {""},
		"sysparm_action": {action},
	}

	if opts != nil {
		vals.Set("sysparm_query", opts.Encode())
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

func (c Client) GetFor(table string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "get", opts, out)
}

func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error) {
	var out struct {
		Records []map[string]interface{}
	}
	err := c.PerformFor(table, "getRecords", opts, &out)
	return out.Records, err
}
