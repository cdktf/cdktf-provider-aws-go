// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eksnodegroup


type EksNodeGroupRemoteAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/eks_node_group#ec2_ssh_key EksNodeGroup#ec2_ssh_key}.
	Ec2SshKey *string `field:"optional" json:"ec2SshKey" yaml:"ec2SshKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/eks_node_group#source_security_group_ids EksNodeGroup#source_security_group_ids}.
	SourceSecurityGroupIds *[]*string `field:"optional" json:"sourceSecurityGroupIds" yaml:"sourceSecurityGroupIds"`
}

