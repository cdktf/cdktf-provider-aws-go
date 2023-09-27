// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateBlockDeviceMappingsEbs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#delete_on_termination LaunchTemplate#delete_on_termination}.
	DeleteOnTermination *string `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#encrypted LaunchTemplate#encrypted}.
	Encrypted *string `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#iops LaunchTemplate#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#kms_key_id LaunchTemplate#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#snapshot_id LaunchTemplate#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#throughput LaunchTemplate#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#volume_size LaunchTemplate#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/launch_template#volume_type LaunchTemplate#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

