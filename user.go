package servicenow

import "net/url"

// GetUsers returns all users that match a query.
func (c Client) GetUsers(query url.Values) ([]map[string]interface{}, error) {
	var m struct{ Records []map[string]interface{} }
	err := c.GetRecordsFor("sys_user", query, &m)
	return m.Records, err
}
