package servicenow

import "net/url"

const TableCMDB = "cmdb_ci"

func (c Client) GetCMDBItems(query url.Values) ([]CMDBItem, error) {
	var res struct {
		Records []CMDBItem
	}
	err := c.GetRecordsFor(TableCMDB, query, &res)
	return res.Records, err
}

type CMDBItem struct {
	Status              string `json:"__status"`
	Asset               string `json:"asset"`
	AssetTag            string `json:"asset_tag"`
	Assigned            string `json:"assigned"`
	AssignedTo          string `json:"assigned_to"`
	AssignmentGroup     string `json:"assignment_group"`
	Attributes          string `json:"attributes"`
	CanPrint            string `json:"can_print"`
	Category            string `json:"category"`
	ChangeControl       string `json:"change_control"`
	CheckedIn           string `json:"checked_in"`
	CheckedOut          string `json:"checked_out"`
	Comments            string `json:"comments"`
	Company             string `json:"company"`
	CorrelationID       string `json:"correlation_id"`
	Cost                string `json:"cost"`
	CostCc              string `json:"cost_cc"`
	CostCenter          string `json:"cost_center"`
	DeliveryDate        string `json:"delivery_date"`
	Department          string `json:"department"`
	DiscoverySource     string `json:"discovery_source"`
	DNSDomain           string `json:"dns_domain"`
	Due                 string `json:"due"`
	DueIn               string `json:"due_in"`
	FaultCount          string `json:"fault_count"`
	FirstDiscovered     string `json:"first_discovered"`
	Fqdn                string `json:"fqdn"`
	GlAccount           string `json:"gl_account"`
	InstallDate         string `json:"install_date"`
	InstallStatus       string `json:"install_status"`
	InvoiceNumber       string `json:"invoice_number"`
	IPAddress           string `json:"ip_address"`
	Justification       string `json:"justification"`
	LastDiscovered      string `json:"last_discovered"`
	LeaseID             string `json:"lease_id"`
	Location            string `json:"location"`
	MacAddress          string `json:"mac_address"`
	MaintenanceSchedule string `json:"maintenance_schedule"`
	ManagedBy           string `json:"managed_by"`
	Manufacturer        string `json:"manufacturer"`
	ModelID             string `json:"model_id"`
	ModelNumber         string `json:"model_number"`
	Monitor             string `json:"monitor"`
	Name                string `json:"name"`
	OperationalStatus   string `json:"operational_status"`
	OrderDate           string `json:"order_date"`
	OwnedBy             string `json:"owned_by"`
	PoNumber            string `json:"po_number"`
	PurchaseDate        string `json:"purchase_date"`
	Schedule            string `json:"schedule"`
	SerialNumber        string `json:"serial_number"`
	ShortDescription    string `json:"short_description"`
	SkipSync            string `json:"skip_sync"`
	StartDate           string `json:"start_date"`
	Subcategory         string `json:"subcategory"`
	SupportGroup        string `json:"support_group"`
	SupportedBy         string `json:"supported_by"`
	SysClassName        string `json:"sys_class_name"`
	SysClassPath        string `json:"sys_class_path"`
	SysCreatedBy        string `json:"sys_created_by"`
	SysCreatedOn        string `json:"sys_created_on"`
	SysDomain           string `json:"sys_domain"`
	SysDomainPath       string `json:"sys_domain_path"`
	SysID               string `json:"sys_id"`
	SysModCount         string `json:"sys_mod_count"`
	SysTags             string `json:"sys_tags"`
	SysUpdatedBy        string `json:"sys_updated_by"`
	SysUpdatedOn        string `json:"sys_updated_on"`
	Unverified          string `json:"unverified"`
	Vendor              string `json:"vendor"`
	WarrantyExpiration  string `json:"warranty_expiration"`
}
