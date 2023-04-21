package mskserverlesscluster


type MskServerlessClusterClientAuthenticationSasl struct {
	// iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/msk_serverless_cluster#iam MskServerlessCluster#iam}
	Iam *MskServerlessClusterClientAuthenticationSaslIam `field:"required" json:"iam" yaml:"iam"`
}

