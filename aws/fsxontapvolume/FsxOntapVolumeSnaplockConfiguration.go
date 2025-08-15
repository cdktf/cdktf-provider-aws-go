// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume


type FsxOntapVolumeSnaplockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#snaplock_type FsxOntapVolume#snaplock_type}.
	SnaplockType *string `field:"required" json:"snaplockType" yaml:"snaplockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#audit_log_volume FsxOntapVolume#audit_log_volume}.
	AuditLogVolume interface{} `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// autocommit_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#autocommit_period FsxOntapVolume#autocommit_period}
	AutocommitPeriod *FsxOntapVolumeSnaplockConfigurationAutocommitPeriod `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#privileged_delete FsxOntapVolume#privileged_delete}.
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// retention_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#retention_period FsxOntapVolume#retention_period}
	RetentionPeriod *FsxOntapVolumeSnaplockConfigurationRetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/fsx_ontap_volume#volume_append_mode_enabled FsxOntapVolume#volume_append_mode_enabled}.
	VolumeAppendModeEnabled interface{} `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

