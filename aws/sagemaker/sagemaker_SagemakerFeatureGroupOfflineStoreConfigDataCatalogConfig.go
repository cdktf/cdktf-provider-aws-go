package sagemaker


type SagemakerFeatureGroupOfflineStoreConfigDataCatalogConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_feature_group#catalog SagemakerFeatureGroup#catalog}.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_feature_group#database SagemakerFeatureGroup#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/sagemaker_feature_group#table_name SagemakerFeatureGroup#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}
