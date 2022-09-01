package dynamodb


type DynamodbTableTtl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dynamodb_table#attribute_name DynamodbTable#attribute_name}.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dynamodb_table#enabled DynamodbTable#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

