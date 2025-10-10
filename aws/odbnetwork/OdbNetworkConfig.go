// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbNetworkConfig struct {
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
	// The AZ ID of the AZ where the ODB network is located.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#availability_zone_id OdbNetwork#availability_zone_id}
	AvailabilityZoneId *string `field:"required" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The CIDR range of the backup subnet for the ODB network.
	//
	// Changing this will force terraform to create new resource.
	// 	Constraints:
	// 	   - Must not overlap with the CIDR range of the client subnet.
	// 	   - Must not overlap with the CIDR ranges of the VPCs that are connected to the
	// 	   ODB network.
	// 	   - Must not use the following CIDR ranges that are reserved by OCI:
	// 	   - 100.106.0.0/16 and 100.107.0.0/16
	// 	   - 169.254.0.0/16
	// 	   - 224.0.0.0 - 239.255.255.255
	// 	   - 240.0.0.0 - 255.255.255.255
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#backup_subnet_cidr OdbNetwork#backup_subnet_cidr}
	BackupSubnetCidr *string `field:"required" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// The CIDR notation for the network resource.
	//
	// Changing this will force terraform to create new resource.
	//  Constraints:
	//   	 - Must not overlap with the CIDR range of the backup subnet.
	//    	- Must not overlap with the CIDR ranges of the VPCs that are connected to the
	//    ODB network.
	//   	- Must not use the following CIDR ranges that are reserved by OCI:
	//   	 - 100.106.0.0/16 and 100.107.0.0/16
	//   	 - 169.254.0.0/16
	//    	- 224.0.0.0 - 239.255.255.255
	//    	- 240.0.0.0 - 255.255.255.255
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#client_subnet_cidr OdbNetwork#client_subnet_cidr}
	ClientSubnetCidr *string `field:"required" json:"clientSubnetCidr" yaml:"clientSubnetCidr"`
	// The user-friendly name for the odb network. Changing this will force terraform to create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#display_name OdbNetwork#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Specifies the configuration for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#s3_access OdbNetwork#s3_access}
	S3Access *string `field:"required" json:"s3Access" yaml:"s3Access"`
	// Specifies the configuration for Zero-ETL access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#zero_etl_access OdbNetwork#zero_etl_access}
	ZeroEtlAccess *string `field:"required" json:"zeroEtlAccess" yaml:"zeroEtlAccess"`
	// The name of the Availability Zone (AZ) where the odb network is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#availability_zone OdbNetwork#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The name of the custom domain that the network is located. custom_domain_name and default_dns_prefix both can't be given.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#custom_domain_name OdbNetwork#custom_domain_name}
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// The default DNS prefix for the network resource. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#default_dns_prefix OdbNetwork#default_dns_prefix}
	DefaultDnsPrefix *string `field:"optional" json:"defaultDnsPrefix" yaml:"defaultDnsPrefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#region OdbNetwork#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies the endpoint policy for Amazon S3 access from the ODB network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#s3_policy_document OdbNetwork#s3_policy_document}
	S3PolicyDocument *string `field:"optional" json:"s3PolicyDocument" yaml:"s3PolicyDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#tags OdbNetwork#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/odb_network#timeouts OdbNetwork#timeouts}
	Timeouts *OdbNetworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

