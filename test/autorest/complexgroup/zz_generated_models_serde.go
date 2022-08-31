//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package complexgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ArrayWrapper.
func (a ArrayWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "array", a.Array)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArrayWrapper.
func (a *ArrayWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "array":
			err = unpopulate(val, "Array", &a.Array)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Basic.
func (b Basic) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "color", b.Color)
	populate(objectMap, "id", b.ID)
	populate(objectMap, "name", b.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Basic.
func (b *Basic) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "color":
			err = unpopulate(val, "Color", &b.Color)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &b.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &b.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type BooleanWrapper.
func (b BooleanWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field_false", b.FieldFalse)
	populate(objectMap, "field_true", b.FieldTrue)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type BooleanWrapper.
func (b *BooleanWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field_false":
			err = unpopulate(val, "FieldFalse", &b.FieldFalse)
			delete(rawMsg, key)
		case "field_true":
			err = unpopulate(val, "FieldTrue", &b.FieldTrue)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ByteWrapper.
func (b ByteWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateByteArray(objectMap, "field", b.Field, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ByteWrapper.
func (b *ByteWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", b, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = runtime.DecodeByteArray(string(val), &b.Field, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", b, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Cookiecuttershark.
func (c Cookiecuttershark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "age", c.Age)
	populateTimeRFC3339(objectMap, "birthday", c.Birthday)
	objectMap["fishtype"] = "cookiecuttershark"
	populate(objectMap, "length", c.Length)
	populate(objectMap, "siblings", c.Siblings)
	populate(objectMap, "species", c.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Cookiecuttershark.
func (c *Cookiecuttershark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &c.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &c.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &c.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &c.Length)
			delete(rawMsg, key)
		case "siblings":
			c.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &c.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DateWrapper.
func (d DateWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateDateType(objectMap, "field", d.Field)
	populateDateType(objectMap, "leap", d.Leap)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateWrapper.
func (d *DateWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateDateType(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "leap":
			err = unpopulateDateType(val, "Leap", &d.Leap)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DatetimeWrapper.
func (d DatetimeWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "field", d.Field)
	populateTimeRFC3339(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DatetimeWrapper.
func (d *DatetimeWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC3339(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC3339(val, "Now", &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Datetimerfc1123Wrapper.
func (d Datetimerfc1123Wrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC1123(objectMap, "field", d.Field)
	populateTimeRFC1123(objectMap, "now", d.Now)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Datetimerfc1123Wrapper.
func (d *Datetimerfc1123Wrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulateTimeRFC1123(val, "Field", &d.Field)
			delete(rawMsg, key)
		case "now":
			err = unpopulateTimeRFC1123(val, "Now", &d.Now)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DictionaryWrapper.
func (d DictionaryWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "defaultProgram", d.DefaultProgram)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DictionaryWrapper.
func (d *DictionaryWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "defaultProgram":
			err = unpopulate(val, "DefaultProgram", &d.DefaultProgram)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Dog.
func (d Dog) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "food", d.Food)
	populate(objectMap, "id", d.ID)
	populate(objectMap, "name", d.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Dog.
func (d *Dog) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "food":
			err = unpopulate(val, "Food", &d.Food)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &d.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &d.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DotFish.
func (d DotFish) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["fish.type"] = d.FishType
	populate(objectMap, "species", d.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFish.
func (d *DotFish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			err = unpopulate(val, "FishType", &d.FishType)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &d.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DotFishMarket.
func (d DotFishMarket) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "fishes", d.Fishes)
	populate(objectMap, "salmons", d.Salmons)
	populate(objectMap, "sampleFish", d.SampleFish)
	populate(objectMap, "sampleSalmon", d.SampleSalmon)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotFishMarket.
func (d *DotFishMarket) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishes":
			d.Fishes, err = unmarshalDotFishClassificationArray(val)
			delete(rawMsg, key)
		case "salmons":
			err = unpopulate(val, "Salmons", &d.Salmons)
			delete(rawMsg, key)
		case "sampleFish":
			d.SampleFish, err = unmarshalDotFishClassification(val)
			delete(rawMsg, key)
		case "sampleSalmon":
			err = unpopulate(val, "SampleSalmon", &d.SampleSalmon)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DotSalmon.
func (d DotSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["fish.type"] = "DotSalmon"
	populate(objectMap, "iswild", d.Iswild)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "species", d.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DotSalmon.
func (d *DotSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fish.type":
			err = unpopulate(val, "FishType", &d.FishType)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &d.Iswild)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &d.Location)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &d.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DoubleWrapper.
func (d DoubleWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field1", d.Field1)
	populate(objectMap, "field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose", d.Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DoubleWrapper.
func (d *DoubleWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field1":
			err = unpopulate(val, "Field1", &d.Field1)
			delete(rawMsg, key)
		case "field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose":
			err = unpopulate(val, "Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose", &d.Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DurationWrapper.
func (d DurationWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field", d.Field)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DurationWrapper.
func (d *DurationWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field":
			err = unpopulate(val, "Field", &d.Field)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Fish.
func (f Fish) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["fishtype"] = f.Fishtype
	populate(objectMap, "length", f.Length)
	populate(objectMap, "siblings", f.Siblings)
	populate(objectMap, "species", f.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Fish.
func (f *Fish) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			err = unpopulate(val, "Fishtype", &f.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &f.Length)
			delete(rawMsg, key)
		case "siblings":
			f.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &f.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FloatWrapper.
func (f FloatWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field1", f.Field1)
	populate(objectMap, "field2", f.Field2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FloatWrapper.
func (f *FloatWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field1":
			err = unpopulate(val, "Field1", &f.Field1)
			delete(rawMsg, key)
		case "field2":
			err = unpopulate(val, "Field2", &f.Field2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Goblinshark.
func (g Goblinshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "age", g.Age)
	populateTimeRFC3339(objectMap, "birthday", g.Birthday)
	populate(objectMap, "color", g.Color)
	objectMap["fishtype"] = "goblin"
	populate(objectMap, "jawsize", g.Jawsize)
	populate(objectMap, "length", g.Length)
	populate(objectMap, "siblings", g.Siblings)
	populate(objectMap, "species", g.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Goblinshark.
func (g *Goblinshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &g.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &g.Birthday)
			delete(rawMsg, key)
		case "color":
			err = unpopulate(val, "Color", &g.Color)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &g.Fishtype)
			delete(rawMsg, key)
		case "jawsize":
			err = unpopulate(val, "Jawsize", &g.Jawsize)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &g.Length)
			delete(rawMsg, key)
		case "siblings":
			g.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &g.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IntWrapper.
func (i IntWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field1", i.Field1)
	populate(objectMap, "field2", i.Field2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IntWrapper.
func (i *IntWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field1":
			err = unpopulate(val, "Field1", &i.Field1)
			delete(rawMsg, key)
		case "field2":
			err = unpopulate(val, "Field2", &i.Field2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LongWrapper.
func (l LongWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "field1", l.Field1)
	populate(objectMap, "field2", l.Field2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LongWrapper.
func (l *LongWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "field1":
			err = unpopulate(val, "Field1", &l.Field1)
			delete(rawMsg, key)
		case "field2":
			err = unpopulate(val, "Field2", &l.Field2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MyBaseHelperType.
func (m MyBaseHelperType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "propBH1", m.PropBH1)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseHelperType.
func (m *MyBaseHelperType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "propBH1":
			err = unpopulate(val, "PropBH1", &m.PropBH1)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MyBaseType.
func (m MyBaseType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "helper", m.Helper)
	objectMap["kind"] = m.Kind
	populate(objectMap, "propB1", m.PropB1)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyBaseType.
func (m *MyBaseType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			err = unpopulate(val, "Helper", &m.Helper)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &m.Kind)
			delete(rawMsg, key)
		case "propB1":
			err = unpopulate(val, "PropB1", &m.PropB1)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MyDerivedType.
func (m MyDerivedType) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "helper", m.Helper)
	objectMap["kind"] = MyKindKind1
	populate(objectMap, "propB1", m.PropB1)
	populate(objectMap, "propD1", m.PropD1)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MyDerivedType.
func (m *MyDerivedType) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "helper":
			err = unpopulate(val, "Helper", &m.Helper)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &m.Kind)
			delete(rawMsg, key)
		case "propB1":
			err = unpopulate(val, "PropB1", &m.PropB1)
			delete(rawMsg, key)
		case "propD1":
			err = unpopulate(val, "PropD1", &m.PropD1)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReadonlyObj.
func (r ReadonlyObj) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "size", r.Size)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReadonlyObj.
func (r *ReadonlyObj) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &r.ID)
			delete(rawMsg, key)
		case "size":
			err = unpopulate(val, "Size", &r.Size)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Salmon.
func (s Salmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	objectMap["fishtype"] = "salmon"
	populate(objectMap, "iswild", s.Iswild)
	populate(objectMap, "length", s.Length)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Salmon.
func (s *Salmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &s.Iswild)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Sawshark.
func (s Sawshark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "age", s.Age)
	populateTimeRFC3339(objectMap, "birthday", s.Birthday)
	objectMap["fishtype"] = "sawshark"
	populate(objectMap, "length", s.Length)
	populateByteArray(objectMap, "picture", s.Picture, runtime.Base64StdFormat)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Sawshark.
func (s *Sawshark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &s.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &s.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "picture":
			err = runtime.DecodeByteArray(string(val), &s.Picture, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Shark.
func (s Shark) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "age", s.Age)
	populateTimeRFC3339(objectMap, "birthday", s.Birthday)
	objectMap["fishtype"] = "shark"
	populate(objectMap, "length", s.Length)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Shark.
func (s *Shark) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "age":
			err = unpopulate(val, "Age", &s.Age)
			delete(rawMsg, key)
		case "birthday":
			err = unpopulateTimeRFC3339(val, "Birthday", &s.Birthday)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Siamese.
func (s Siamese) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "breed", s.Breed)
	populate(objectMap, "color", s.Color)
	populate(objectMap, "hates", s.Hates)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Siamese.
func (s *Siamese) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "breed":
			err = unpopulate(val, "Breed", &s.Breed)
			delete(rawMsg, key)
		case "color":
			err = unpopulate(val, "Color", &s.Color)
			delete(rawMsg, key)
		case "hates":
			err = unpopulate(val, "Hates", &s.Hates)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SmartSalmon.
func (s SmartSalmon) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "college_degree", s.CollegeDegree)
	objectMap["fishtype"] = "smart_salmon"
	populate(objectMap, "iswild", s.Iswild)
	populate(objectMap, "length", s.Length)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "siblings", s.Siblings)
	populate(objectMap, "species", s.Species)
	if s.AdditionalProperties != nil {
		for key, val := range s.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SmartSalmon.
func (s *SmartSalmon) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "college_degree":
			err = unpopulate(val, "CollegeDegree", &s.CollegeDegree)
			delete(rawMsg, key)
		case "fishtype":
			err = unpopulate(val, "Fishtype", &s.Fishtype)
			delete(rawMsg, key)
		case "iswild":
			err = unpopulate(val, "Iswild", &s.Iswild)
			delete(rawMsg, key)
		case "length":
			err = unpopulate(val, "Length", &s.Length)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "siblings":
			s.Siblings, err = unmarshalFishClassificationArray(val)
			delete(rawMsg, key)
		case "species":
			err = unpopulate(val, "Species", &s.Species)
			delete(rawMsg, key)
		default:
			if s.AdditionalProperties == nil {
				s.AdditionalProperties = map[string]any{}
			}
			if val != nil {
				var aux any
				err = json.Unmarshal(val, &aux)
				s.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type StringWrapper.
func (s StringWrapper) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "empty", s.Empty)
	populate(objectMap, "field", s.Field)
	populate(objectMap, "null", s.Null)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type StringWrapper.
func (s *StringWrapper) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "empty":
			err = unpopulate(val, "Empty", &s.Empty)
			delete(rawMsg, key)
		case "field":
			err = unpopulate(val, "Field", &s.Field)
			delete(rawMsg, key)
		case "null":
			err = unpopulate(val, "Null", &s.Null)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
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

func populateByteArray(m map[string]any, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
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
