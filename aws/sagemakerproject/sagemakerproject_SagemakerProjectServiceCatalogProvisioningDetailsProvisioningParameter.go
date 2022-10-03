package sagemakerproject


type SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_project#key SagemakerProject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_project#value SagemakerProject#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

