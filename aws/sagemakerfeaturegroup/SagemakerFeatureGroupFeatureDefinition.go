package sagemakerfeaturegroup


type SagemakerFeatureGroupFeatureDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/sagemaker_feature_group#feature_name SagemakerFeatureGroup#feature_name}.
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/sagemaker_feature_group#feature_type SagemakerFeatureGroup#feature_type}.
	FeatureType *string `field:"optional" json:"featureType" yaml:"featureType"`
}

