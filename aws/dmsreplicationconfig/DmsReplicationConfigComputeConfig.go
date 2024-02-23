// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreplicationconfig


type DmsReplicationConfigComputeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#replication_subnet_group_id DmsReplicationConfig#replication_subnet_group_id}.
	ReplicationSubnetGroupId *string `field:"required" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#availability_zone DmsReplicationConfig#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#dns_name_servers DmsReplicationConfig#dns_name_servers}.
	DnsNameServers *string `field:"optional" json:"dnsNameServers" yaml:"dnsNameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#kms_key_id DmsReplicationConfig#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#max_capacity_units DmsReplicationConfig#max_capacity_units}.
	MaxCapacityUnits *float64 `field:"optional" json:"maxCapacityUnits" yaml:"maxCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#min_capacity_units DmsReplicationConfig#min_capacity_units}.
	MinCapacityUnits *float64 `field:"optional" json:"minCapacityUnits" yaml:"minCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#multi_az DmsReplicationConfig#multi_az}.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#preferred_maintenance_window DmsReplicationConfig#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/dms_replication_config#vpc_security_group_ids DmsReplicationConfig#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

