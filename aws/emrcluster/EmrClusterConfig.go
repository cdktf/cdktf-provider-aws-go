// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmrClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#name EmrCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#release_label EmrCluster#release_label}.
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#service_role EmrCluster#service_role}.
	ServiceRole *string `field:"required" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#additional_info EmrCluster#additional_info}.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#applications EmrCluster#applications}.
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#autoscaling_role EmrCluster#autoscaling_role}.
	AutoscalingRole *string `field:"optional" json:"autoscalingRole" yaml:"autoscalingRole"`
	// auto_termination_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#auto_termination_policy EmrCluster#auto_termination_policy}
	AutoTerminationPolicy *EmrClusterAutoTerminationPolicy `field:"optional" json:"autoTerminationPolicy" yaml:"autoTerminationPolicy"`
	// bootstrap_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#bootstrap_action EmrCluster#bootstrap_action}
	BootstrapAction interface{} `field:"optional" json:"bootstrapAction" yaml:"bootstrapAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#configurations EmrCluster#configurations}.
	Configurations *string `field:"optional" json:"configurations" yaml:"configurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#configurations_json EmrCluster#configurations_json}.
	ConfigurationsJson *string `field:"optional" json:"configurationsJson" yaml:"configurationsJson"`
	// core_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#core_instance_fleet EmrCluster#core_instance_fleet}
	CoreInstanceFleet *EmrClusterCoreInstanceFleet `field:"optional" json:"coreInstanceFleet" yaml:"coreInstanceFleet"`
	// core_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#core_instance_group EmrCluster#core_instance_group}
	CoreInstanceGroup *EmrClusterCoreInstanceGroup `field:"optional" json:"coreInstanceGroup" yaml:"coreInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#custom_ami_id EmrCluster#custom_ami_id}.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#ebs_root_volume_size EmrCluster#ebs_root_volume_size}.
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// ec2_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#ec2_attributes EmrCluster#ec2_attributes}
	Ec2Attributes *EmrClusterEc2Attributes `field:"optional" json:"ec2Attributes" yaml:"ec2Attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#id EmrCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#keep_job_flow_alive_when_no_steps EmrCluster#keep_job_flow_alive_when_no_steps}.
	KeepJobFlowAliveWhenNoSteps interface{} `field:"optional" json:"keepJobFlowAliveWhenNoSteps" yaml:"keepJobFlowAliveWhenNoSteps"`
	// kerberos_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#kerberos_attributes EmrCluster#kerberos_attributes}
	KerberosAttributes *EmrClusterKerberosAttributes `field:"optional" json:"kerberosAttributes" yaml:"kerberosAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#list_steps_states EmrCluster#list_steps_states}.
	ListStepsStates *[]*string `field:"optional" json:"listStepsStates" yaml:"listStepsStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#log_encryption_kms_key_id EmrCluster#log_encryption_kms_key_id}.
	LogEncryptionKmsKeyId *string `field:"optional" json:"logEncryptionKmsKeyId" yaml:"logEncryptionKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#log_uri EmrCluster#log_uri}.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// master_instance_fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#master_instance_fleet EmrCluster#master_instance_fleet}
	MasterInstanceFleet *EmrClusterMasterInstanceFleet `field:"optional" json:"masterInstanceFleet" yaml:"masterInstanceFleet"`
	// master_instance_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#master_instance_group EmrCluster#master_instance_group}
	MasterInstanceGroup *EmrClusterMasterInstanceGroup `field:"optional" json:"masterInstanceGroup" yaml:"masterInstanceGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#placement_group_config EmrCluster#placement_group_config}.
	PlacementGroupConfig interface{} `field:"optional" json:"placementGroupConfig" yaml:"placementGroupConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#scale_down_behavior EmrCluster#scale_down_behavior}.
	ScaleDownBehavior *string `field:"optional" json:"scaleDownBehavior" yaml:"scaleDownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#security_configuration EmrCluster#security_configuration}.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#step EmrCluster#step}.
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#step_concurrency_level EmrCluster#step_concurrency_level}.
	StepConcurrencyLevel *float64 `field:"optional" json:"stepConcurrencyLevel" yaml:"stepConcurrencyLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#tags EmrCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#tags_all EmrCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#termination_protection EmrCluster#termination_protection}.
	TerminationProtection interface{} `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#unhealthy_node_replacement EmrCluster#unhealthy_node_replacement}.
	UnhealthyNodeReplacement interface{} `field:"optional" json:"unhealthyNodeReplacement" yaml:"unhealthyNodeReplacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/emr_cluster#visible_to_all_users EmrCluster#visible_to_all_users}.
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

