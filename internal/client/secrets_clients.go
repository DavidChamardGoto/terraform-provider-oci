// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_secrets "github.com/oracle/oci-go-sdk/v62/secrets"

	oci_common "github.com/oracle/oci-go-sdk/v62/common"
)

func init() {
	RegisterOracleClient("oci_secrets.SecretsClient", &OracleClient{InitClientFn: initSecretsSecretsClient})
}

func initSecretsSecretsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_secrets.NewSecretsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) SecretsClient() *oci_secrets.SecretsClient {
	return m.GetClient("oci_secrets.SecretsClient").(*oci_secrets.SecretsClient)
}