// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ISiKDiagnose is documented here https://gematik.de/fhir/ISiK/StructureDefinition/ISiKDiagnose
type ISiKDiagnose struct {
	Id            *string     `bson:"id,omitempty" json:"id,omitempty"`
	Meta          *Meta       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules *string     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language      *string     `bson:"language,omitempty" json:"language,omitempty"`
	Text          *Narrative  `bson:"text,omitempty" json:"text,omitempty"`
	Extension     []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}
type OtherISiKDiagnose ISiKDiagnose

// MarshalJSON marshals the given ISiKDiagnose as JSON into a byte slice
func (r ISiKDiagnose) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherISiKDiagnose
		ResourceType string `json:"resourceType"`
	}{
		OtherISiKDiagnose: OtherISiKDiagnose(r),
		ResourceType:      "ISiKDiagnose",
	})
}

// UnmarshalISiKDiagnose unmarshals a ISiKDiagnose.
func UnmarshalISiKDiagnose(b []byte) (ISiKDiagnose, error) {
	var iSiKDiagnose ISiKDiagnose
	if err := json.Unmarshal(b, &iSiKDiagnose); err != nil {
		return iSiKDiagnose, err
	}
	return iSiKDiagnose, nil
}
