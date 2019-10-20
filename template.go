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
	Macro     string `xml:"macro,allowempty" json:"macro,allowempty" yaml:"macro,allowempty"`
	Value     string `xml:"value,allowempty" json:"value,allowempty" yaml:"value,allowempty"`
	Operator  string `xml:"operator,allowempty" json:"operator,allowempty" yaml:"operator,allowempty"`
	Formulaid string `xml:"formulaid,allowempty" json:"formulaid,allowempty" yaml:"formulaid,allowempty"`
}

// Dependency representation
type Dependency struct {
	Name               string `xml:"name"`
	Expression         string `xml:"expression"`
	RecoveryExpression string `xml:"recovery_expression,allowempty" json:"recovery_expression,allowempty" yaml:"recovery_expression,allowempty"`
}

// DiscoveryRule representation
type DiscoveryRule struct {
	Name              string           `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Type              string           `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	SnmpCommunity     string           `xml:"snmp_community,allowempty" json:"snmp_community,allowempty" yaml:"snmp_community,allowempty"`
	SnmpOid           string           `xml:"snmp_oid,allowempty" json:"snmp_oid,allowempty" yaml:"snmp_oid,allowempty"`
	Key               string           `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
	Delay             string           `xml:"delay,allowempty" json:"delay,allowempty" yaml:"delay,allowempty"`
	Filter            Filter           `xml:"filter,allowempty" json:"filter,allowempty" yaml:"filter,allowempty"`
	Description       string           `xml:"description,allowempty" json:"decription,allowempty" yaml:"description,allowempty"`
	ItemPrototypes    []ItemPrototype  `xml:"item_prototypes>item_prototype,allowempty" json:"item_prototypes,allowempty" yaml:"item_prototypes,allowempty"`
	TriggerPrototypes []Trigger        `xml:"trigger_prototypes>trigger_prototype,allowempty" json:"trigger_prototypes,allowempty" yaml:"trigger_prototypes,allowempty"`
	GraphPrototypes   []GraphPrototype `xml:"graph_prototypes>graph_prototype,allowempty" json:"graph_prototypes,allowempty" yaml:"graph_prototypes,allowempty"`
	MasterItem        MasterItem       `xml:"master_item,allowempty" json:"master_item,allowempty" yaml:"master_item,allowempty"`
	LldMacroPaths     []LldMacroPath   `xml:"lld_macro_paths>lld_macro_path,allowempty" json:"lld_macro_paths,allowempty" yaml:"lld_macro_paths,allowempty"`
	Preprocessing     []Step           `xml:"preprocessing>step,allowempty" json:"preprocessing,allowempty" yaml:"preprocessing,allowempty"`
}

// Filter representation
type Filter struct {
	Evaltype   string      `xml:"evaltype,allowempty" json:"evaltype,allowempty" yaml:"evaltype,allowempty"`
	Conditions []Condition `xml:"conditions>condition,allowempty" json:"conditions,allowempty" yaml:"conditions,allowempty"`
}

// GItem representation
type GItem struct {
	Host string `xml:"host,allowempty" json:"host,allowempty" yaml:"host,allowempty"`
	Key  string `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
}

// Graph representation
type Graph struct {
	Name       string      `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Width      int         `xml:"width,allowempty" json:"width,allowempty" yaml:"width,allowempty"`
	Height     int         `xml:"height,allowempty" json:"height,allowempty" yaml:"height,allowempty"`
	Type       string      `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	YminType1  string      `xml:"ymin_type_1,allowempty" json:"ymin_type_1,allowempty" yaml:"ymin_type_1,allowempty"`
	YmaxType1  string      `xml:"ymax_type_1,allowempty" json:"ymax_type_1,allowempty" yaml:"ymax_type_1,allowempty"`
	Show3d     string      `xml:"show_3d,allowempty" json:"show_3d,allowempty" yaml:"show_3d,allowempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,allowempty" json:"graph_items,allowempty" yaml:"graph_items,allowempty"`
}

// GraphItem representation
type GraphItem struct {
	Sortorder int    `xml:"sortorder,allowempty" json:"sortorder,allowempty" yaml:"sortorder,allowempty"`
	Drawtype  string `xml:"drawtype,allowempty" json:"drawtype,allowempty" yaml:"drawtype,allowempty"`
	Color     string `xml:"color,allowempty" json:"color,allowempty" yaml:"color,allowempty"`
	Yaxisside string `xml:"yaxisside,allowempty" json:"yaxisside,allowempty" yaml:"yaxisside,allowempty"`
	CalcFnc   string `xml:"calc_fnc,allowempty" json:"calc_fnc,allowempty" yaml:"calc_fnc,allowempty"`
	Type      string `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	Item      GItem  `xml:"item,allowempty" json:"item,allowempty" yaml:"item,allowempty"`
}

// GraphPrototype representation
type GraphPrototype struct {
	Name       string      `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Width      int         `xml:"width,allowempty" json:"width,allowempty" yaml:"width,allowempty"`
	Height     int         `xml:"height,allowempty" json:"height,allowempty" yaml:"height,allowempty"`
	Type       string      `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	YminType1  string      `xml:"ymin_type_1,allowempty" json:"ymin_type_1,allowempty" yaml:"ymin_type_1,allowempty"`
	YmaxType1  string      `xml:"ymax_type_1,allowempty" json:"ymax_type_1,allowempty" yaml:"ymax_type_1,allowempty"`
	Show3d     string      `xml:"show_3d,allowempty" json:"show_3d,allowempty" yaml:"show_3d,allowempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,allowempty" json:"graph_items,allowempty" yaml:"graph_items,allowempty"`
}

// Group representation
type Group struct {
	Name string `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
}

// Item representation
type Item struct {
	Name          string        `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Type          string        `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	SnmpCommunity string        `xml:"snmp_community,allowempty" json:"snmp_community,allowempty" yaml:"snmp_community,allowempty"`
	SnmpOid       string        `xml:"snmp_oid,allowempty" json:"snmp_oid,allowempty" yaml:"snmp_oid,allowempty"`
	Key           string        `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
	Delay         string        `xml:"delay,allowempty" json:"delay,allowempty" yaml:"delay,allowempty"`
	History       string        `xml:"history,allowempty" json:"history,allowempty" yaml:"history,allowempty"`
	Status        string        `xml:"status,allowempty" json:"status,allowempty" yaml:"status,allowempty"`
	Trends        string        `xml:"trends,allowempty" json:"trends,allowempty" yaml:"trends,allowempty"`
	ValueType     string        `xml:"value_type,allowempty" json:"value_type,allowempty" yaml:"value_type,allowempty"`
	Authtype      string        `xml:"authtype,allowempty" json:"authtype,allowempty" yaml:"authtype,allowempty"`
	Username      string        `xml:"username,allowempty" json:"username,allowempty" yaml:"username,allowempty"`
	Password      string        `xml:"password,allowempty" json:"password,allowempty" yaml:"password,allowempty"`
	Units         string        `xml:"units,allowempty" json:"units,allowempty" yaml:"units,allowempty"`
	Params        string        `xml:"params,allowempty" json:"params,allowempty" yaml:"params,allowempty"`
	Description   string        `xml:"description,allowempty" json:"decription,allowempty" yaml:"description,allowempty"`
	InventoryLink string        `xml:"inventory_link,allowempty" json:"inventory_link,allowempty" yaml:"inventory_link,allowempty"`
	Applications  []Application `xml:"applications>application,allowempty" json:"applications,allowempty" yaml:"applications,allowempty"`
	Logtimefmt    string        `xml:"logtimefmt,allowempty" json:"logtimefmt,allowempty" yaml:"logtimefmt,allowempty"`
	ValueMap      ValueMap      `xml:"valuemap,allowempty" json:"valuemap,allowempty" yaml:"valuemap,allowempty"`
	Preprocessing []Step        `xml:"preprocessing>step,allowempty" json:"preprocessing,allowempty" yaml:"preprocessing,allowempty"`
	MasterItem    MasterItem    `xml:"master_item,allowempty" json:"master_item,allowempty" yaml:"master_item,allowempty"`
	URL           string        `xml:"url,allowempty" json:"url,allowempty" yaml:"url,allowempty"`
	RetrieveMode  string        `xml:"retrieve_mode,allowempty" json:"retrieve_mode,allowempty" yaml:"retrieve_mode,allowempty"`
	Triggers      []Trigger     `xml:"triggers>trigger,allowempty" json:"triggers,allowempty" yaml:"triggers,allowempty"`
}

// ItemPrototype representation
type ItemPrototype struct {
	Name                  string        `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Type                  string        `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	SnmpCommunity         string        `xml:"snmp_community,allowempty" json:"snmp_community,allowempty" yaml:"snmp_community,allowempty"`
	SnmpOid               string        `xml:"snmp_oid,allowempty" json:"snmp_oid,allowempty" yaml:"snmp_oid,allowempty"`
	Key                   string        `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
	Delay                 string        `xml:"delay,allowempty" json:"delay,allowempty" yaml:"delay,allowempty"`
	History               string        `xml:"history,allowempty" json:"history,allowempty" yaml:"history,allowempty"`
	Trends                string        `xml:"trends,allowempty" json:"trends,allowempty" yaml:"trends,allowempty"`
	ValueType             string        `xml:"value_type,allowempty" json:"value_type,allowempty" yaml:"value_type,allowempty"`
	Authtype              string        `xml:"authtype,allowempty" json:"authtype,allowempty" yaml:"authtype,allowempty"`
	Username              string        `xml:"username,allowempty" json:"username,allowempty" yaml:"username,allowempty"`
	Password              string        `xml:"password,allowempty" json:"password,allowempty" yaml:"password,allowempty"`
	Units                 string        `xml:"units,allowempty" json:"units,allowempty" yaml:"units,allowempty"`
	Params                string        `xml:"params,allowempty" json:"params,allowempty" yaml:"params,allowempty"`
	Description           string        `xml:"description,allowempty" json:"decription,allowempty" yaml:"description,allowempty"`
	ApplicationPrototypes []Application `xml:"application_prototypes>application_prototype,allowempty" json:"application_prototypes,allowempty" yaml:"application_prototypes,allowempty"`
	InventoryLink         string        `xml:"inventory_link,allowempty" json:"inventory_link,allowempty" yaml:"inventory_link,allowempty"`
	Applications          []Application `xml:"applications>application,allowempty" json:"applications,allowempty" yaml:"applications,allowempty"`
	Logtimefmt            string        `xml:"logtimefmt,allowempty" json:"logtimefmt,allowempty" yaml:"logtimefmt,allowempty"`
	ValueMap              ValueMap      `xml:"valuemap,allowempty" json:"valuemap,allowempty" yaml:"valuemap,allowempty"`
	Preprocessing         []Step        `xml:"preprocessing>step,allowempty" json:"preprocessing,allowempty" yaml:"preprocessing,allowempty"`
	MasterItem            MasterItem    `xml:"master_item,allowempty" json:"master_item,allowempty" yaml:"master_item,allowempty"`
	TriggerPrototypes     []Trigger     `xml:"trigger_prototypes>trigger_prototype,allowempty" json:"trigger_prototypes,allowempty" yaml:"trigger_prototypes,allowempty"`
	URL                   string        `xml:"url,allowempty" json:"url,allowempty" yaml:"url,allowempty"`
	RetrieveMode          string        `xml:"retrieve_mode,allowempty" json:"retrieve_mode,allowempty" yaml:"retrieve_mode,allowempty"`
}

// LldMacroPath representation
type LldMacroPath struct {
	LldMacro string `xml:"lld_macro,allowempty" json:"lld_macro,allowempty" yaml:"lld_macro,allowempty"`
	Path     string `xml:"path,allowempty" json:"path,allowempty" yaml:"path,allowempty"`
}

// Macro representation
type Macro struct {
	Macro       string `xml:"macro,allowempty" json:"macro,allowempty" yaml:"macro,allowempty"`
	Value       string `xml:"value,allowempty" json:"value,allowempty" yaml:"value,allowempty"`
	Description string `xml:"description,allowempty" json:"description,allowempty" yaml:"description,allowempty"`
}

// Mapping representation
type Mapping struct {
	Value    string `xml:"value,allowempty" json:"value,allowempty" yaml:"value,allowempty"`
	Newvalue string `xml:"newvalue,allowempty" json:"newvalue,allowempty" yaml:"newvalue,allowempty"`
}

// MasterItem representation
type MasterItem struct {
	Key string `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
}

// Resource representation
type Resource struct {
	Name string `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Key  string `xml:"key,allowempty" json:"key,allowempty" yaml:"key,allowempty"`
	Host string `xml:"host,allowempty" json:"host,allowempty" yaml:"host,allowempty"`
}

// Screen representation
type Screen struct {
	Name        string       `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Hsize       int          `xml:"hsize,allowempty" json:"hsize,allowempty" yaml:"hsize,allowempty"`
	Vsize       int          `xml:"vsize,allowempty" json:"vsize,allowempty" yaml:"vsize,allowempty"`
	ScreenItems []ScreenItem `xml:"screen_items>screen_item,allowempty" json:"screen_items,allowempty" yaml:"screen_items,allowempty"`
}

// ScreenItem representation
type ScreenItem struct {
	Resourcetype int      `xml:"resourcetype" json:"resourcetype,allowempty" yaml:"resourcetype,allowempty"`
	Style        int      `xml:"style" json:"style,allowempty" yaml:"style,allowempty"`
	Resource     Resource `xml:"resource,allowempty" json:"resource,allowempty" yaml:"resource,allowempty"`
	Width        int      `xml:"width,allowempty" json:"width,allowempty" yaml:"width,allowempty"`
	Height       int      `xml:"height,allowempty" json:"height,allowempty" yaml:"height,allowempty"`
	X            int      `xml:"x" json:"x" yaml:"x"`
	Y            int      `xml:"y" json:"y" yaml:"y"`
	Colspan      int      `xml:"colspan,allowempty" json:"colspan,allowempty" yaml:"colspan,allowempty"`
	Rowspan      int      `xml:"rowspan,allowempty" json:"rowspan,allowempty" yaml:"rowspan,allowempty"`
	Elements     int      `xml:"elements,allowempty" json:"elements,allowempty" yaml:"elements,allowempty"`
	Valign       int      `xml:"valign" json:"valign,allowempty" yaml:"valign,allowempty"`
	Halign       int      `xml:"halign" json:"halign,allowempty" yaml:"halign,allowempty"`
	Dynamic      int      `xml:"dynamic" json:"dynamic,allowempty" yaml:"dynamic,allowempty"`
	SortTriggers int      `xml:"sort_triggers" json:"sort_triggers,allowempty" yaml:"sort_triggers,allowempty"`
	URL          string   `xml:"url" json:"url,allowempty" yaml:"url,allowempty"`
	Application  string   `xml:"application" json:"application,allowempty" yaml:"application,allowempty"`
	MaxColumns   int      `xml:"max_columns,allowempty" json:"max_columns,allowempty" yaml:"max_columns,allowempty"`
}

// Step representation
type Step struct {
	Type               string `xml:"type,allowempty" json:"type,allowempty" yaml:"type,allowempty"`
	Params             string `xml:"params" json:"params,allowempty" yaml:"params,allowempty"`
	ErrorHandler       string `xml:"error_handler,allowempty" json:"error_handler,allowempty" yaml:"error_handler,allowempty"`
	ErrorHandlerParams string `xml:"error_handler_params,allowempty" json:"error_handler_params,allowempty" yaml:"error_handler_params,allowempty"`
}

// Template representation
type Template struct {
	Template       string          `xml:"template,allowempty" json:"template,allowempty" yaml:"template,allowempty"`
	Name           string          `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Description    string          `xml:"description,allowempty" json:"description,allowempty" yaml:"description,allowempty"`
	Templates      []Template      `xml:"templates>template,allowempty" json:"templates,allowempty" yaml:"templates,allowempty"`
	Groups         []Group         `xml:"groups>group,allowempty" json:"groups,allowempty" yaml:"groups,allowempty"`
	Applications   []Application   `xml:"applications>application,allowempty" json:"applications,allowempty" yaml:"applications,allowempty"`
	Items          []Item          `xml:"items>item,allowempty" json:"items,allowempty" yaml:"items,allowempty"`
	DiscoveryRules []DiscoveryRule `xml:"discovery_rules>discovery_rule,allowempty" json:"discovery_rules,allowempty" yaml:"discovery_rules,allowempty"`
	Macros         []Macro         `xml:"macros>macro,allowempty" json:"macros,allowempty" yaml:"macros,allowempty"`
	Screens        []Screen        `xml:"screens>screen,allowempty" json:"screens,allowempty" yaml:"screens,allowempty"`
}

// Trigger representation
type Trigger struct {
	Expression         string       `xml:"expression,allowempty" json:"expression,allowempty" yaml:"expression,allowempty"`
	RecoveryMode       string       `xml:"recovery_mode,allowempty" json:"recovery_mode,allowempty" yaml:"recovery_mode,allowempty"`
	RecoveryExpression string       `xml:"recovery_expression,allowempty" json:"recovery_expression,allowempty" yaml:"recovery_expression,allowempty"`
	Name               string       `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Opdata             string       `xml:"opdata,allowempty" json:"opdata,allowempty" yaml:"opdata,allowempty"`
	Priority           string       `xml:"priority,allowempty" json:"priority,allowempty" yaml:"priority,allowempty"`
	Description        string       `xml:"description,allowempty" json:"description,allowempty" yaml:"description,allowempty"`
	ManualClose        string       `xml:"manual_close,allowempty" json:"manual_close,allowempty" yaml:"manual_close,allowempty"`
	Dependencies       []Dependency `xml:"dependencies>dependency,allowempty" json:"dependencies,allowempty" yaml:"dependencies,allowempty"`
}

// ValueMap representation
type ValueMap struct {
	Name     string    `xml:"name,allowempty" json:"name,allowempty" yaml:"name,allowempty"`
	Mappings []Mapping `xml:"mappings>mapping,allowempty" json:"mappings,allowempty" yaml:"mappings,allowempty"`
}

// ZabbixExport representation
type ZabbixExport struct {
	XMLName   xml.Name   `xml:"zabbix_export" json:"-" yaml:"-"`
	Version   string     `xml:"version,allowempty" json:"version,allowempty" yaml:"version,allowempty"`
	Date      string     `xml:"date,allowempty" json:"date,allowempty" yaml:"date,allowempty"`
	Groups    []Group    `xml:"groups>group,allowempty" json:"groups,allowempty" yaml:"groups,allowempty"`
	Templates []Template `xml:"templates>template,allowempty" json:"templates,allowempty" yaml:"templates,allowempty"`
	Triggers  []Trigger  `xml:"triggers>trigger,allowempty" json:"triggers,allowempty" yaml:"triggers,allowempty"`
	Graphs    []Graph    `xml:"graphs>graph,allowempty" json:"graphs,allowempty" yaml:"graphs,allowempty"`
	ValueMaps []ValueMap `xml:"value_maps>value_map,allowempty" json:"value_maps,allowempty" yaml:"value_maps,allowempty"`
}
