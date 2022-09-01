// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#document_type AppflowFlow#document_type}.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#include_all_versions AppflowFlow#include_all_versions}.
	IncludeAllVersions interface{} `field:"optional" json:"includeAllVersions" yaml:"includeAllVersions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#include_renditions AppflowFlow#include_renditions}.
	IncludeRenditions interface{} `field:"optional" json:"includeRenditions" yaml:"includeRenditions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appflow_flow#include_source_files AppflowFlow#include_source_files}.
	IncludeSourceFiles interface{} `field:"optional" json:"includeSourceFiles" yaml:"includeSourceFiles"`
}

