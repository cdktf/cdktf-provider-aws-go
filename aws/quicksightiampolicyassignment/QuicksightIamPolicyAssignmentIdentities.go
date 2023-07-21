package quicksightiampolicyassignment


type QuicksightIamPolicyAssignmentIdentities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_iam_policy_assignment#group QuicksightIamPolicyAssignment#group}.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/quicksight_iam_policy_assignment#user QuicksightIamPolicyAssignment#user}.
	User *[]*string `field:"optional" json:"user" yaml:"user"`
}

