package dynamodbglobaltable


type DynamodbGlobalTableReplica struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/dynamodb_global_table#region_name DynamodbGlobalTable#region_name}.
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
}

