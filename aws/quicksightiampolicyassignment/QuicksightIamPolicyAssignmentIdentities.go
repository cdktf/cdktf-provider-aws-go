package quicksightiampolicyassignment


type QuicksightIamPolicyAssignmentIdentities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_iam_policy_assignment#group QuicksightIamPolicyAssignment#group}.
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/quicksight_iam_policy_assignment#user QuicksightIamPolicyAssignment#user}.
	User *[]*string `field:"optional" json:"user" yaml:"user"`
}

