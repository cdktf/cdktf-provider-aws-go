// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksjavaapplayer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksJavaAppLayerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#stack_id OpsworksJavaAppLayer#stack_id}.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#app_server OpsworksJavaAppLayer#app_server}.
	AppServer *string `field:"optional" json:"appServer" yaml:"appServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#app_server_version OpsworksJavaAppLayer#app_server_version}.
	AppServerVersion *string `field:"optional" json:"appServerVersion" yaml:"appServerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#auto_assign_elastic_ips OpsworksJavaAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `field:"optional" json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#auto_assign_public_ips OpsworksJavaAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `field:"optional" json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#auto_healing OpsworksJavaAppLayer#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#cloudwatch_configuration OpsworksJavaAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksJavaAppLayerCloudwatchConfiguration `field:"optional" json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_configure_recipes OpsworksJavaAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `field:"optional" json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_deploy_recipes OpsworksJavaAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `field:"optional" json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_instance_profile_arn OpsworksJavaAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `field:"optional" json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_json OpsworksJavaAppLayer#custom_json}.
	CustomJson *string `field:"optional" json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_security_group_ids OpsworksJavaAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `field:"optional" json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_setup_recipes OpsworksJavaAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `field:"optional" json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_shutdown_recipes OpsworksJavaAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `field:"optional" json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#custom_undeploy_recipes OpsworksJavaAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `field:"optional" json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#drain_elb_on_shutdown OpsworksJavaAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `field:"optional" json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#ebs_volume OpsworksJavaAppLayer#ebs_volume}
	EbsVolume interface{} `field:"optional" json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#elastic_load_balancer OpsworksJavaAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `field:"optional" json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#id OpsworksJavaAppLayer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#install_updates_on_boot OpsworksJavaAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#instance_shutdown_timeout OpsworksJavaAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `field:"optional" json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#jvm_options OpsworksJavaAppLayer#jvm_options}.
	JvmOptions *string `field:"optional" json:"jvmOptions" yaml:"jvmOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#jvm_type OpsworksJavaAppLayer#jvm_type}.
	JvmType *string `field:"optional" json:"jvmType" yaml:"jvmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#jvm_version OpsworksJavaAppLayer#jvm_version}.
	JvmVersion *string `field:"optional" json:"jvmVersion" yaml:"jvmVersion"`
	// load_based_auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#load_based_auto_scaling OpsworksJavaAppLayer#load_based_auto_scaling}
	LoadBasedAutoScaling *OpsworksJavaAppLayerLoadBasedAutoScaling `field:"optional" json:"loadBasedAutoScaling" yaml:"loadBasedAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#name OpsworksJavaAppLayer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#system_packages OpsworksJavaAppLayer#system_packages}.
	SystemPackages *[]*string `field:"optional" json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#tags OpsworksJavaAppLayer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#tags_all OpsworksJavaAppLayer#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/opsworks_java_app_layer#use_ebs_optimized_instances OpsworksJavaAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `field:"optional" json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

