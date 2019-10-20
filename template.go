package domain

import (
	"encoding/xml"
)

// Application representation
type Application struct {
	Name string `xml:"name" json:"name" yaml:"name"`
}

// Condition representation
type Condition struct {
	Macro     string `xml:"macro,allowempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value     string `xml:"value,allowempty" json:"value,omitempty" yaml:"value,omitempty"`
	Operator  string `xml:"operator,allowempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	Formulaid string `xml:"formulaid,allowempty" json:"formulaid,omitempty" yaml:"formulaid,omitempty"`
}

// Dependency representation
type Dependency struct {
	Name               string `xml:"name,allowempty"`
	Expression         string `xml:"expression,allowempty"`
	RecoveryExpression string `xml:"recovery_expression,allowempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
}

// DiscoveryRule representation
type DiscoveryRule struct {
	Name                 string           `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                 string           `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity        string           `xml:"snmp_community,allowempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid              string           `xml:"snmp_oid,allowempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                  string           `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                string           `xml:"delay,allowempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	Status               string           `xml:"status,allowempty" json:"status,omitempty" yaml:"status,omitempty"`
	AllowedHosts         string           `xml:"allowed_hosts,allowempty" json:"allowed_hosts,omitempty" yaml:"allowed_hosts,omitempty"`
	Snmpv3Contextname    string           `xml:"snmpv3_contextname,allowempty" json:"snmpv3_contextname,omitempty" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname   string           `xml:"snmpv3_securityname,allowempty" json:"snmpv3_securityname,omitempty" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel  int              `xml:"snmpv3_securitylevel,allowempty" json:"snmpv3_securitylevel,omitempty" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol   int              `xml:"snmpv3_authprotocol,allowempty" json:"snmpv3_authprotocol,omitempty" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase string           `xml:"snmpv3_authpassphrase,allowempty" json:"snmpv3_authpassphrase,omitempty" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol   int              `xml:"snmpv3_privprotocol,allowempty" json:"snmpv3_privprotocol,omitempty" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase string           `xml:"snmpv3_privpassphrase,allowempty" json:"snmpv3_privpassphrase,omitempty" yaml:"snmpv3_privpassphrase,omitempty"`
	Params               string           `xml:"params,allowempty" json:"params,omitempty" yaml:"params,omitempty"`
	IpmiSensor           string           `xml:"ipmi_sensor,allowempty" json:"ipmi_sensor,omitempty" yaml:"ipmi_sensor,omitempty"`
	Authtype             string           `xml:"authtype,allowempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username             string           `xml:"username,allowempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password             string           `xml:"password,allowempty" json:"password,omitempty" yaml:"password,omitempty"`
	Publickey            string           `xml:"publickey,allowempty" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	Privatekey           string           `xml:"privatekey,allowempty" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Port                 string           `xml:"port,allowempty" json:"port,omitempty" yaml:"port,omitempty"`
	Filter               Filter           `xml:"filter,allowempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Lifetime             string           `xml:"lifetime,allowempty" json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
	Description          string           `xml:"description,allowempty" json:"decription,omitempty" yaml:"description,omitempty"`
	ItemPrototypes       []ItemPrototype  `xml:"item_prototypes>item_prototype,allowempty" json:"item_prototypes,omitempty" yaml:"item_prototypes,omitempty"`
	TriggerPrototypes    []Trigger        `xml:"trigger_prototypes>trigger_prototype,allowempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	GraphPrototypes      []GraphPrototype `xml:"graph_prototypes>graph_prototype,allowempty" json:"graph_prototypes,omitempty" yaml:"graph_prototypes,omitempty"`
	MasterItem           MasterItem       `xml:"master_item,allowempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	LldMacroPaths        []LldMacroPath   `xml:"lld_macro_paths>lld_macro_path,allowempty" json:"lld_macro_paths,omitempty" yaml:"lld_macro_paths,omitempty"`
	Preprocessing        []Step           `xml:"preprocessing>step,allowempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
}

// Filter representation
type Filter struct {
	Evaltype   string      `xml:"evaltype,allowempty" json:"evaltype,omitempty" yaml:"evaltype,omitempty"`
	Formula    string      `xml:"formula,allowempty" json:"formula,omitempty" yaml:"formula,omitempty"`
	Conditions []Condition `xml:"conditions>condition,allowempty" json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// GItem representation
type GItem struct {
	Host string `xml:"host,allowempty" json:"host,omitempty" yaml:"host,omitempty"`
	Key  string `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Graph representation
type Graph struct {
	Name       string      `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Width      int         `xml:"width,allowempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height     int         `xml:"height,allowempty" json:"height,omitempty" yaml:"height,omitempty"`
	Type       string      `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	YminType1  string      `xml:"ymin_type_1,allowempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
	YmaxType1  string      `xml:"ymax_type_1,allowempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	Show3d     string      `xml:"show_3d,allowempty" json:"show_3d,omitempty" yaml:"show_3d,omitempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,allowempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// GraphItem representation
type GraphItem struct {
	Sortorder int    `xml:"sortorder,allowempty" json:"sortorder,omitempty" yaml:"sortorder,omitempty"`
	Drawtype  string `xml:"drawtype,allowempty" json:"drawtype,omitempty" yaml:"drawtype,omitempty"`
	Color     string `xml:"color,allowempty" json:"color,omitempty" yaml:"color,omitempty"`
	Yaxisside string `xml:"yaxisside,allowempty" json:"yaxisside,omitempty" yaml:"yaxisside,omitempty"`
	CalcFnc   string `xml:"calc_fnc,allowempty" json:"calc_fnc,omitempty" yaml:"calc_fnc,omitempty"`
	Type      string `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	Item      GItem  `xml:"item,allowempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GraphPrototype representation
type GraphPrototype struct {
	Name           string      `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Width          int         `xml:"width,allowempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height         int         `xml:"height,allowempty" json:"height,omitempty" yaml:"height,omitempty"`
	Yaxismin       float32     `xml:"yaxismin,allowempty" json:"yaxismin,omitempty" yaml:"yaxismin,omitempty"`
	Yaxismax       float32     `xml:"yaxismax,allowempty" json:"yaxismax,omitempty" yaml:"yaxismax,omitempty"`
	ShowWorkPeriod int         `xml:"show_work_period,allowempty" json:"show_work_period,omitempty" yaml:"show_work_period,omitempty"`
	ShowTriggers   int         `xml:"show_triggers,allowempty" json:"show_triggers,omitempty" yaml:"show_triggers,omitempty"`
	Type           string      `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	ShowLegend     int         `xml:"show_legend,allowempty" json:"show_legend,omitempty" yaml:"show_legend,omitempty"`
	Show3d         int         `xml:"show_3d,allowempty" json:"show_3d,omitempty" yaml:"show_3d,omitempty"`
	PercentLeft    float32     `xml:"percent_left,allowempty" json:"percent_left,omitempty" yaml:"percent_left,omitempty"`
	PercentRight   float32     `xml:"percent_right,allowempty" json:"percent_right,omitempty" yaml:"percent_right,omitempty"`
	YminType1      string      `xml:"ymin_type_1,allowempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
	YmaxType1      string      `xml:"ymax_type_1,allowempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	GraphItems     []GraphItem `xml:"graph_items>graph_item,allowempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// Group representation
type Group struct {
	Name string `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Item representation
type Item struct {
	Name                 string        `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                 string        `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity        string        `xml:"snmp_community,allowempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid              string        `xml:"snmp_oid,allowempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                  string        `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                string        `xml:"delay,allowempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History              string        `xml:"history,allowempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends               string        `xml:"trends,allowempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	Status               string        `xml:"status,allowempty" json:"status,omitempty" yaml:"status,omitempty"`
	ValueType            string        `xml:"value_type,allowempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	AllowedHosts         string        `xml:"allowed_hosts,allowempty" json:"allowed_hosts,omitempty" yaml:"allowed_hosts,omitempty"`
	Units                string        `xml:"units,allowempty" json:"units,omitempty" yaml:"units,omitempty"`
	Snmpv3Contextname    string        `xml:"snmpv3_contextname,allowempty" json:"snmpv3_contextname,omitempty" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname   string        `xml:"snmpv3_securityname,allowempty" json:"snmpv3_securityname,omitempty" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel  int           `xml:"snmpv3_securitylevel,allowempty" json:"snmpv3_securitylevel,omitempty" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol   int           `xml:"snmpv3_authprotocol,allowempty" json:"snmpv3_authprotocol,omitempty" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase string        `xml:"snmpv3_authpassphrase,allowempty" json:"snmpv3_authpassphrase,omitempty" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol   int           `xml:"snmpv3_privprotocol,allowempty" json:"snmpv3_privprotocol,omitempty" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase string        `xml:"snmpv3_privpassphrase,allowempty" json:"snmpv3_privpassphrase,omitempty" yaml:"snmpv3_privpassphrase,omitempty"`
	Params               string        `xml:"params,allowempty" json:"params,omitempty" yaml:"params,omitempty"`
	IpmiSensor           string        `xml:"ipmi_sensor,allowempty" json:"ipmi_sensor,omitempty" yaml:"ipmi_sensor,omitempty"`
	Authtype             string        `xml:"authtype,allowempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username             string        `xml:"username,allowempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password             string        `xml:"password,allowempty" json:"password,omitempty" yaml:"password,omitempty"`
	Publickey            string        `xml:"publickey,allowempty" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	Privatekey           string        `xml:"privatekey,allowempty" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Port                 string        `xml:"port,allowempty" json:"port,omitempty" yaml:"port,omitempty"`
	Description          string        `xml:"description,allowempty" json:"decription,omitempty" yaml:"description,omitempty"`
	InventoryLink        string        `xml:"inventory_link,allowempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications         []Application `xml:"applications>application,allowempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	ValueMap             ValueMap      `xml:"valuemap,allowempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	Logtimefmt           string        `xml:"logtimefmt,allowempty" json:"logtimefmt,omitempty" yaml:"logtimefmt,omitempty"`
	Preprocessing        []Step        `xml:"preprocessing>step,allowempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	JmxEndpoint          string        `xml:"jmx_endpoint,allowempty" json:"jmx_endpoint,omitempty" yaml:"jmx_endpoint,omitempty"`
	Timeout              string        `xml:"timeout,allowempty" json:"timeout,omitempty" yaml:"timeout,omitempty"`
	URL                  string        `xml:"url,allowempty" json:"url,omitempty" yaml:"url,omitempty"`
	StatusCodes          int           `xml:"status_codes,allowempty" json:"status_codes,omitempty" yaml:"status_codes,omitempty"`
	FollowRedirects      int           `xml:"follow_redirects,allowempty" json:"follow_redirects,omitempty" yaml:"follow_redirects,omitempty"`
	PostType             int           `xml:"post_type,allowempty" json:"post_type,omitempty" yaml:"post_type,omitempty"`
	RetrieveMode         string        `xml:"retrieve_mode,allowempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
	RequestMethod        int           `xml:"request_method,allowempty" json:"request_method,omitempty" yaml:"request_method,omitempty"`
	OutputFormat         int           `xml:"output_format,allowempty" json:"output_format,omitempty" yaml:"output_format,omitempty"`
	AllowTraps           int           `xml:"allow_traps,allowempty" json:"allow_traps,omitempty" yaml:"allow_traps,omitempty"`
	SslCertFile          string        `xml:"ssl_cert_file,allowempty" json:"ssl_cert_file,omitempty" yaml:"ssl_cert_file,omitempty"`
	SslKeyFile           string        `xml:"ssl_key_file,allowempty" json:"ssl_key_file,omitempty" yaml:"ssl_key_file,omitempty"`
	SslKeyPassword       string        `xml:"ssl_key_password,allowempty" json:"ssl_key_password,omitempty" yaml:"ssl_key_password,omitempty"`
	VerifyPeer           int           `xml:"verify_peer,allowempty" json:"verify_peer,omitempty" yaml:"verify_peer,omitempty"`
	VerifyHost           int           `xml:"verify_host,allowempty" json:"verify_host,omitempty" yaml:"verify_host,omitempty"`
	MasterItem           MasterItem    `xml:"master_item,allowempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	Triggers             []Trigger     `xml:"triggers>trigger,allowempty" json:"triggers,omitempty" yaml:"triggers,omitempty"`
}

// ItemPrototype representation
type ItemPrototype struct {
	Name                  string        `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                  string        `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity         string        `xml:"snmp_community,allowempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid               string        `xml:"snmp_oid,allowempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                   string        `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                 string        `xml:"delay,allowempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History               string        `xml:"history,allowempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends                string        `xml:"trends,allowempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	Status                string        `xml:"status,allowempty" json:"status,omitempty" yaml:"status,omitempty"`
	ValueType             string        `xml:"value_type,allowempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	AllowedHosts          string        `xml:"allowed_hosts,allowempty" json:"allowed_hosts,omitempty" yaml:"allowed_hosts,omitempty"`
	Units                 string        `xml:"units,allowempty" json:"units,omitempty" yaml:"units,omitempty"`
	Snmpv3Contextname     string        `xml:"snmpv3_contextname,allowempty" json:"snmpv3_contextname,omitempty" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname    string        `xml:"snmpv3_securityname,allowempty" json:"snmpv3_securityname,omitempty" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel   int           `xml:"snmpv3_securitylevel,allowempty" json:"snmpv3_securitylevel,omitempty" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol    int           `xml:"snmpv3_authprotocol,allowempty" json:"snmpv3_authprotocol,omitempty" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase  string        `xml:"snmpv3_authpassphrase,allowempty" json:"snmpv3_authpassphrase,omitempty" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol    int           `xml:"snmpv3_privprotocol,allowempty" json:"snmpv3_privprotocol,omitempty" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase  string        `xml:"snmpv3_privpassphrase,allowempty" json:"snmpv3_privpassphrase,omitempty" yaml:"snmpv3_privpassphrase,omitempty"`
	Params                string        `xml:"params,allowempty" json:"params,omitempty" yaml:"params,omitempty"`
	IpmiSensor            string        `xml:"ipmi_sensor,allowempty" json:"ipmi_sensor,omitempty" yaml:"ipmi_sensor,omitempty"`
	Authtype              string        `xml:"authtype,allowempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username              string        `xml:"username,allowempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password              string        `xml:"password,allowempty" json:"password,omitempty" yaml:"password,omitempty"`
	Publickey             string        `xml:"publickey,allowempty" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	Privatekey            string        `xml:"privatekey,allowempty" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Port                  string        `xml:"port,allowempty" json:"port,omitempty" yaml:"port,omitempty"`
	Description           string        `xml:"description,allowempty" json:"decription,omitempty" yaml:"description,omitempty"`
	InventoryLink         string        `xml:"inventory_link,allowempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications          []Application `xml:"applications>application,allowempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	ValueMap              ValueMap      `xml:"valuemap,allowempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	Logtimefmt            string        `xml:"logtimefmt,allowempty" json:"logtimefmt,omitempty" yaml:"logtimefmt,omitempty"`
	Preprocessing         []Step        `xml:"preprocessing>step,allowempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	JmxEndpoint           string        `xml:"jmx_endpoint,allowempty" json:"jmx_endpoint,omitempty" yaml:"jmx_endpoint,omitempty"`
	Timeout               string        `xml:"timeout,allowempty" json:"timeout,omitempty" yaml:"timeout,omitempty"`
	TriggerPrototypes     []Trigger     `xml:"trigger_prototypes>trigger_prototype,allowempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	URL                   string        `xml:"url,allowempty" json:"url,omitempty" yaml:"url,omitempty"`
	StatusCodes           int           `xml:"status_codes,allowempty" json:"status_codes,omitempty" yaml:"status_codes,omitempty"`
	FollowRedirects       int           `xml:"follow_redirects,allowempty" json:"follow_redirects,omitempty" yaml:"follow_redirects,omitempty"`
	PostType              int           `xml:"post_type,allowempty" json:"post_type,omitempty" yaml:"post_type,omitempty"`
	RetrieveMode          string        `xml:"retrieve_mode,allowempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
	RequestMethod         int           `xml:"request_method,allowempty" json:"request_method,omitempty" yaml:"request_method,omitempty"`
	OutputFormat          int           `xml:"output_format,allowempty" json:"output_format,omitempty" yaml:"output_format,omitempty"`
	AllowTraps            int           `xml:"allow_traps,allowempty" json:"allow_traps,omitempty" yaml:"allow_traps,omitempty"`
	SslCertFile           string        `xml:"ssl_cert_file,allowempty" json:"ssl_cert_file,omitempty" yaml:"ssl_cert_file,omitempty"`
	SslKeyFile            string        `xml:"ssl_key_file,allowempty" json:"ssl_key_file,omitempty" yaml:"ssl_key_file,omitempty"`
	SslKeyPassword        string        `xml:"ssl_key_password,allowempty" json:"ssl_key_password,omitempty" yaml:"ssl_key_password,omitempty"`
	VerifyPeer            int           `xml:"verify_peer,allowempty" json:"verify_peer,omitempty" yaml:"verify_peer,omitempty"`
	VerifyHost            int           `xml:"verify_host,allowempty" json:"verify_host,omitempty" yaml:"verify_host,omitempty"`
	ApplicationPrototypes []Application `xml:"application_prototypes>application_prototype,allowempty" json:"application_prototypes,omitempty" yaml:"application_prototypes,omitempty"`
	MasterItem            MasterItem    `xml:"master_item,allowempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
}

// LldMacroPath representation
type LldMacroPath struct {
	LldMacro string `xml:"lld_macro,allowempty" json:"lld_macro,omitempty" yaml:"lld_macro,omitempty"`
	Path     string `xml:"path,allowempty" json:"path,omitempty" yaml:"path,omitempty"`
}

// Macro representation
type Macro struct {
	Macro       string `xml:"macro,allowempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value       string `xml:"value,allowempty" json:"value,omitempty" yaml:"value,omitempty"`
	Description string `xml:"description,allowempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// Mapping representation
type Mapping struct {
	Value    string `xml:"value,allowempty" json:"value,omitempty" yaml:"value,omitempty"`
	Newvalue string `xml:"newvalue,allowempty" json:"newvalue,omitempty" yaml:"newvalue,omitempty"`
}

// MasterItem representation
type MasterItem struct {
	Key string `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Resource representation
type Resource struct {
	Name string `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Key  string `xml:"key,allowempty" json:"key,omitempty" yaml:"key,omitempty"`
	Host string `xml:"host,allowempty" json:"host,omitempty" yaml:"host,omitempty"`
}

// Screen representation
type Screen struct {
	Name        string       `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Hsize       int          `xml:"hsize,allowempty" json:"hsize,omitempty" yaml:"hsize,omitempty"`
	Vsize       int          `xml:"vsize,allowempty" json:"vsize,omitempty" yaml:"vsize,omitempty"`
	ScreenItems []ScreenItem `xml:"screen_items>screen_item,allowempty" json:"screen_items,omitempty" yaml:"screen_items,omitempty"`
}

// ScreenItem representation
type ScreenItem struct {
	Resourcetype int      `xml:"resourcetype" json:"resourcetype,omitempty" yaml:"resourcetype,omitempty"`
	Style        int      `xml:"style" json:"style,omitempty" yaml:"style,omitempty"`
	Resource     Resource `xml:"resource,allowempty" json:"resource,omitempty" yaml:"resource,omitempty"`
	Width        int      `xml:"width,allowempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height       int      `xml:"height,allowempty" json:"height,omitempty" yaml:"height,omitempty"`
	X            int      `xml:"x" json:"x" yaml:"x"`
	Y            int      `xml:"y" json:"y" yaml:"y"`
	Colspan      int      `xml:"colspan,allowempty" json:"colspan,omitempty" yaml:"colspan,omitempty"`
	Rowspan      int      `xml:"rowspan,allowempty" json:"rowspan,omitempty" yaml:"rowspan,omitempty"`
	Elements     int      `xml:"elements,allowempty" json:"elements,omitempty" yaml:"elements,omitempty"`
	Valign       int      `xml:"valign" json:"valign,omitempty" yaml:"valign,omitempty"`
	Halign       int      `xml:"halign" json:"halign,omitempty" yaml:"halign,omitempty"`
	Dynamic      int      `xml:"dynamic" json:"dynamic,omitempty" yaml:"dynamic,omitempty"`
	SortTriggers int      `xml:"sort_triggers" json:"sort_triggers,omitempty" yaml:"sort_triggers,omitempty"`
	URL          string   `xml:"url" json:"url,allowempty" yaml:"url,omitempty"`
	Application  string   `xml:"application" json:"application,omitempty" yaml:"application,omitempty"`
	MaxColumns   int      `xml:"max_columns,allowempty" json:"max_columns,omitempty" yaml:"max_columns,omitempty"`
}

// Step representation
type Step struct {
	Type               string `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	Params             string `xml:"params" json:"params,omitempty" yaml:"params,omitempty"`
	ErrorHandler       string `xml:"error_handler,allowempty" json:"error_handler,omitempty" yaml:"error_handler,omitempty"`
	ErrorHandlerParams string `xml:"error_handler_params,allowempty" json:"error_handler_params,omitempty" yaml:"error_handler_params,omitempty"`
}

// Template representation
type Template struct {
	Template       string          `xml:"template,allowempty" json:"template,omitempty" yaml:"template,omitempty"`
	Name           string          `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description    string          `xml:"description,allowempty" json:"description,omitempty" yaml:"description,omitempty"`
	Templates      []Template      `xml:"templates>template,allowempty" json:"templates,omitempty" yaml:"templates,omitempty"`
	Groups         []Group         `xml:"groups>group,allowempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Applications   []Application   `xml:"applications>application,allowempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	Items          []Item          `xml:"items>item,allowempty" json:"items,omitempty" yaml:"items,omitempty"`
	DiscoveryRules []DiscoveryRule `xml:"discovery_rules>discovery_rule,allowempty" json:"discovery_rules,omitempty" yaml:"discovery_rules,omitempty"`
	Macros         []Macro         `xml:"macros>macro,allowempty" json:"macros,omitempty" yaml:"macros,omitempty"`
	Screens        []Screen        `xml:"screens>screen,allowempty" json:"screens,omitempty" yaml:"screens,omitempty"`
}

// Trigger representation
type Trigger struct {
	Expression         string       `xml:"expression,allowempty" json:"expression,omitempty" yaml:"expression,omitempty"`
	RecoveryMode       string       `xml:"recovery_mode,allowempty" json:"recovery_mode,omitempty" yaml:"recovery_mode,omitempty"`
	RecoveryExpression string       `xml:"recovery_expression,allowempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
	Name               string       `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	CorrelationMode    int          `xml:"correlation_mode,allowempty" json:"correlation_mode,omitempty" yaml:"correlation_mode,omitempty"`
	CorrelationTag     string       `xml:"correlation_tag,allowempty" json:"correlation_tag,omitempty" yaml:"correlation_tag,omitempty"`
	URL                string       `xml:"url,allowempty" json:"url,omitempty" yaml:"url,omitempty"`
	Status             int          `xml:"status,allowempty" json:"status,omitempty" yaml:"status,omitempty"`
	Priority           string       `xml:"priority,allowempty" json:"priority,omitempty" yaml:"priority,omitempty"`
	Description        string       `xml:"description,allowempty" json:"description,omitempty" yaml:"description,omitempty"`
	Type               int          `xml:"type,allowempty" json:"type,omitempty" yaml:"type,omitempty"`
	ManualClose        string       `xml:"manual_close,allowempty" json:"manual_close,omitempty" yaml:"manual_close,omitempty"`
	Dependencies       []Dependency `xml:"dependencies>dependency,allowempty" json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}

// ValueMap representation
type ValueMap struct {
	Name     string    `xml:"name,allowempty" json:"name,omitempty" yaml:"name,omitempty"`
	Mappings []Mapping `xml:"mappings>mapping,allowempty" json:"mappings,omitempty" yaml:"mappings,omitempty"`
}

// ZabbixExport representation
type ZabbixExport struct {
	XMLName   xml.Name   `xml:"zabbix_export" json:"-" yaml:"-"`
	Version   string     `xml:"version,allowempty" json:"version,omitempty" yaml:"version,omitempty"`
	Date      string     `xml:"date,allowempty" json:"date,omitempty" yaml:"date,omitempty"`
	Groups    []Group    `xml:"groups>group,allowempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Templates []Template `xml:"templates>template,allowempty" json:"templates,omitempty" yaml:"templates,omitempty"`
	Triggers  []Trigger  `xml:"triggers>trigger,allowempty" json:"triggers,omitempty" yaml:"triggers,omitempty"`
	Graphs    []Graph    `xml:"graphs>graph,allowempty" json:"graphs,omitempty" yaml:"graphs,omitempty"`
	ValueMaps []ValueMap `xml:"value_maps>value_map,allowempty" json:"value_maps,omitempty" yaml:"value_maps,omitempty"`
}
