package iam


type DataAwsIamPolicyDocumentStatementPrincipals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/iam_policy_document#identifiers DataAwsIamPolicyDocument#identifiers}.
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/iam_policy_document#type DataAwsIamPolicyDocument#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

