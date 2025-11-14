// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package odbcloudvmcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OdbCloudVmClusterConfig struct {
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
	// The unique identifier of the Exadata infrastructure for this VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#cloud_exadata_infrastructure_id OdbCloudVmCluster#cloud_exadata_infrastructure_id}
	CloudExadataInfrastructureId *string `field:"required" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The number of CPU cores to enable on the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#cpu_core_count OdbCloudVmCluster#cpu_core_count}
	CpuCoreCount *float64 `field:"required" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The size of the data disk group, in terabytes (TBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#data_storage_size_in_tbs OdbCloudVmCluster#data_storage_size_in_tbs}
	DataStorageSizeInTbs *float64 `field:"required" json:"dataStorageSizeInTbs" yaml:"dataStorageSizeInTbs"`
	// The list of database servers for the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#db_servers OdbCloudVmCluster#db_servers}
	DbServers *[]*string `field:"required" json:"dbServers" yaml:"dbServers"`
	// A user-friendly name for the VM cluster. This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#display_name OdbCloudVmCluster#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// A valid software version of Oracle Grid Infrastructure (GI).
	//
	// To get the list of valid values, use the ListGiVersions operation and specify the shape of the Exadata infrastructure. Example: 19.0.0.0 This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#gi_version OdbCloudVmCluster#gi_version}
	GiVersion *string `field:"required" json:"giVersion" yaml:"giVersion"`
	// The host name prefix for the VM cluster.
	//
	// Constraints: - Can't be "localhost" or "hostname". - Can't contain "-version". - The maximum length of the combined hostname and domain is 63 characters. - The hostname must be unique within the subnet. This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#hostname_prefix OdbCloudVmCluster#hostname_prefix}
	HostnamePrefix *string `field:"required" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#odb_network_id OdbCloudVmCluster#odb_network_id}
	OdbNetworkId *string `field:"required" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#ssh_public_keys OdbCloudVmCluster#ssh_public_keys}
	SshPublicKeys *[]*string `field:"required" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The name of the Grid Infrastructure (GI) cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#cluster_name OdbCloudVmCluster#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#data_collection_options OdbCloudVmCluster#data_collection_options}
	DataCollectionOptions interface{} `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The amount of local node storage, in gigabytes (GBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#db_node_storage_size_in_gbs OdbCloudVmCluster#db_node_storage_size_in_gbs}
	DbNodeStorageSizeInGbs *float64 `field:"optional" json:"dbNodeStorageSizeInGbs" yaml:"dbNodeStorageSizeInGbs"`
	// Specifies whether to enable database backups to local Exadata storage for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#is_local_backup_enabled OdbCloudVmCluster#is_local_backup_enabled}
	IsLocalBackupEnabled interface{} `field:"optional" json:"isLocalBackupEnabled" yaml:"isLocalBackupEnabled"`
	// Specifies whether to create a sparse disk group for the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#is_sparse_diskgroup_enabled OdbCloudVmCluster#is_sparse_diskgroup_enabled}
	IsSparseDiskgroupEnabled interface{} `field:"optional" json:"isSparseDiskgroupEnabled" yaml:"isSparseDiskgroupEnabled"`
	// The Oracle license model to apply to the VM cluster. Default: LICENSE_INCLUDED. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#license_model OdbCloudVmCluster#license_model}
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The amount of memory, in gigabytes (GBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#memory_size_in_gbs OdbCloudVmCluster#memory_size_in_gbs}
	MemorySizeInGbs *float64 `field:"optional" json:"memorySizeInGbs" yaml:"memorySizeInGbs"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#region OdbCloudVmCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The port number for TCP connections to the single client access name (SCAN) listener.
	//
	// Valid values: 1024–8999 with the following exceptions: 2484 , 6100 , 6200 , 7060, 7070 , 7085 , and 7879Default: 1521. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#scan_listener_port_tcp OdbCloudVmCluster#scan_listener_port_tcp}
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#tags OdbCloudVmCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#timeouts OdbCloudVmCluster#timeouts}
	Timeouts *OdbCloudVmClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The configured time zone of the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/odb_cloud_vm_cluster#timezone OdbCloudVmCluster#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

