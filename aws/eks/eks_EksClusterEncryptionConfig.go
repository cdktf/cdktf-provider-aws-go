package eks


type EksClusterEncryptionConfig struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_cluster#provider EksCluster#provider}
	Provider *EksClusterEncryptionConfigProvider `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_cluster#resources EksCluster#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

