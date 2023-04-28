package opsworksapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpsworksApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#name OpsworksApplication#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#stack_id OpsworksApplication#stack_id}.
	StackId *string `field:"required" json:"stackId" yaml:"stackId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#type OpsworksApplication#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// app_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#app_source OpsworksApplication#app_source}
	AppSource interface{} `field:"optional" json:"appSource" yaml:"appSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#auto_bundle_on_deploy OpsworksApplication#auto_bundle_on_deploy}.
	AutoBundleOnDeploy *string `field:"optional" json:"autoBundleOnDeploy" yaml:"autoBundleOnDeploy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#aws_flow_ruby_settings OpsworksApplication#aws_flow_ruby_settings}.
	AwsFlowRubySettings *string `field:"optional" json:"awsFlowRubySettings" yaml:"awsFlowRubySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#data_source_arn OpsworksApplication#data_source_arn}.
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#data_source_database_name OpsworksApplication#data_source_database_name}.
	DataSourceDatabaseName *string `field:"optional" json:"dataSourceDatabaseName" yaml:"dataSourceDatabaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#data_source_type OpsworksApplication#data_source_type}.
	DataSourceType *string `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#description OpsworksApplication#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#document_root OpsworksApplication#document_root}.
	DocumentRoot *string `field:"optional" json:"documentRoot" yaml:"documentRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#domains OpsworksApplication#domains}.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#enable_ssl OpsworksApplication#enable_ssl}.
	EnableSsl interface{} `field:"optional" json:"enableSsl" yaml:"enableSsl"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#environment OpsworksApplication#environment}
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#id OpsworksApplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#rails_env OpsworksApplication#rails_env}.
	RailsEnv *string `field:"optional" json:"railsEnv" yaml:"railsEnv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#short_name OpsworksApplication#short_name}.
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
	// ssl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.65.0/docs/resources/opsworks_application#ssl_configuration OpsworksApplication#ssl_configuration}
	SslConfiguration interface{} `field:"optional" json:"sslConfiguration" yaml:"sslConfiguration"`
}

