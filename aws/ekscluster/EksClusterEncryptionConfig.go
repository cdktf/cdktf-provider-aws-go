package ekscluster


type EksClusterEncryptionConfig struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/eks_cluster#provider EksCluster#provider}
	Provider *EksClusterEncryptionConfigProvider `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/eks_cluster#resources EksCluster#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

