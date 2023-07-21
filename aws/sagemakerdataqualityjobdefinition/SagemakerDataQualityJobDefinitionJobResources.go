package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionJobResources struct {
	// cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.9.0/docs/resources/sagemaker_data_quality_job_definition#cluster_config SagemakerDataQualityJobDefinition#cluster_config}
	ClusterConfig *SagemakerDataQualityJobDefinitionJobResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

