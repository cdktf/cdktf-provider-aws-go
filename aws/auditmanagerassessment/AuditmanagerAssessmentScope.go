package auditmanagerassessment


type AuditmanagerAssessmentScope struct {
	// aws_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/auditmanager_assessment#aws_accounts AuditmanagerAssessment#aws_accounts}
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// aws_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/auditmanager_assessment#aws_services AuditmanagerAssessment#aws_services}
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}

