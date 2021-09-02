// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// UpdateDbManagementPrivateEndpointDetails The details used to update a Database Management private endpoint.
type UpdateDbManagementPrivateEndpointDetails struct {

	// The display name of the private endpoint.
	Name *string `mandatory:"false" json:"name"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m UpdateDbManagementPrivateEndpointDetails) String() string {
	return common.PointerString(m)
}
