package provider


type AwsProviderAssumeRoleWithWebIdentity struct {
	// The duration, between 15 minutes and 12 hours, of the role session.
	//
	// Valid time units are ns, us (or µs), ms, s, h, or m.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#duration AwsProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#policy AwsProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#policy_arns AwsProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#role_arn AwsProvider#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An identifier for the assumed role session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#session_name AwsProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#web_identity_token AwsProvider#web_identity_token}.
	WebIdentityToken *string `field:"optional" json:"webIdentityToken" yaml:"webIdentityToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs#web_identity_token_file AwsProvider#web_identity_token_file}.
	WebIdentityTokenFile *string `field:"optional" json:"webIdentityTokenFile" yaml:"webIdentityTokenFile"`
}

