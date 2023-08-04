package securityhubinsight


type SecurityhubInsightFilters struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#aws_account_id SecurityhubInsight#aws_account_id}
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// company_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#company_name SecurityhubInsight#company_name}
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// compliance_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#compliance_status SecurityhubInsight#compliance_status}
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// confidence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#confidence SecurityhubInsight#confidence}
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// created_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#created_at SecurityhubInsight#created_at}
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#criticality SecurityhubInsight#criticality}
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// description block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#description SecurityhubInsight#description}
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// finding_provider_fields_confidence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_confidence SecurityhubInsight#finding_provider_fields_confidence}
	FindingProviderFieldsConfidence interface{} `field:"optional" json:"findingProviderFieldsConfidence" yaml:"findingProviderFieldsConfidence"`
	// finding_provider_fields_criticality block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_criticality SecurityhubInsight#finding_provider_fields_criticality}
	FindingProviderFieldsCriticality interface{} `field:"optional" json:"findingProviderFieldsCriticality" yaml:"findingProviderFieldsCriticality"`
	// finding_provider_fields_related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_related_findings_id SecurityhubInsight#finding_provider_fields_related_findings_id}
	FindingProviderFieldsRelatedFindingsId interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsId" yaml:"findingProviderFieldsRelatedFindingsId"`
	// finding_provider_fields_related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_related_findings_product_arn SecurityhubInsight#finding_provider_fields_related_findings_product_arn}
	FindingProviderFieldsRelatedFindingsProductArn interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsProductArn" yaml:"findingProviderFieldsRelatedFindingsProductArn"`
	// finding_provider_fields_severity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_severity_label SecurityhubInsight#finding_provider_fields_severity_label}
	FindingProviderFieldsSeverityLabel interface{} `field:"optional" json:"findingProviderFieldsSeverityLabel" yaml:"findingProviderFieldsSeverityLabel"`
	// finding_provider_fields_severity_original block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_severity_original SecurityhubInsight#finding_provider_fields_severity_original}
	FindingProviderFieldsSeverityOriginal interface{} `field:"optional" json:"findingProviderFieldsSeverityOriginal" yaml:"findingProviderFieldsSeverityOriginal"`
	// finding_provider_fields_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#finding_provider_fields_types SecurityhubInsight#finding_provider_fields_types}
	FindingProviderFieldsTypes interface{} `field:"optional" json:"findingProviderFieldsTypes" yaml:"findingProviderFieldsTypes"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#first_observed_at SecurityhubInsight#first_observed_at}
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// generator_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#generator_id SecurityhubInsight#generator_id}
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#id SecurityhubInsight#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// keyword block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#keyword SecurityhubInsight#keyword}
	Keyword interface{} `field:"optional" json:"keyword" yaml:"keyword"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#last_observed_at SecurityhubInsight#last_observed_at}
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// malware_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#malware_name SecurityhubInsight#malware_name}
	MalwareName interface{} `field:"optional" json:"malwareName" yaml:"malwareName"`
	// malware_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#malware_path SecurityhubInsight#malware_path}
	MalwarePath interface{} `field:"optional" json:"malwarePath" yaml:"malwarePath"`
	// malware_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#malware_state SecurityhubInsight#malware_state}
	MalwareState interface{} `field:"optional" json:"malwareState" yaml:"malwareState"`
	// malware_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#malware_type SecurityhubInsight#malware_type}
	MalwareType interface{} `field:"optional" json:"malwareType" yaml:"malwareType"`
	// network_destination_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_destination_domain SecurityhubInsight#network_destination_domain}
	NetworkDestinationDomain interface{} `field:"optional" json:"networkDestinationDomain" yaml:"networkDestinationDomain"`
	// network_destination_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_destination_ipv4 SecurityhubInsight#network_destination_ipv4}
	NetworkDestinationIpv4 interface{} `field:"optional" json:"networkDestinationIpv4" yaml:"networkDestinationIpv4"`
	// network_destination_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_destination_ipv6 SecurityhubInsight#network_destination_ipv6}
	NetworkDestinationIpv6 interface{} `field:"optional" json:"networkDestinationIpv6" yaml:"networkDestinationIpv6"`
	// network_destination_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_destination_port SecurityhubInsight#network_destination_port}
	NetworkDestinationPort interface{} `field:"optional" json:"networkDestinationPort" yaml:"networkDestinationPort"`
	// network_direction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_direction SecurityhubInsight#network_direction}
	NetworkDirection interface{} `field:"optional" json:"networkDirection" yaml:"networkDirection"`
	// network_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_protocol SecurityhubInsight#network_protocol}
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// network_source_domain block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_source_domain SecurityhubInsight#network_source_domain}
	NetworkSourceDomain interface{} `field:"optional" json:"networkSourceDomain" yaml:"networkSourceDomain"`
	// network_source_ipv4 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_source_ipv4 SecurityhubInsight#network_source_ipv4}
	NetworkSourceIpv4 interface{} `field:"optional" json:"networkSourceIpv4" yaml:"networkSourceIpv4"`
	// network_source_ipv6 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_source_ipv6 SecurityhubInsight#network_source_ipv6}
	NetworkSourceIpv6 interface{} `field:"optional" json:"networkSourceIpv6" yaml:"networkSourceIpv6"`
	// network_source_mac block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_source_mac SecurityhubInsight#network_source_mac}
	NetworkSourceMac interface{} `field:"optional" json:"networkSourceMac" yaml:"networkSourceMac"`
	// network_source_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#network_source_port SecurityhubInsight#network_source_port}
	NetworkSourcePort interface{} `field:"optional" json:"networkSourcePort" yaml:"networkSourcePort"`
	// note_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#note_text SecurityhubInsight#note_text}
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// note_updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#note_updated_at SecurityhubInsight#note_updated_at}
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// note_updated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#note_updated_by SecurityhubInsight#note_updated_by}
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// process_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_launched_at SecurityhubInsight#process_launched_at}
	ProcessLaunchedAt interface{} `field:"optional" json:"processLaunchedAt" yaml:"processLaunchedAt"`
	// process_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_name SecurityhubInsight#process_name}
	ProcessName interface{} `field:"optional" json:"processName" yaml:"processName"`
	// process_parent_pid block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_parent_pid SecurityhubInsight#process_parent_pid}
	ProcessParentPid interface{} `field:"optional" json:"processParentPid" yaml:"processParentPid"`
	// process_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_path SecurityhubInsight#process_path}
	ProcessPath interface{} `field:"optional" json:"processPath" yaml:"processPath"`
	// process_pid block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_pid SecurityhubInsight#process_pid}
	ProcessPid interface{} `field:"optional" json:"processPid" yaml:"processPid"`
	// process_terminated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#process_terminated_at SecurityhubInsight#process_terminated_at}
	ProcessTerminatedAt interface{} `field:"optional" json:"processTerminatedAt" yaml:"processTerminatedAt"`
	// product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#product_arn SecurityhubInsight#product_arn}
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// product_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#product_fields SecurityhubInsight#product_fields}
	ProductFields interface{} `field:"optional" json:"productFields" yaml:"productFields"`
	// product_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#product_name SecurityhubInsight#product_name}
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// recommendation_text block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#recommendation_text SecurityhubInsight#recommendation_text}
	RecommendationText interface{} `field:"optional" json:"recommendationText" yaml:"recommendationText"`
	// record_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#record_state SecurityhubInsight#record_state}
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// related_findings_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#related_findings_id SecurityhubInsight#related_findings_id}
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// related_findings_product_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#related_findings_product_arn SecurityhubInsight#related_findings_product_arn}
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// resource_aws_ec2_instance_iam_instance_profile_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_iam_instance_profile_arn SecurityhubInsight#resource_aws_ec2_instance_iam_instance_profile_arn}
	ResourceAwsEc2InstanceIamInstanceProfileArn interface{} `field:"optional" json:"resourceAwsEc2InstanceIamInstanceProfileArn" yaml:"resourceAwsEc2InstanceIamInstanceProfileArn"`
	// resource_aws_ec2_instance_image_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_image_id SecurityhubInsight#resource_aws_ec2_instance_image_id}
	ResourceAwsEc2InstanceImageId interface{} `field:"optional" json:"resourceAwsEc2InstanceImageId" yaml:"resourceAwsEc2InstanceImageId"`
	// resource_aws_ec2_instance_ipv4_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_ipv4_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv4_addresses}
	ResourceAwsEc2InstanceIpv4Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpv4Addresses" yaml:"resourceAwsEc2InstanceIpv4Addresses"`
	// resource_aws_ec2_instance_ipv6_addresses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_ipv6_addresses SecurityhubInsight#resource_aws_ec2_instance_ipv6_addresses}
	ResourceAwsEc2InstanceIpv6Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpv6Addresses" yaml:"resourceAwsEc2InstanceIpv6Addresses"`
	// resource_aws_ec2_instance_key_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_key_name SecurityhubInsight#resource_aws_ec2_instance_key_name}
	ResourceAwsEc2InstanceKeyName interface{} `field:"optional" json:"resourceAwsEc2InstanceKeyName" yaml:"resourceAwsEc2InstanceKeyName"`
	// resource_aws_ec2_instance_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_launched_at SecurityhubInsight#resource_aws_ec2_instance_launched_at}
	ResourceAwsEc2InstanceLaunchedAt interface{} `field:"optional" json:"resourceAwsEc2InstanceLaunchedAt" yaml:"resourceAwsEc2InstanceLaunchedAt"`
	// resource_aws_ec2_instance_subnet_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_subnet_id SecurityhubInsight#resource_aws_ec2_instance_subnet_id}
	ResourceAwsEc2InstanceSubnetId interface{} `field:"optional" json:"resourceAwsEc2InstanceSubnetId" yaml:"resourceAwsEc2InstanceSubnetId"`
	// resource_aws_ec2_instance_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_type SecurityhubInsight#resource_aws_ec2_instance_type}
	ResourceAwsEc2InstanceType interface{} `field:"optional" json:"resourceAwsEc2InstanceType" yaml:"resourceAwsEc2InstanceType"`
	// resource_aws_ec2_instance_vpc_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_ec2_instance_vpc_id SecurityhubInsight#resource_aws_ec2_instance_vpc_id}
	ResourceAwsEc2InstanceVpcId interface{} `field:"optional" json:"resourceAwsEc2InstanceVpcId" yaml:"resourceAwsEc2InstanceVpcId"`
	// resource_aws_iam_access_key_created_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_created_at SecurityhubInsight#resource_aws_iam_access_key_created_at}
	ResourceAwsIamAccessKeyCreatedAt interface{} `field:"optional" json:"resourceAwsIamAccessKeyCreatedAt" yaml:"resourceAwsIamAccessKeyCreatedAt"`
	// resource_aws_iam_access_key_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_status SecurityhubInsight#resource_aws_iam_access_key_status}
	ResourceAwsIamAccessKeyStatus interface{} `field:"optional" json:"resourceAwsIamAccessKeyStatus" yaml:"resourceAwsIamAccessKeyStatus"`
	// resource_aws_iam_access_key_user_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_user_name SecurityhubInsight#resource_aws_iam_access_key_user_name}
	ResourceAwsIamAccessKeyUserName interface{} `field:"optional" json:"resourceAwsIamAccessKeyUserName" yaml:"resourceAwsIamAccessKeyUserName"`
	// resource_aws_s3_bucket_owner_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_s3_bucket_owner_id SecurityhubInsight#resource_aws_s3_bucket_owner_id}
	ResourceAwsS3BucketOwnerId interface{} `field:"optional" json:"resourceAwsS3BucketOwnerId" yaml:"resourceAwsS3BucketOwnerId"`
	// resource_aws_s3_bucket_owner_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_aws_s3_bucket_owner_name SecurityhubInsight#resource_aws_s3_bucket_owner_name}
	ResourceAwsS3BucketOwnerName interface{} `field:"optional" json:"resourceAwsS3BucketOwnerName" yaml:"resourceAwsS3BucketOwnerName"`
	// resource_container_image_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_container_image_id SecurityhubInsight#resource_container_image_id}
	ResourceContainerImageId interface{} `field:"optional" json:"resourceContainerImageId" yaml:"resourceContainerImageId"`
	// resource_container_image_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_container_image_name SecurityhubInsight#resource_container_image_name}
	ResourceContainerImageName interface{} `field:"optional" json:"resourceContainerImageName" yaml:"resourceContainerImageName"`
	// resource_container_launched_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_container_launched_at SecurityhubInsight#resource_container_launched_at}
	ResourceContainerLaunchedAt interface{} `field:"optional" json:"resourceContainerLaunchedAt" yaml:"resourceContainerLaunchedAt"`
	// resource_container_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_container_name SecurityhubInsight#resource_container_name}
	ResourceContainerName interface{} `field:"optional" json:"resourceContainerName" yaml:"resourceContainerName"`
	// resource_details_other block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_details_other SecurityhubInsight#resource_details_other}
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_id SecurityhubInsight#resource_id}
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resource_partition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_partition SecurityhubInsight#resource_partition}
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// resource_region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_region SecurityhubInsight#resource_region}
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_tags SecurityhubInsight#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#resource_type SecurityhubInsight#resource_type}
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// severity_label block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#severity_label SecurityhubInsight#severity_label}
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// source_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#source_url SecurityhubInsight#source_url}
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// threat_intel_indicator_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_category SecurityhubInsight#threat_intel_indicator_category}
	ThreatIntelIndicatorCategory interface{} `field:"optional" json:"threatIntelIndicatorCategory" yaml:"threatIntelIndicatorCategory"`
	// threat_intel_indicator_last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_last_observed_at SecurityhubInsight#threat_intel_indicator_last_observed_at}
	ThreatIntelIndicatorLastObservedAt interface{} `field:"optional" json:"threatIntelIndicatorLastObservedAt" yaml:"threatIntelIndicatorLastObservedAt"`
	// threat_intel_indicator_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_source SecurityhubInsight#threat_intel_indicator_source}
	ThreatIntelIndicatorSource interface{} `field:"optional" json:"threatIntelIndicatorSource" yaml:"threatIntelIndicatorSource"`
	// threat_intel_indicator_source_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_source_url SecurityhubInsight#threat_intel_indicator_source_url}
	ThreatIntelIndicatorSourceUrl interface{} `field:"optional" json:"threatIntelIndicatorSourceUrl" yaml:"threatIntelIndicatorSourceUrl"`
	// threat_intel_indicator_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_type SecurityhubInsight#threat_intel_indicator_type}
	ThreatIntelIndicatorType interface{} `field:"optional" json:"threatIntelIndicatorType" yaml:"threatIntelIndicatorType"`
	// threat_intel_indicator_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#threat_intel_indicator_value SecurityhubInsight#threat_intel_indicator_value}
	ThreatIntelIndicatorValue interface{} `field:"optional" json:"threatIntelIndicatorValue" yaml:"threatIntelIndicatorValue"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#title SecurityhubInsight#title}
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#type SecurityhubInsight#type}
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#updated_at SecurityhubInsight#updated_at}
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// user_defined_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#user_defined_values SecurityhubInsight#user_defined_values}
	UserDefinedValues interface{} `field:"optional" json:"userDefinedValues" yaml:"userDefinedValues"`
	// verification_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#verification_state SecurityhubInsight#verification_state}
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// workflow_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.11.0/docs/resources/securityhub_insight#workflow_status SecurityhubInsight#workflow_status}
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

