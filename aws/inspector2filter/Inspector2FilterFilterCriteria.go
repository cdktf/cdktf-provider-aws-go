// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package inspector2filter


type Inspector2FilterFilterCriteria struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#aws_account_id Inspector2Filter#aws_account_id}
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// code_vulnerability_detector_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#code_vulnerability_detector_name Inspector2Filter#code_vulnerability_detector_name}
	CodeVulnerabilityDetectorName interface{} `field:"optional" json:"codeVulnerabilityDetectorName" yaml:"codeVulnerabilityDetectorName"`
	// code_vulnerability_detector_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#code_vulnerability_detector_tags Inspector2Filter#code_vulnerability_detector_tags}
	CodeVulnerabilityDetectorTags interface{} `field:"optional" json:"codeVulnerabilityDetectorTags" yaml:"codeVulnerabilityDetectorTags"`
	// code_vulnerability_file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#code_vulnerability_file_path Inspector2Filter#code_vulnerability_file_path}
	CodeVulnerabilityFilePath interface{} `field:"optional" json:"codeVulnerabilityFilePath" yaml:"codeVulnerabilityFilePath"`
	// component_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#component_id Inspector2Filter#component_id}
	ComponentId interface{} `field:"optional" json:"componentId" yaml:"componentId"`
	// component_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#component_type Inspector2Filter#component_type}
	ComponentType interface{} `field:"optional" json:"componentType" yaml:"componentType"`
	// ec2_instance_image_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ec2_instance_image_id Inspector2Filter#ec2_instance_image_id}
	Ec2InstanceImageId interface{} `field:"optional" json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// ec2_instance_subnet_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ec2_instance_subnet_id Inspector2Filter#ec2_instance_subnet_id}
	Ec2InstanceSubnetId interface{} `field:"optional" json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// ec2_instance_vpc_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ec2_instance_vpc_id Inspector2Filter#ec2_instance_vpc_id}
	Ec2InstanceVpcId interface{} `field:"optional" json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// ecr_image_architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_architecture Inspector2Filter#ecr_image_architecture}
	EcrImageArchitecture interface{} `field:"optional" json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// ecr_image_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_hash Inspector2Filter#ecr_image_hash}
	EcrImageHash interface{} `field:"optional" json:"ecrImageHash" yaml:"ecrImageHash"`
	// ecr_image_pushed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_pushed_at Inspector2Filter#ecr_image_pushed_at}
	EcrImagePushedAt interface{} `field:"optional" json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// ecr_image_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_registry Inspector2Filter#ecr_image_registry}
	EcrImageRegistry interface{} `field:"optional" json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// ecr_image_repository_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_repository_name Inspector2Filter#ecr_image_repository_name}
	EcrImageRepositoryName interface{} `field:"optional" json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// ecr_image_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#ecr_image_tags Inspector2Filter#ecr_image_tags}
	EcrImageTags interface{} `field:"optional" json:"ecrImageTags" yaml:"ecrImageTags"`
	// epss_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#epss_score Inspector2Filter#epss_score}
	EpssScore interface{} `field:"optional" json:"epssScore" yaml:"epssScore"`
	// exploit_available block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#exploit_available Inspector2Filter#exploit_available}
	ExploitAvailable interface{} `field:"optional" json:"exploitAvailable" yaml:"exploitAvailable"`
	// finding_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#finding_arn Inspector2Filter#finding_arn}
	FindingArn interface{} `field:"optional" json:"findingArn" yaml:"findingArn"`
	// finding_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#finding_status Inspector2Filter#finding_status}
	FindingStatus interface{} `field:"optional" json:"findingStatus" yaml:"findingStatus"`
	// finding_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#finding_type Inspector2Filter#finding_type}
	FindingType interface{} `field:"optional" json:"findingType" yaml:"findingType"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#first_observed_at Inspector2Filter#first_observed_at}
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// fix_available block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#fix_available Inspector2Filter#fix_available}
	FixAvailable interface{} `field:"optional" json:"fixAvailable" yaml:"fixAvailable"`
	// inspector_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#inspector_score Inspector2Filter#inspector_score}
	InspectorScore interface{} `field:"optional" json:"inspectorScore" yaml:"inspectorScore"`
	// lambda_function_execution_role_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#lambda_function_execution_role_arn Inspector2Filter#lambda_function_execution_role_arn}
	LambdaFunctionExecutionRoleArn interface{} `field:"optional" json:"lambdaFunctionExecutionRoleArn" yaml:"lambdaFunctionExecutionRoleArn"`
	// lambda_function_last_modified_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#lambda_function_last_modified_at Inspector2Filter#lambda_function_last_modified_at}
	LambdaFunctionLastModifiedAt interface{} `field:"optional" json:"lambdaFunctionLastModifiedAt" yaml:"lambdaFunctionLastModifiedAt"`
	// lambda_function_layers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#lambda_function_layers Inspector2Filter#lambda_function_layers}
	LambdaFunctionLayers interface{} `field:"optional" json:"lambdaFunctionLayers" yaml:"lambdaFunctionLayers"`
	// lambda_function_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#lambda_function_name Inspector2Filter#lambda_function_name}
	LambdaFunctionName interface{} `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// lambda_function_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#lambda_function_runtime Inspector2Filter#lambda_function_runtime}
	LambdaFunctionRuntime interface{} `field:"optional" json:"lambdaFunctionRuntime" yaml:"lambdaFunctionRuntime"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#last_observed_at Inspector2Filter#last_observed_at}
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// network_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#network_protocol Inspector2Filter#network_protocol}
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#port_range Inspector2Filter#port_range}
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// related_vulnerabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#related_vulnerabilities Inspector2Filter#related_vulnerabilities}
	RelatedVulnerabilities interface{} `field:"optional" json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#resource_id Inspector2Filter#resource_id}
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#resource_tags Inspector2Filter#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#resource_type Inspector2Filter#resource_type}
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#severity Inspector2Filter#severity}
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#title Inspector2Filter#title}
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#updated_at Inspector2Filter#updated_at}
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// vendor_severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#vendor_severity Inspector2Filter#vendor_severity}
	VendorSeverity interface{} `field:"optional" json:"vendorSeverity" yaml:"vendorSeverity"`
	// vulnerability_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#vulnerability_id Inspector2Filter#vulnerability_id}
	VulnerabilityId interface{} `field:"optional" json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// vulnerability_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#vulnerability_source Inspector2Filter#vulnerability_source}
	VulnerabilitySource interface{} `field:"optional" json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// vulnerable_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/inspector2_filter#vulnerable_packages Inspector2Filter#vulnerable_packages}
	VulnerablePackages interface{} `field:"optional" json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

