// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchdomain


type OpensearchDomainClusterConfig struct {
	// cold_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#cold_storage_options OpensearchDomain#cold_storage_options}
	ColdStorageOptions *OpensearchDomainClusterConfigColdStorageOptions `field:"optional" json:"coldStorageOptions" yaml:"coldStorageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#dedicated_master_count OpensearchDomain#dedicated_master_count}.
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#dedicated_master_enabled OpensearchDomain#dedicated_master_enabled}.
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#dedicated_master_type OpensearchDomain#dedicated_master_type}.
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#instance_count OpensearchDomain#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#instance_type OpensearchDomain#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#multi_az_with_standby_enabled OpensearchDomain#multi_az_with_standby_enabled}.
	MultiAzWithStandbyEnabled interface{} `field:"optional" json:"multiAzWithStandbyEnabled" yaml:"multiAzWithStandbyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#warm_count OpensearchDomain#warm_count}.
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#warm_enabled OpensearchDomain#warm_enabled}.
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#warm_type OpensearchDomain#warm_type}.
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// zone_awareness_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#zone_awareness_config OpensearchDomain#zone_awareness_config}
	ZoneAwarenessConfig *OpensearchDomainClusterConfigZoneAwarenessConfig `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/opensearch_domain#zone_awareness_enabled OpensearchDomain#zone_awareness_enabled}.
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}

