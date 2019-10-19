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
	Macro     string `xml:"macro,omitempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value     string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Formulaid string `xml:"formulaid,omitempty" json:"formulaid,omitempty" yaml:"formulaid,omitempty"`
}

// Dependency representation
type Dependency struct {
	Name               string `xml:"name"`
	Expression         string `xml:"expression"`
	RecoveryExpression string `xml:"recovery_expression,omitempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
}

// DiscoveryRule representation
type DiscoveryRule struct {
	Name              string           `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type              string           `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity     string           `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid           string           `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key               string           `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay             string           `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	Description       string           `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	Filter            Filter           `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	ItemPrototypes    []ItemPrototype  `xml:"item_prototypes>item_prototype,omitempty" json:"item_prototypes,omitempty" yaml:"item_prototypes,omitempty"`
	TriggerPrototypes []Trigger        `xml:"trigger_prototypes>trigger_prototype,omitempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	GraphPrototypes   []GraphPrototype `xml:"graph_prototypes>graph_prototype,omitempty" json:"graph_prototypes,omitempty" yaml:"graph_prototypes,omitempty"`
	MasterItem        MasterItem       `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
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

// Graph representation
type Graph struct {
	Name       string      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
        YminType1  string      `xml:"ymin_type_1,omitempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
        YmaxType1  string      `xml:"ymax_type_1,omitempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	Type       string      `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,omitempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// GraphItem representation
type GraphItem struct {
	Sortorder int    `xml:"sortorder,omitempty" json:"sortorder,omitempty" yaml:"sortorder,omitempty"`
	Drawtype  string `xml:"drawtype,omitempty" json:"drawtype,omitempty" yaml:"drawtype,omitempty"`
	Color     string `xml:"color,omitempty" json:"color,omitempty" yaml:"color,omitempty"`
	Item      GItem  `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// GraphPrototype representation
type GraphPrototype struct {
	Name       string      `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	YminType1  string      `xml:"ymin_type_1,omitempty" json:"ymin_type_1,omitempty" yaml:"ymin_type_1,omitempty"`
	YmaxType1  string      `xml:"ymax_type_1,omitempty" json:"ymax_type_1,omitempty" yaml:"ymax_type_1,omitempty"`
	GraphItems []GraphItem `xml:"graph_items>graph_item,omitempty" json:"graph_items,omitempty" yaml:"graph_items,omitempty"`
}

// Group representation
type Group struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Item representation
type Item struct {
	Name          string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type          string        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity string        `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid       string        `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key           string        `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay         string        `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History       string        `xml:"history,omitempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends        string        `xml:"trends,omitempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	ValueType     string        `xml:"value_type,omitempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	Units         string        `xml:"units,omitempty" json:"units,omitempty" yaml:"units,omitempty"`
	Params        string        `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
	Description   string        `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	InventoryLink string        `xml:"inventory_link,omitempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications  []Application `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	ValueMap      ValueMap      `xml:"valuemap,omitempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	Preprocessing []Step        `xml:"preprocessing>step,omitempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	MasterItem    MasterItem    `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	URL           string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	RetrieveMode  string        `xml:"retrieve_mode,omitempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
	Triggers      []Trigger     `xml:"triggers>trigger,omitempty" json:"triggers,omitempty" yaml:"triggers,omitempty"`
}

// ItemPrototype representation
type ItemPrototype struct {
	Name              string        `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Type              string        `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	SnmpCommunity     string        `xml:"snmp_community,omitempty" json:"snmp_community,omitempty" yaml:"snmp_community,omitempty"`
	SnmpOid           string        `xml:"snmp_oid,omitempty" json:"snmp_oid,omitempty" yaml:"snmp_oid,omitempty"`
	Key               string        `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Delay             string        `xml:"delay,omitempty" json:"delay,omitempty" yaml:"delay,omitempty"`
	History           string        `xml:"history,omitempty" json:"history,omitempty" yaml:"history,omitempty"`
	Trends            string        `xml:"trends,omitempty" json:"trends,omitempty" yaml:"trends,omitempty"`
	ValueType         string        `xml:"value_type,omitempty" json:"value_type,omitempty" yaml:"value_type,omitempty"`
	Units             string        `xml:"units,omitempty" json:"units,omitempty" yaml:"units,omitempty"`
	Params            string        `xml:"params,omitempty" json:"params,omitempty" yaml:"params,omitempty"`
	Description       string        `xml:"description,omitempty" json:"decription,omitempty" yaml:"description,omitempty"`
	InventoryLink     string        `xml:"inventory_link,omitempty" json:"inventory_link,omitempty" yaml:"inventory_link,omitempty"`
	Applications      []Application `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	Preprocessing     []Step        `xml:"preprocessing>step,omitempty" json:"preprocessing,omitempty" yaml:"preprocessing,omitempty"`
	ValueMap          ValueMap      `xml:"valuemap,omitempty" json:"valuemap,omitempty" yaml:"valuemap,omitempty"`
	TriggerPrototypes []Trigger     `xml:"trigger_prototypes>trigger_prototype,omitempty" json:"trigger_prototypes,omitempty" yaml:"trigger_prototypes,omitempty"`
	MasterItem        MasterItem    `xml:"master_item,omitempty" json:"master_item,omitempty" yaml:"master_item,omitempty"`
	URL               string        `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	RetrieveMode      string        `xml:"retrieve_mode,omitempty" json:"retrieve_mode,omitempty" yaml:"retrieve_mode,omitempty"`
}

// Macro representation
type Macro struct {
	Macro       string `xml:"macro,omitempty" json:"macro,omitempty" yaml:"macro,omitempty"`
	Value       string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Description string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// Mapping representation
type Mapping struct {
	Value    string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
	Newvalue string `xml:"newvalue,omitempty" json:"newvalue,omitempty" yaml:"newvalue,omitempty"`
}

// MasterItem representation
type MasterItem struct {
	Key string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Resource representation
type Resource struct {
	Name string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Host string `xml:"host,omitempty" json:"host,omitempty" yaml:"host,omitempty"`
}

// Screen representation
type Screen struct {
	Name        string       `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Hsize       int          `xml:"hsize,omitempty" json:"hsize,omitempty" yaml:"hsize,omitempty"`
	Vsize       int          `xml:"vsize,omitempty" json:"vsize,omitempty" yaml:"vsize,omitempty"`
	ScreenItems []ScreenItem `xml:"screen_items>screen_item,omitempty" json:"screen_items,omitempty" yaml:"screen_items,omitempty"`
}

// ScreenItem representation
type ScreenItem struct {
	Resourcetype int      `xml:"resourcetype" json:"resourcetype,omitempty" yaml:"resourcetype,omitempty"`
	Style        int      `xml:"style" json:"style,omitempty" yaml:"style,omitempty"`
	Resource     Resource `xml:"resource,omitempty" json:"resource,omitempty" yaml:"resource,omitempty"`
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
	Application  string   `xml:"application" json:"application,omitempty" yaml:"application,omitempty"`
	MaxColumns   int      `xml:"max_columns,omitempty" json:"max_columns,omitempty" yaml:"max_columns,omitempty"`
}

// Step representation
type Step struct {
	Type   string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Params string `xml:"params" json:"params,omitempty" yaml:"params,omitempty"`
}

// Template representation
type Template struct {
	Template       string          `xml:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
	Name           string          `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Description    string          `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Templates      []Template      `xml:"templates>template,omitempty" json:"templates,omitempty" yaml:"templates,omitempty"`
	Groups         []Group         `xml:"groups>group,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Applications   []Application   `xml:"applications>application,omitempty" json:"applications,omitempty" yaml:"applications,omitempty"`
	Items          []Item          `xml:"items>item,omitempty" json:"items,omitempty" yaml:"items,omitempty"`
	DiscoveryRules []DiscoveryRule `xml:"discovery_rules>discovery_rule,omitempty" json:"discovery_rules,omitempty" yaml:"discovery_rules,omitempty"`
	Macros         []Macro         `xml:"macros>macro,omitempty" json:"macros,omitempty" yaml:"macros,omitempty"`
	Screens        []Screen        `xml:"screens>screen,omitempty" json:"screens,omitempty" yaml:"screens,omitempty"`
}

// Trigger representation
type Trigger struct {
	Expression         string       `xml:"expression,omitempty" json:"expression,omitempty" yaml:"expression,omitempty"`
	RecoveryMode       string       `xml:"recovery_mode,omitempty" json:"recovery_mode,omitempty" yaml:"recovery_mode,omitempty"`
	RecoveryExpression string       `xml:"recovery_expression,omitempty" json:"recovery_expression,omitempty" yaml:"recovery_expression,omitempty"`
	Name               string       `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
	Opdata             string       `xml:"opdata,omitempty" json:"opdata,omitempty" yaml:"opdata,omitempty"`
	Priority           string       `xml:"priority,omitempty" json:"priority,omitempty" yaml:"priority,omitempty"`
	Description        string       `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	ManualClose        string       `xml:"manual_close,omitempty" json:"manual_close,omitempty" yaml:"manual_close,omitempty"`
	Dependencies       []Dependency `xml:"dependencies>dependency,omitempty" json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
}

// ValueMap representation
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
	Graphs    []Graph    `xml:"graphs>graph,omitempty" json:"graphs,omitempty" yaml:"graphs,omitempty"`
	ValueMaps []ValueMap `xml:"value_maps>value_map,omitempty" json:"value_maps,omitempty" yaml:"value_maps,omitempty"`
}
