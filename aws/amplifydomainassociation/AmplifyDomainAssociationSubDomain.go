package amplifydomainassociation


type AmplifyDomainAssociationSubDomain struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/amplify_domain_association#branch_name AmplifyDomainAssociation#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.6.1/docs/resources/amplify_domain_association#prefix AmplifyDomainAssociation#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

