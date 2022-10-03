package provider


type AwsProviderAssumeRole struct {
	// The duration, between 15 minutes and 12 hours, of the role session.
	//
	// Valid time units are ns, us (or µs), ms, s, h, or m.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#duration AwsProvider#duration}
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// The duration, in seconds, of the role session.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#duration_seconds AwsProvider#duration_seconds}
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// A unique identifier that might be required when you assume a role in another account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#external_id AwsProvider#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// IAM Policy JSON describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#policy AwsProvider#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// Amazon Resource Names (ARNs) of IAM Policies describing further restricting permissions for the IAM Role being assumed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#policy_arns AwsProvider#policy_arns}
	PolicyArns *[]*string `field:"optional" json:"policyArns" yaml:"policyArns"`
	// Amazon Resource Name (ARN) of an IAM Role to assume prior to making API calls.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#role_arn AwsProvider#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An identifier for the assumed role session.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#session_name AwsProvider#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
	// Source identity specified by the principal assuming the role.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#source_identity AwsProvider#source_identity}
	SourceIdentity *string `field:"optional" json:"sourceIdentity" yaml:"sourceIdentity"`
	// Assume role session tags.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#tags AwsProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Assume role session tag keys to pass to any subsequent sessions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws#transitive_tag_keys AwsProvider#transitive_tag_keys}
	TransitiveTagKeys *[]*string `field:"optional" json:"transitiveTagKeys" yaml:"transitiveTagKeys"`
}

