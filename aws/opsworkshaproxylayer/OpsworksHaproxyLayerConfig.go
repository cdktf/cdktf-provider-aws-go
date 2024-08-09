// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworkshaproxylayer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksHaproxyLayerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#stack_id OpsworksHaproxyLayer#stack_id}.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#stats_password OpsworksHaproxyLayer#stats_password}.
	StatsPassword *string `field:"required" json:"statsPassword" yaml:"statsPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#auto_assign_elastic_ips OpsworksHaproxyLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `field:"optional" json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#auto_assign_public_ips OpsworksHaproxyLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `field:"optional" json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#auto_healing OpsworksHaproxyLayer#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#cloudwatch_configuration OpsworksHaproxyLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksHaproxyLayerCloudwatchConfiguration `field:"optional" json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_configure_recipes OpsworksHaproxyLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `field:"optional" json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_deploy_recipes OpsworksHaproxyLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `field:"optional" json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_instance_profile_arn OpsworksHaproxyLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `field:"optional" json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_json OpsworksHaproxyLayer#custom_json}.
	CustomJson *string `field:"optional" json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_security_group_ids OpsworksHaproxyLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `field:"optional" json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_setup_recipes OpsworksHaproxyLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `field:"optional" json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_shutdown_recipes OpsworksHaproxyLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `field:"optional" json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#custom_undeploy_recipes OpsworksHaproxyLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `field:"optional" json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#drain_elb_on_shutdown OpsworksHaproxyLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `field:"optional" json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#ebs_volume OpsworksHaproxyLayer#ebs_volume}
	EbsVolume interface{} `field:"optional" json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#elastic_load_balancer OpsworksHaproxyLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `field:"optional" json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#healthcheck_method OpsworksHaproxyLayer#healthcheck_method}.
	HealthcheckMethod *string `field:"optional" json:"healthcheckMethod" yaml:"healthcheckMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#healthcheck_url OpsworksHaproxyLayer#healthcheck_url}.
	HealthcheckUrl *string `field:"optional" json:"healthcheckUrl" yaml:"healthcheckUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#id OpsworksHaproxyLayer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#install_updates_on_boot OpsworksHaproxyLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#instance_shutdown_timeout OpsworksHaproxyLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `field:"optional" json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// load_based_auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#load_based_auto_scaling OpsworksHaproxyLayer#load_based_auto_scaling}
	LoadBasedAutoScaling *OpsworksHaproxyLayerLoadBasedAutoScaling `field:"optional" json:"loadBasedAutoScaling" yaml:"loadBasedAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#name OpsworksHaproxyLayer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#stats_enabled OpsworksHaproxyLayer#stats_enabled}.
	StatsEnabled interface{} `field:"optional" json:"statsEnabled" yaml:"statsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#stats_url OpsworksHaproxyLayer#stats_url}.
	StatsUrl *string `field:"optional" json:"statsUrl" yaml:"statsUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#stats_user OpsworksHaproxyLayer#stats_user}.
	StatsUser *string `field:"optional" json:"statsUser" yaml:"statsUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#system_packages OpsworksHaproxyLayer#system_packages}.
	SystemPackages *[]*string `field:"optional" json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#tags OpsworksHaproxyLayer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#tags_all OpsworksHaproxyLayer#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/opsworks_haproxy_layer#use_ebs_optimized_instances OpsworksHaproxyLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `field:"optional" json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

