package ekscluster


type EksClusterEncryptionConfigProvider struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/eks_cluster#key_arn EksCluster#key_arn}.
	KeyArn *string `field:"required" json:"keyArn" yaml:"keyArn"`
}

