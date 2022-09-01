package eks


type EksNodeGroupRemoteAccess struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#ec2_ssh_key EksNodeGroup#ec2_ssh_key}.
	Ec2SshKey *string `field:"optional" json:"ec2SshKey" yaml:"ec2SshKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_node_group#source_security_group_ids EksNodeGroup#source_security_group_ids}.
	SourceSecurityGroupIds *[]*string `field:"optional" json:"sourceSecurityGroupIds" yaml:"sourceSecurityGroupIds"`
}

