// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v62/common"
	"strings"
)

// InstanceConfigurationLaunchInstancePlatformConfig The platform configuration requested for the instance.
// If you provide the parameter, the instance is created with the platform configuration that you specify.
// For any values that you omit, the instance uses the default configuration values for the `shape` that you
// specify. If you don't provide the parameter, the default values for the `shape` are used.
// Each shape only supports certain configurable values. If the values that you provide are not valid for the
// specified `shape`, an error is returned.
type InstanceConfigurationLaunchInstancePlatformConfig interface {

	// Whether Secure Boot is enabled on the instance.
	GetIsSecureBootEnabled() *bool

	// Whether the Trusted Platform Module (TPM) is enabled on the instance.
	GetIsTrustedPlatformModuleEnabled() *bool

	// Whether the Measured Boot feature is enabled on the instance.
	GetIsMeasuredBootEnabled() *bool
}

type instanceconfigurationlaunchinstanceplatformconfig struct {
	JsonData                       []byte
	IsSecureBootEnabled            *bool  `mandatory:"false" json:"isSecureBootEnabled"`
	IsTrustedPlatformModuleEnabled *bool  `mandatory:"false" json:"isTrustedPlatformModuleEnabled"`
	IsMeasuredBootEnabled          *bool  `mandatory:"false" json:"isMeasuredBootEnabled"`
	Type                           string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *instanceconfigurationlaunchinstanceplatformconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceconfigurationlaunchinstanceplatformconfig instanceconfigurationlaunchinstanceplatformconfig
	s := struct {
		Model Unmarshalerinstanceconfigurationlaunchinstanceplatformconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsSecureBootEnabled = s.Model.IsSecureBootEnabled
	m.IsTrustedPlatformModuleEnabled = s.Model.IsTrustedPlatformModuleEnabled
	m.IsMeasuredBootEnabled = s.Model.IsMeasuredBootEnabled
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceconfigurationlaunchinstanceplatformconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "AMD_MILAN_BM":
		mm := InstanceConfigurationAmdMilanBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEL_VM":
		mm := InstanceConfigurationIntelVmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMD_ROME_BM":
		mm := InstanceConfigurationAmdRomeBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEL_SKYLAKE_BM":
		mm := InstanceConfigurationIntelSkylakeBmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMD_VM":
		mm := InstanceConfigurationAmdVmLaunchInstancePlatformConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetIsSecureBootEnabled returns IsSecureBootEnabled
func (m instanceconfigurationlaunchinstanceplatformconfig) GetIsSecureBootEnabled() *bool {
	return m.IsSecureBootEnabled
}

//GetIsTrustedPlatformModuleEnabled returns IsTrustedPlatformModuleEnabled
func (m instanceconfigurationlaunchinstanceplatformconfig) GetIsTrustedPlatformModuleEnabled() *bool {
	return m.IsTrustedPlatformModuleEnabled
}

//GetIsMeasuredBootEnabled returns IsMeasuredBootEnabled
func (m instanceconfigurationlaunchinstanceplatformconfig) GetIsMeasuredBootEnabled() *bool {
	return m.IsMeasuredBootEnabled
}

func (m instanceconfigurationlaunchinstanceplatformconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m instanceconfigurationlaunchinstanceplatformconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceConfigurationLaunchInstancePlatformConfigTypeEnum Enum with underlying type: string
type InstanceConfigurationLaunchInstancePlatformConfigTypeEnum string

// Set of constants representing the allowable values for InstanceConfigurationLaunchInstancePlatformConfigTypeEnum
const (
	InstanceConfigurationLaunchInstancePlatformConfigTypeAmdMilanBm     InstanceConfigurationLaunchInstancePlatformConfigTypeEnum = "AMD_MILAN_BM"
	InstanceConfigurationLaunchInstancePlatformConfigTypeAmdRomeBm      InstanceConfigurationLaunchInstancePlatformConfigTypeEnum = "AMD_ROME_BM"
	InstanceConfigurationLaunchInstancePlatformConfigTypeIntelSkylakeBm InstanceConfigurationLaunchInstancePlatformConfigTypeEnum = "INTEL_SKYLAKE_BM"
	InstanceConfigurationLaunchInstancePlatformConfigTypeAmdVm          InstanceConfigurationLaunchInstancePlatformConfigTypeEnum = "AMD_VM"
	InstanceConfigurationLaunchInstancePlatformConfigTypeIntelVm        InstanceConfigurationLaunchInstancePlatformConfigTypeEnum = "INTEL_VM"
)

var mappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnum = map[string]InstanceConfigurationLaunchInstancePlatformConfigTypeEnum{
	"AMD_MILAN_BM":     InstanceConfigurationLaunchInstancePlatformConfigTypeAmdMilanBm,
	"AMD_ROME_BM":      InstanceConfigurationLaunchInstancePlatformConfigTypeAmdRomeBm,
	"INTEL_SKYLAKE_BM": InstanceConfigurationLaunchInstancePlatformConfigTypeIntelSkylakeBm,
	"AMD_VM":           InstanceConfigurationLaunchInstancePlatformConfigTypeAmdVm,
	"INTEL_VM":         InstanceConfigurationLaunchInstancePlatformConfigTypeIntelVm,
}

var mappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnumLowerCase = map[string]InstanceConfigurationLaunchInstancePlatformConfigTypeEnum{
	"amd_milan_bm":     InstanceConfigurationLaunchInstancePlatformConfigTypeAmdMilanBm,
	"amd_rome_bm":      InstanceConfigurationLaunchInstancePlatformConfigTypeAmdRomeBm,
	"intel_skylake_bm": InstanceConfigurationLaunchInstancePlatformConfigTypeIntelSkylakeBm,
	"amd_vm":           InstanceConfigurationLaunchInstancePlatformConfigTypeAmdVm,
	"intel_vm":         InstanceConfigurationLaunchInstancePlatformConfigTypeIntelVm,
}

// GetInstanceConfigurationLaunchInstancePlatformConfigTypeEnumValues Enumerates the set of values for InstanceConfigurationLaunchInstancePlatformConfigTypeEnum
func GetInstanceConfigurationLaunchInstancePlatformConfigTypeEnumValues() []InstanceConfigurationLaunchInstancePlatformConfigTypeEnum {
	values := make([]InstanceConfigurationLaunchInstancePlatformConfigTypeEnum, 0)
	for _, v := range mappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceConfigurationLaunchInstancePlatformConfigTypeEnumStringValues Enumerates the set of values in String for InstanceConfigurationLaunchInstancePlatformConfigTypeEnum
func GetInstanceConfigurationLaunchInstancePlatformConfigTypeEnumStringValues() []string {
	return []string{
		"AMD_MILAN_BM",
		"AMD_ROME_BM",
		"INTEL_SKYLAKE_BM",
		"AMD_VM",
		"INTEL_VM",
	}
}

// GetMappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnum(val string) (InstanceConfigurationLaunchInstancePlatformConfigTypeEnum, bool) {
	enum, ok := mappingInstanceConfigurationLaunchInstancePlatformConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}