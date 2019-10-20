package domain

import (
	"encoding/xml"
)

// Application representation OK
type Application struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Condition representation
type Condition struct {
	Macro     string `xml:"macro,omitempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value     string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Operator  string `xml:"operator,omitempty" json:"operator,omitempty" yaml:"operator,omitempty"`
	Formulaid string `xml:"formulaid,omitempty" json:"formulaid,omitempty" yaml:"formulaid,omitempty"`
}

// Dependency representation OK
type Dependency struct {
	Name               string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Expression         string `xml:"expression,omitempty" json:"expression,omitempty" yaml:"expression,omitempty"`
	RecoveryExpression string `xml:"recovery_expression,omitempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
}

// DiscoveryRule representation TODO
type DiscoveryRule struct {
	Name                 string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                 string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity        string `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid              string `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                  string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                string `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	Status               string `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	AllowedHosts         string `xml:"allowed_hosts,omitempty" json:"allowed_hosts,omitempty" yaml:"allowed_hosts,omitempty"`
	Snmpv3Contextname    string `xml:"snmpv3_contextname,omitempty" json:"snmpv3_contextname,omitempty" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname   string `xml:"snmpv3_securityname,omitempty" json:"snmpv3_securityname,omitempty" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel  int    `xml:"snmpv3_securitylevel,omitempty" json:"snmpv3_securitylevel,omitempty" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol   int    `xml:"snmpv3_authprotocol,omitempty" json:"snmpv3_authprotocol,omitempty" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase string `xml:"snmpv3_authpassphrase,omitempty" json:"snmpv3_authpassphrase,omitempty" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol   int    `xml:"snmpv3_privprotocol,omitempty" json:"snmpv3_privprotocol,omitempty" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase string `xml:"snmpv3_privpassphrase,omitempty" json:"snmpv3_privpassphrase,omitempty" yaml:"snmpv3_privpassphrase,omitempty"`
	Params               string `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
	IpmiSensor           string `xml:"ipmi_sensor,omitempty" json:"ipmi_sensor,omitempty" yaml:"ipmi_sensor,omitempty"`
	Authtype             string `xml:"authtype,omitempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username             string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password             string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Publickey            string `xml:"publickey,omitempty" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	Privatekey           string `xml:"privatekey,omitempty" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Port                 string `xml:"port,omitempty" json:"port,omitempty" yaml:"port,omitempty"`

	Filter            Filter           `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Lifetime          string           `xml:"lifetime,omitempty" json:"lifetime,omitempty" yaml:"lifetime,omitempty"`
	Description       string           `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	ItemPrototypes    []ItemPrototype  `xml:"item_prototypes>item_prototype,omitempty" json:"item_prototypes,omitempty" yaml:"item_prototypes,omitempty"`
	TriggerPrototypes []Trigger        `xml:"trigger_prototypes>trigger_prototype,omitempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	GraphPrototypes   []GraphPrototype `xml:"graph_prototypes>graph_prototype,omitempty" json:"graph_prototypes,omitempty" yaml:"graph_prototypes,omitempty"`
	MasterItem        MasterItem       `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	LldMacroPaths     []LldMacroPath   `xml:"lld_macro_paths>lld_macro_path,omitempty" json:"lld_macro_paths,omitempty" yaml:"lld_macro_paths,omitempty"`
	Preprocessing     []Step           `xml:"preprocessing>step,omitempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
}

// Filter representation
type Filter struct {
	Evaltype   string      `xml:"evaltype,omitempty" json:"evaltype,omitempty" yaml:"evaltype,omitempty"`
	Conditions []Condition `xml:"conditions>condition,omitempty" json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// GItem representation
type GItem struct {
	Host string `xml:"host,omitempty" json:"host,omitempty" yaml:"host,omitempty"`
	Key  string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Graph representation TODO
type Graph struct {
	Name       string      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Width      int         `xml:"width,omitempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height     int         `xml:"height,omitempty" json:"height,omitempty" yaml:"height,omitempty"`
	Type       string      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	YminType1  string      `xml:"ymin_type_1,omitempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
	YmaxType1  string      `xml:"ymax_type_1,omitempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	Show3d     string      `xml:"show_3d,omitempty" json:"show_3d,omitempty" yaml:"show_3d,omitempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,omitempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// GraphItem representation OK
type GraphItem struct {
	Sortorder int    `xml:"sortorder,omitempty" json:"sortorder,omitempty" yaml:"sortorder,omitempty"`
	Drawtype  string `xml:"drawtype,omitempty" json:"drawtype,omitempty" yaml:"drawtype,omitempty"`
	Color     string `xml:"color,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	Yaxisside string `xml:"yaxisside,omitempty" json:"yaxisside,omitempty" yaml:"yaxisside,omitempty"`
	CalcFnc   string `xml:"calc_fnc,omitempty" json:"calc_fnc,omitempty" yaml:"calc_fnc,omitempty"`
	Type      string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Item      GItem  `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GraphPrototype representation TODO
type GraphPrototype struct {
	Name       string      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Width      int         `xml:"width,omitempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height     int         `xml:"height,omitempty" json:"height,omitempty" yaml:"height,omitempty"`
	Type       string      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	YminType1  string      `xml:"ymin_type_1,omitempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
	YmaxType1  string      `xml:"ymax_type_1,omitempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	Show3d     string      `xml:"show_3d,omitempty" json:"show_3d,omitempty" yaml:"show_3d,omitempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,omitempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// Group representation OK
type Group struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Httptest representation TODO
type Httptest struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Item representation TODO
type Item struct {
	Name                 string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                 string        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity        string        `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid              string        `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                  string        `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                string        `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History              string        `xml:"history,omitempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends               string        `xml:"trends,omitempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	Status               string        `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	ValueType            string        `xml:"value_type,omitempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	AllowedHosts         string        `xml:"allowed_hosts,omitempty" json:"allowed_hosts,omitempty" yaml:"allowed_hosts,omitempty"`
	Units                string        `xml:"units,omitempty" json:"units,omitempty" yaml:"units,omitempty"`
	Snmpv3Contextname    string        `xml:"snmpv3_contextname,omitempty" json:"snmpv3_contextname,omitempty" yaml:"snmpv3_contextname,omitempty"`
	Snmpv3Securityname   string        `xml:"snmpv3_securityname,omitempty" json:"snmpv3_securityname,omitempty" yaml:"snmpv3_securityname,omitempty"`
	Snmpv3Securitylevel  int           `xml:"snmpv3_securitylevel,omitempty" json:"snmpv3_securitylevel,omitempty" yaml:"snmpv3_securitylevel,omitempty"`
	Snmpv3Authprotocol   int           `xml:"snmpv3_authprotocol,omitempty" json:"snmpv3_authprotocol,omitempty" yaml:"snmpv3_authprotocol,omitempty"`
	Snmpv3Authpassphrase string        `xml:"snmpv3_authpassphrase,omitempty" json:"snmpv3_authpassphrase,omitempty" yaml:"snmpv3_authpassphrase,omitempty"`
	Snmpv3Privprotocol   int           `xml:"snmpv3_privprotocol,omitempty" json:"snmpv3_privprotocol,omitempty" yaml:"snmpv3_privprotocol,omitempty"`
	Snmpv3Privpassphrase string        `xml:"snmpv3_privpassphrase,omitempty" json:"snmpv3_privpassphrase,omitempty" yaml:"snmpv3_privpassphrase,omitempty"`
	Params               string        `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
	IpmiSensor           string        `xml:"ipmi_sensor,omitempty" json:"ipmi_sensor,omitempty" yaml:"ipmi_sensor,omitempty"`
	Authtype             string        `xml:"authtype,omitempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username             string        `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password             string        `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Publickey            string        `xml:"publickey,omitempty" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	Privatekey           string        `xml:"privatekey,omitempty" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Port                 string        `xml:"port,omitempty" json:"port,omitempty" yaml:"port,omitempty"`
	Description          string        `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	InventoryLink        string        `xml:"inventory_link,omitempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications         []Application `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	ValueMap             ValueMap      `xml:"valuemap,omitempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	Logtimefmt           string        `xml:"logtimefmt,omitempty" json:"logtimefmt,omitempty" yaml:"logtimefmt,omitempty"`
	Preprocessing        []Step        `xml:"preprocessing>step,omitempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	JmxEndpoint          string        `xml:"jmx_endpoint,omitempty" json:"jmx_endpoint,omitempty" yaml:"jmx_endpoint,omitempty"`
	Timeout              string        `xml:"timeout,omitempty" json:"timeout,omitempty" yaml:"timeout,omitempty"`
	URL                  string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	MasterItem           MasterItem    `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	RetrieveMode         string        `xml:"retrieve_mode,omitempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
	Triggers             []Trigger     `xml:"triggers>trigger,omitempty" json:"triggers,omitempty" yaml:"triggers,omitempty"`
}

// ItemPrototype representation TODO
type ItemPrototype struct {
	Name                  string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type                  string        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity         string        `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid               string        `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key                   string        `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay                 string        `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History               string        `xml:"history,omitempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends                string        `xml:"trends,omitempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	ValueType             string        `xml:"value_type,omitempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	Authtype              string        `xml:"authtype,omitempty" json:"authtype,omitempty" yaml:"authtype,omitempty"`
	Username              string        `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password              string        `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Units                 string        `xml:"units,omitempty" json:"units,omitempty" yaml:"units,omitempty"`
	Params                string        `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
	Description           string        `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	ApplicationPrototypes []Application `xml:"application_prototypes>application_prototype,omitempty" json:"application_prototypes,omitempty" yaml:"application_prototypes,omitempty"`
	InventoryLink         string        `xml:"inventory_link,omitempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications          []Application `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	Logtimefmt            string        `xml:"logtimefmt,omitempty" json:"logtimefmt,omitempty" yaml:"logtimefmt,omitempty"`
	ValueMap              ValueMap      `xml:"valuemap,omitempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	Preprocessing         []Step        `xml:"preprocessing>step,omitempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	MasterItem            MasterItem    `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	TriggerPrototypes     []Trigger     `xml:"trigger_prototypes>trigger_prototype,omitempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	URL                   string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	RetrieveMode          string        `xml:"retrieve_mode,omitempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
}

// LldMacroPath representation
type LldMacroPath struct {
	LldMacro string `xml:"lld_macro,omitempty" json:"lld_macro,omitempty" yaml:"lld_macro,omitempty"`
	Path     string `xml:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty"`
}

// Macro representation OK
type Macro struct {
	Macro       string `xml:"macro,omitempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value       string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// Mapping representation OK
type Mapping struct {
	Value    string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Newvalue string `xml:"newvalue,omitempty" json:"newvalue,omitempty" yaml:"newvalue,omitempty"`
}

// MasterItem representation OK
type MasterItem struct {
	Key string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Resource representation TODO
type Resource struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Key  string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Host string `xml:"host,omitempty" json:"host,omitempty" yaml:"host,omitempty"`
}

// Screen representation OK
type Screen struct {
	Name        string       `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Hsize       int          `xml:"hsize,omitempty" json:"hsize,omitempty" yaml:"hsize,omitempty"`
	Vsize       int          `xml:"vsize,omitempty" json:"vsize,omitempty" yaml:"vsize,omitempty"`
	ScreenItems []ScreenItem `xml:"screen_items>screen_item,omitempty" json:"screen_items,omitempty" yaml:"screen_items,omitempty"`
}

// ScreenItem representation TODO
type ScreenItem struct {
	Resourcetype int      `xml:"resourcetype" json:"resourcetype,omitempty" yaml:"resourcetype,omitempty"`
	Style        int      `xml:"style" json:"style,omitempty" yaml:"style,omitempty"`
	Width        int      `xml:"width,omitempty" json:"width,omitempty" yaml:"width,omitempty"`
	Height       int      `xml:"height,omitempty" json:"height,omitempty" yaml:"height,omitempty"`
	X            int      `xml:"x" json:"x" yaml:"x"`
	Y            int      `xml:"y" json:"y" yaml:"y"`
	Colspan      int      `xml:"colspan,omitempty" json:"colspan,omitempty" yaml:"colspan,omitempty"`
	Rowspan      int      `xml:"rowspan,omitempty" json:"rowspan,omitempty" yaml:"rowspan,omitempty"`
	Elements     int      `xml:"elements,omitempty" json:"elements,omitempty" yaml:"elements,omitempty"`
	Valign       int      `xml:"valign" json:"valign,omitempty" yaml:"valign,omitempty"`
	Halign       int      `xml:"halign" json:"halign,omitempty" yaml:"halign,omitempty"`
	Dynamic      int      `xml:"dynamic" json:"dynamic,omitempty" yaml:"dynamic,omitempty"`
	SortTriggers int      `xml:"sort_triggers" json:"sort_triggers,omitempty" yaml:"sort_triggers,omitempty"`
	URL          string   `xml:"url" json:"url,omitempty" yaml:"url,omitempty"`
	Resource     Resource `xml:"resource,omitempty" json:"resource,omitempty" yaml:"resource,omitempty"`
	MaxColumns   int      `xml:"max_columns,omitempty" json:"max_columns,omitempty" yaml:"max_columns,omitempty"`
	Application  string   `xml:"application" json:"application,omitempty" yaml:"application,omitempty"`
}

// Step representation
type Step struct {
	Type               string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Params             string `xml:"params" json:"params,omitempty" yaml:"params,omitempty"`
	ErrorHandler       string `xml:"error_handler,omitempty" json:"error_handler,omitempty" yaml:"error_handler,omitempty"`
	ErrorHandlerParams string `xml:"error_handler_params,omitempty" json:"error_handler_params,omitempty" yaml:"error_handler_params,omitempty"`
}

// Tag representation OK
type Tag struct {
	Tag   string `xml:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty"`
	Value string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// Template representation TODO
type Template struct {
	Template       string          `xml:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
	Name           string          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description    string          `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Templates      []Template      `xml:"templates>template,omitempty" json:"templates,omitempty" yaml:"templates,omitempty"`
	Groups         []Group         `xml:"groups>group,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Applications   []Application   `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	Items          []Item          `xml:"items>item,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	DiscoveryRules []DiscoveryRule `xml:"discovery_rules>discovery_rule,omitempty" json:"discovery_rules,omitempty" yaml:"discovery_rules,omitempty"`
	Httptests      []Httptest      `xml:"httptests,omitempty" json:"httptests,omitempty" yaml:"httptests,omitempty"`
	Macros         []Macro         `xml:"macros>macro,omitempty" json:"macros,omitempty" yaml:"macros,omitempty"`
	Screens        []Screen        `xml:"screens>screen,omitempty" json:"screens,omitempty" yaml:"screens,omitempty"`
	Tags           []Tag           `xml:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Trigger representation OK
type Trigger struct {
	Expression         string       `xml:"expression,omitempty" json:"expression,omitempty" yaml:"expression,omitempty"`
	RecoveryMode       string       `xml:"recovery_mode,omitempty" json:"recovery_mode,omitempty" yaml:"recovery_mode,omitempty"`
	RecoveryExpression string       `xml:"recovery_expression,omitempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
	Name               string       `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Opdata             string       `xml:"opdata,omitempty" json:"opdata,omitempty" yaml:"opdata,omitempty"`
	CorrelationMode    int          `xml:"correlation_mode,omitempty" json:"correlation_mode,omitempty" yaml:"correlation_mode,omitempty"`
	CorrelationTag     string       `xml:"correlation_tag,omitempty" json:"correlation_tag,omitempty" yaml:"correlation_tag,omitempty"`
	URL                string       `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	Status             int          `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Priority           string       `xml:"priority,omitempty" json:"priority,omitempty" yaml:"priority,omitempty"`
	Description        string       `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Type               int          `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	ManualClose        string       `xml:"manual_close,omitempty" json:"manual_close,omitempty" yaml:"manual_close,omitempty"`
	Dependencies       []Dependency `xml:"dependencies>dependency,omitempty" json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Tags               []Tag        `xml:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty"`
}

// ValueMap representation OK
type ValueMap struct {
	Name     string    `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Mappings []Mapping `xml:"mappings>mapping,omitempty" json:"mappings,omitempty" yaml:"mappings,omitempty"`
}

// ZabbixExport representation
type ZabbixExport struct {
	XMLName   xml.Name   `xml:"zabbix_export" json:"-" yaml:"-"`
	Version   string     `xml:"version,omitempty" json:"version,omitempty" yaml:"version,omitempty"`
	Date      string     `xml:"date,omitempty" json:"date,omitempty" yaml:"date,omitempty"`
	Groups    []Group    `xml:"groups>group,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Templates []Template `xml:"templates>template,omitempty" json:"templates,omitempty" yaml:"templates,omitempty"`
	Triggers  []Trigger  `xml:"triggers>trigger,omitempty" json:"triggers,omitempty" yaml:"triggers,omitempty"`
	Graphs    []Graph    `xml:"graphs>graph,omitempty" json:"graphs,omitempty" yaml:"graphs,omitempty"`
	ValueMaps []ValueMap `xml:"value_maps>value_map,omitempty" json:"value_maps,omitempty" yaml:"value_maps,omitempty"`
}
