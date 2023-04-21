package mskcluster


type MskClusterClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/msk_cluster#iam MskCluster#iam}.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/msk_cluster#scram MskCluster#scram}.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

