package opsworksrailsapplayer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksRailsAppLayerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#stack_id OpsworksRailsAppLayer#stack_id}.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#app_server OpsworksRailsAppLayer#app_server}.
	AppServer *string `field:"optional" json:"appServer" yaml:"appServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_assign_elastic_ips OpsworksRailsAppLayer#auto_assign_elastic_ips}.
	AutoAssignElasticIps interface{} `field:"optional" json:"autoAssignElasticIps" yaml:"autoAssignElasticIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_assign_public_ips OpsworksRailsAppLayer#auto_assign_public_ips}.
	AutoAssignPublicIps interface{} `field:"optional" json:"autoAssignPublicIps" yaml:"autoAssignPublicIps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#auto_healing OpsworksRailsAppLayer#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#bundler_version OpsworksRailsAppLayer#bundler_version}.
	BundlerVersion *string `field:"optional" json:"bundlerVersion" yaml:"bundlerVersion"`
	// cloudwatch_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#cloudwatch_configuration OpsworksRailsAppLayer#cloudwatch_configuration}
	CloudwatchConfiguration *OpsworksRailsAppLayerCloudwatchConfiguration `field:"optional" json:"cloudwatchConfiguration" yaml:"cloudwatchConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_configure_recipes OpsworksRailsAppLayer#custom_configure_recipes}.
	CustomConfigureRecipes *[]*string `field:"optional" json:"customConfigureRecipes" yaml:"customConfigureRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_deploy_recipes OpsworksRailsAppLayer#custom_deploy_recipes}.
	CustomDeployRecipes *[]*string `field:"optional" json:"customDeployRecipes" yaml:"customDeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_instance_profile_arn OpsworksRailsAppLayer#custom_instance_profile_arn}.
	CustomInstanceProfileArn *string `field:"optional" json:"customInstanceProfileArn" yaml:"customInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_json OpsworksRailsAppLayer#custom_json}.
	CustomJson *string `field:"optional" json:"customJson" yaml:"customJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_security_group_ids OpsworksRailsAppLayer#custom_security_group_ids}.
	CustomSecurityGroupIds *[]*string `field:"optional" json:"customSecurityGroupIds" yaml:"customSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_setup_recipes OpsworksRailsAppLayer#custom_setup_recipes}.
	CustomSetupRecipes *[]*string `field:"optional" json:"customSetupRecipes" yaml:"customSetupRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_shutdown_recipes OpsworksRailsAppLayer#custom_shutdown_recipes}.
	CustomShutdownRecipes *[]*string `field:"optional" json:"customShutdownRecipes" yaml:"customShutdownRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#custom_undeploy_recipes OpsworksRailsAppLayer#custom_undeploy_recipes}.
	CustomUndeployRecipes *[]*string `field:"optional" json:"customUndeployRecipes" yaml:"customUndeployRecipes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#drain_elb_on_shutdown OpsworksRailsAppLayer#drain_elb_on_shutdown}.
	DrainElbOnShutdown interface{} `field:"optional" json:"drainElbOnShutdown" yaml:"drainElbOnShutdown"`
	// ebs_volume block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#ebs_volume OpsworksRailsAppLayer#ebs_volume}
	EbsVolume interface{} `field:"optional" json:"ebsVolume" yaml:"ebsVolume"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#elastic_load_balancer OpsworksRailsAppLayer#elastic_load_balancer}.
	ElasticLoadBalancer *string `field:"optional" json:"elasticLoadBalancer" yaml:"elasticLoadBalancer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#id OpsworksRailsAppLayer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#install_updates_on_boot OpsworksRailsAppLayer#install_updates_on_boot}.
	InstallUpdatesOnBoot interface{} `field:"optional" json:"installUpdatesOnBoot" yaml:"installUpdatesOnBoot"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#instance_shutdown_timeout OpsworksRailsAppLayer#instance_shutdown_timeout}.
	InstanceShutdownTimeout *float64 `field:"optional" json:"instanceShutdownTimeout" yaml:"instanceShutdownTimeout"`
	// load_based_auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#load_based_auto_scaling OpsworksRailsAppLayer#load_based_auto_scaling}
	LoadBasedAutoScaling *OpsworksRailsAppLayerLoadBasedAutoScaling `field:"optional" json:"loadBasedAutoScaling" yaml:"loadBasedAutoScaling"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#manage_bundler OpsworksRailsAppLayer#manage_bundler}.
	ManageBundler interface{} `field:"optional" json:"manageBundler" yaml:"manageBundler"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#name OpsworksRailsAppLayer#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#passenger_version OpsworksRailsAppLayer#passenger_version}.
	PassengerVersion *string `field:"optional" json:"passengerVersion" yaml:"passengerVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#rubygems_version OpsworksRailsAppLayer#rubygems_version}.
	RubygemsVersion *string `field:"optional" json:"rubygemsVersion" yaml:"rubygemsVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#ruby_version OpsworksRailsAppLayer#ruby_version}.
	RubyVersion *string `field:"optional" json:"rubyVersion" yaml:"rubyVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#system_packages OpsworksRailsAppLayer#system_packages}.
	SystemPackages *[]*string `field:"optional" json:"systemPackages" yaml:"systemPackages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#tags OpsworksRailsAppLayer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#tags_all OpsworksRailsAppLayer#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/opsworks_rails_app_layer#use_ebs_optimized_instances OpsworksRailsAppLayer#use_ebs_optimized_instances}.
	UseEbsOptimizedInstances interface{} `field:"optional" json:"useEbsOptimizedInstances" yaml:"useEbsOptimizedInstances"`
}

