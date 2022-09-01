package msk


type MskServerlessClusterClientAuthentication struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/msk_serverless_cluster#sasl MskServerlessCluster#sasl}
	Sasl *MskServerlessClusterClientAuthenticationSasl `field:"required" json:"sasl" yaml:"sasl"`
}

