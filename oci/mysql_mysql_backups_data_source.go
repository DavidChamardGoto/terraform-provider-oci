// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v25/mysql"
)

func init() {
	RegisterDatasource("oci_mysql_mysql_backups", MysqlMysqlBackupsDataSource())
}

func MysqlMysqlBackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlMysqlBackups,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(MysqlMysqlBackupResource()),
			},
		},
	}
}

func readMysqlMysqlBackups(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlMysqlBackupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dbBackupsClient()

	return ReadResource(sync)
}

type MysqlMysqlBackupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.DbBackupsClient
	Res    *oci_mysql.ListBackupsResponse
}

func (s *MysqlMysqlBackupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlMysqlBackupsDataSourceCrud) Get() error {
	request := oci_mysql.ListBackupsRequest{}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_mysql.BackupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "mysql")

	response, err := s.Client.ListBackups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBackups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MysqlMysqlBackupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		mysqlBackup := map[string]interface{}{}

		if r.BackupSizeInGBs != nil {
			mysqlBackup["backup_size_in_gbs"] = *r.BackupSizeInGBs
		}

		mysqlBackup["backup_type"] = r.BackupType

		if r.DataStorageSizeInGBs != nil {
			mysqlBackup["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DbSystemId != nil {
			mysqlBackup["db_system_id"] = *r.DbSystemId
		}

		if r.DefinedTags != nil {
			mysqlBackup["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			mysqlBackup["description"] = *r.Description
		}

		if r.DisplayName != nil {
			mysqlBackup["display_name"] = *r.DisplayName
		}

		mysqlBackup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			mysqlBackup["id"] = *r.Id
		}

		if r.MysqlVersion != nil {
			mysqlBackup["mysql_version"] = *r.MysqlVersion
		}

		if r.RetentionInDays != nil {
			mysqlBackup["retention_in_days"] = *r.RetentionInDays
		}

		if r.ShapeName != nil {
			mysqlBackup["shape_name"] = *r.ShapeName
		}

		mysqlBackup["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			mysqlBackup["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, mysqlBackup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MysqlMysqlBackupsDataSource().Schema["backups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backups", resources); err != nil {
		return err
	}

	return nil
}