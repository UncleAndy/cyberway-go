package cyberway

import (
	"encoding/json"
	"fmt"
	"io"
)

// see: libraries/chain/contracts/abi_serializer.cpp:53...
// see: libraries/chain/include/eosio/chain/contracts/types.hpp:100
type ABI struct {
	Version          string            	`json:"version"`
	Types            []ABIType         	`json:"types,omitempty"`
	Structs          []StructDef       	`json:"structs,omitempty"`
	Actions          []ActionDef       	`json:"actions,omitempty"`
	Events			 []EventDef			`json:"events,omitempty"`
	Tables           []TableDef        	`json:"tables,omitempty"`
	ErrorMessages    []ABIErrorMessage 	`json:"error_messages,omitempty"`
	Extensions       []ExtensionEntry  	`json:"abi_extensions,omitempty"`
	Variants		 []VariantDef		`json:"variants,omitempty"`
}

func NewABI(r io.Reader) (*ABI, error) {
	abi := &ABI{}
	abiDecoder := json.NewDecoder(r)
	err := abiDecoder.Decode(abi)
	if err != nil {
		return nil, fmt.Errorf("read abi: %s", err)
	}

	return abi, nil

}

func (a *ABI) ActionForName(name ActionName) *ActionDef {
	for _, a := range a.Actions {
		if a.Name == name {
			return &a
		}
	}
	return nil
}

func (a *ABI) StructForName(name string) *StructDef {
	for _, s := range a.Structs {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

func (a *ABI) TableForName(name TableName) *TableDef {
	for _, s := range a.Tables {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

func (a *ABI) VariantForName(name string) *VariantDef {
	for _, s := range a.Variants {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

func (a *ABI) TypeNameForNewTypeName(typeName string) (resolvedTypeName string, isAlias bool) {
	for _, t := range a.Types {
		if t.NewTypeName == typeName {
			return t.Type, true
		}
	}
	return typeName, false
}

// +
type ABIType struct {
	NewTypeName string `json:"new_type_name"`
	Type        string `json:"type"`
}

// +
type StructDef struct {
	Name   string     `json:"name"`
	Base   string     `json:"base"`
	Fields []FieldDef `json:"fields,omitempty"`
}

// +
type FieldDef struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// ++
type ActionDef struct {
	Name              ActionName `json:"name"`
	Type              string     `json:"type"`
}

// TableDef defines a table. See libraries/chain/include/eosio/chain/contracts/types.hpp:78
// ++
type TableDef struct {
	Name      TableName `json:"name"`
	Type      string    `json:"type"`
	ScopeType string    `json:"scope_type"`
	Indexes   []string  `json:"indexes,omitempty"`
}
// +++
type IndexDef struct {
	Name		IndexName	`json:"name"`
	Unique		bool		`json:"unique"`
	Orders		[]OrderDef	`json:"orders"`
}
// +++
type OrderDef struct {
	Field		string		`json:"field"`
	Order		string		`json:"order"`
}

// ClausePair represents clauses, related to Ricardian Contracts.
// +
type ClausePair struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// +
type ABIErrorMessage struct {
	Code    Uint64 `json:"error_code"`
	Message string `json:"error_msg"`
}

// +++
type EventDef struct {
	Name	string		`json:"name"`
	Type	string		`json:"type"`
}

// +++
type VariantDef struct {
	Name	string		`json:"name"`
	Types	[]string	`json:"types,omitempty"`
}

// +++
type ExtensionEntry struct {
	Tag		uint16		`json:"tag"`
	Value	HexBytes	`json:"value"`
}