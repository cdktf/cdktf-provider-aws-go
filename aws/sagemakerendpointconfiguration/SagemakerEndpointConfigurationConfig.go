package sagemakerendpointconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerEndpointConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// production_variants block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#production_variants SagemakerEndpointConfiguration#production_variants}
	ProductionVariants interface{} `field:"required" json:"productionVariants" yaml:"productionVariants"`
	// async_inference_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#async_inference_config SagemakerEndpointConfiguration#async_inference_config}
	AsyncInferenceConfig *SagemakerEndpointConfigurationAsyncInferenceConfig `field:"optional" json:"asyncInferenceConfig" yaml:"asyncInferenceConfig"`
	// data_capture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#data_capture_config SagemakerEndpointConfiguration#data_capture_config}
	DataCaptureConfig *SagemakerEndpointConfigurationDataCaptureConfig `field:"optional" json:"dataCaptureConfig" yaml:"dataCaptureConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#id SagemakerEndpointConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#kms_key_arn SagemakerEndpointConfiguration#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#name SagemakerEndpointConfiguration#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#name_prefix SagemakerEndpointConfiguration#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// shadow_production_variants block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#shadow_production_variants SagemakerEndpointConfiguration#shadow_production_variants}
	ShadowProductionVariants interface{} `field:"optional" json:"shadowProductionVariants" yaml:"shadowProductionVariants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#tags SagemakerEndpointConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/sagemaker_endpoint_configuration#tags_all SagemakerEndpointConfiguration#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

