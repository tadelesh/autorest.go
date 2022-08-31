//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexmodelgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CatalogArray.
func (c CatalogArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "productArray", c.ProductArray)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatalogArray.
func (c *CatalogArray) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productArray":
			err = unpopulate(val, "ProductArray", &c.ProductArray)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CatalogArrayOfDictionary.
func (c CatalogArrayOfDictionary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "productArrayOfDictionary", c.ProductArrayOfDictionary)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatalogArrayOfDictionary.
func (c *CatalogArrayOfDictionary) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productArrayOfDictionary":
			err = unpopulate(val, "ProductArrayOfDictionary", &c.ProductArrayOfDictionary)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CatalogDictionary.
func (c CatalogDictionary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "productDictionary", c.ProductDictionary)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatalogDictionary.
func (c *CatalogDictionary) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productDictionary":
			err = unpopulate(val, "ProductDictionary", &c.ProductDictionary)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type CatalogDictionaryOfArray.
func (c CatalogDictionaryOfArray) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "productDictionaryOfArray", c.ProductDictionaryOfArray)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CatalogDictionaryOfArray.
func (c *CatalogDictionaryOfArray) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productDictionaryOfArray":
			err = unpopulate(val, "ProductDictionaryOfArray", &c.ProductDictionaryOfArray)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Product.
func (p Product) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "capacity", p.Capacity)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "display_name", p.DisplayName)
	populate(objectMap, "image", p.Image)
	populate(objectMap, "product_id", p.ProductID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Product.
func (p *Product) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "capacity":
			err = unpopulate(val, "Capacity", &p.Capacity)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "display_name":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "image":
			err = unpopulate(val, "Image", &p.Image)
			delete(rawMsg, key)
		case "product_id":
			err = unpopulate(val, "ProductID", &p.ProductID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
