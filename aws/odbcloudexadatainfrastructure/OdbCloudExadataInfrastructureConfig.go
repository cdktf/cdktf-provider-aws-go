// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudexadatainfrastructure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbCloudExadataInfrastructureConfig struct {
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
	// The AZ ID of the AZ where the Exadata infrastructure is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#availability_zone_id OdbCloudExadataInfrastructure#availability_zone_id}
	AvailabilityZoneId *string `field:"required" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The user-friendly name for the Exadata infrastructure. Changing this will force terraform to create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#display_name OdbCloudExadataInfrastructure#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The model name of the Exadata infrastructure. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#shape OdbCloudExadataInfrastructure#shape}
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// The name of the Availability Zone (AZ) where the Exadata infrastructure is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#availability_zone OdbCloudExadataInfrastructure#availability_zone}
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of compute instances that the Exadata infrastructure is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#compute_count OdbCloudExadataInfrastructure#compute_count}
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#customer_contacts_to_send_to_oci OdbCloudExadataInfrastructure#customer_contacts_to_send_to_oci}
	CustomerContactsToSendToOci interface{} `field:"optional" json:"customerContactsToSendToOci" yaml:"customerContactsToSendToOci"`
	// The database server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the ListDbSystemShapes operation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#database_server_type OdbCloudExadataInfrastructure#database_server_type}
	DatabaseServerType *string `field:"optional" json:"databaseServerType" yaml:"databaseServerType"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#maintenance_window OdbCloudExadataInfrastructure#maintenance_window}
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#region OdbCloudExadataInfrastructure#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// TThe number of storage servers that are activated for the Exadata infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#storage_count OdbCloudExadataInfrastructure#storage_count}
	StorageCount *float64 `field:"optional" json:"storageCount" yaml:"storageCount"`
	// The storage server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the ListDbSystemShapes operation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#storage_server_type OdbCloudExadataInfrastructure#storage_server_type}
	StorageServerType *string `field:"optional" json:"storageServerType" yaml:"storageServerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#tags OdbCloudExadataInfrastructure#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/odb_cloud_exadata_infrastructure#timeouts OdbCloudExadataInfrastructure#timeouts}
	Timeouts *OdbCloudExadataInfrastructureTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

