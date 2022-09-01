package msk


type MskServerlessClusterClientAuthenticationSasl struct {
	// iam block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_serverless_cluster#iam MskServerlessCluster#iam}
	Iam *MskServerlessClusterClientAuthenticationSaslIam `field:"required" json:"iam" yaml:"iam"`
}

