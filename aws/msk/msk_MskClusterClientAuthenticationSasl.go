package msk


type MskClusterClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#iam MskCluster#iam}.
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_cluster#scram MskCluster#scram}.
	Scram interface{} `field:"optional" json:"scram" yaml:"scram"`
}

