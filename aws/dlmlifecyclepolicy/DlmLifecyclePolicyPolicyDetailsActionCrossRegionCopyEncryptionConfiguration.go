package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dlm_lifecycle_policy#cmk_arn DlmLifecyclePolicy#cmk_arn}.
	CmkArn *string `field:"optional" json:"cmkArn" yaml:"cmkArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.1/docs/resources/dlm_lifecycle_policy#encrypted DlmLifecyclePolicy#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
}

