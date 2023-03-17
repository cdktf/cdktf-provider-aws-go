package dataawselasticacheuser


type DataAwsElasticacheUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticache_user#password_count DataAwsElasticacheUser#password_count}.
	PasswordCount *float64 `field:"optional" json:"passwordCount" yaml:"passwordCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/d/elasticache_user#type DataAwsElasticacheUser#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

