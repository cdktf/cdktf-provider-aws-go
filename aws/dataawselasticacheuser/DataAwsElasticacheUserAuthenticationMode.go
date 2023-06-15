package dataawselasticacheuser


type DataAwsElasticacheUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/data-sources/elasticache_user#password_count DataAwsElasticacheUser#password_count}.
	PasswordCount *float64 `field:"optional" json:"passwordCount" yaml:"passwordCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.67.0/docs/data-sources/elasticache_user#type DataAwsElasticacheUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

