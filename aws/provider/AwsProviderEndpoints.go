// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsProviderEndpoints struct {
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#accessanalyzer AwsProvider#accessanalyzer}
	Accessanalyzer *string `field:"optional" json:"accessanalyzer" yaml:"accessanalyzer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#account AwsProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#acm AwsProvider#acm}
	Acm *string `field:"optional" json:"acm" yaml:"acm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#acmpca AwsProvider#acmpca}
	Acmpca *string `field:"optional" json:"acmpca" yaml:"acmpca"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#amg AwsProvider#amg}
	Amg *string `field:"optional" json:"amg" yaml:"amg"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#amp AwsProvider#amp}
	Amp *string `field:"optional" json:"amp" yaml:"amp"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#amplify AwsProvider#amplify}
	Amplify *string `field:"optional" json:"amplify" yaml:"amplify"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#apigateway AwsProvider#apigateway}
	Apigateway *string `field:"optional" json:"apigateway" yaml:"apigateway"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#apigatewayv2 AwsProvider#apigatewayv2}
	Apigatewayv2 *string `field:"optional" json:"apigatewayv2" yaml:"apigatewayv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appautoscaling AwsProvider#appautoscaling}
	Appautoscaling *string `field:"optional" json:"appautoscaling" yaml:"appautoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appconfig AwsProvider#appconfig}
	Appconfig *string `field:"optional" json:"appconfig" yaml:"appconfig"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appfabric AwsProvider#appfabric}
	Appfabric *string `field:"optional" json:"appfabric" yaml:"appfabric"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appflow AwsProvider#appflow}
	Appflow *string `field:"optional" json:"appflow" yaml:"appflow"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appintegrations AwsProvider#appintegrations}
	Appintegrations *string `field:"optional" json:"appintegrations" yaml:"appintegrations"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appintegrationsservice AwsProvider#appintegrationsservice}
	Appintegrationsservice *string `field:"optional" json:"appintegrationsservice" yaml:"appintegrationsservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#applicationautoscaling AwsProvider#applicationautoscaling}
	Applicationautoscaling *string `field:"optional" json:"applicationautoscaling" yaml:"applicationautoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#applicationinsights AwsProvider#applicationinsights}
	Applicationinsights *string `field:"optional" json:"applicationinsights" yaml:"applicationinsights"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#applicationsignals AwsProvider#applicationsignals}
	Applicationsignals *string `field:"optional" json:"applicationsignals" yaml:"applicationsignals"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appmesh AwsProvider#appmesh}
	Appmesh *string `field:"optional" json:"appmesh" yaml:"appmesh"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appregistry AwsProvider#appregistry}
	Appregistry *string `field:"optional" json:"appregistry" yaml:"appregistry"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#apprunner AwsProvider#apprunner}
	Apprunner *string `field:"optional" json:"apprunner" yaml:"apprunner"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appstream AwsProvider#appstream}
	Appstream *string `field:"optional" json:"appstream" yaml:"appstream"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#appsync AwsProvider#appsync}
	Appsync *string `field:"optional" json:"appsync" yaml:"appsync"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#athena AwsProvider#athena}
	Athena *string `field:"optional" json:"athena" yaml:"athena"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#auditmanager AwsProvider#auditmanager}
	Auditmanager *string `field:"optional" json:"auditmanager" yaml:"auditmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#autoscaling AwsProvider#autoscaling}
	Autoscaling *string `field:"optional" json:"autoscaling" yaml:"autoscaling"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#autoscalingplans AwsProvider#autoscalingplans}
	Autoscalingplans *string `field:"optional" json:"autoscalingplans" yaml:"autoscalingplans"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#backup AwsProvider#backup}
	Backup *string `field:"optional" json:"backup" yaml:"backup"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#batch AwsProvider#batch}
	Batch *string `field:"optional" json:"batch" yaml:"batch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#bcmdataexports AwsProvider#bcmdataexports}
	Bcmdataexports *string `field:"optional" json:"bcmdataexports" yaml:"bcmdataexports"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#beanstalk AwsProvider#beanstalk}
	Beanstalk *string `field:"optional" json:"beanstalk" yaml:"beanstalk"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#bedrock AwsProvider#bedrock}
	Bedrock *string `field:"optional" json:"bedrock" yaml:"bedrock"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#bedrockagent AwsProvider#bedrockagent}
	Bedrockagent *string `field:"optional" json:"bedrockagent" yaml:"bedrockagent"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#billing AwsProvider#billing}
	Billing *string `field:"optional" json:"billing" yaml:"billing"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#budgets AwsProvider#budgets}
	Budgets *string `field:"optional" json:"budgets" yaml:"budgets"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ce AwsProvider#ce}
	Ce *string `field:"optional" json:"ce" yaml:"ce"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#chatbot AwsProvider#chatbot}
	Chatbot *string `field:"optional" json:"chatbot" yaml:"chatbot"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#chime AwsProvider#chime}
	Chime *string `field:"optional" json:"chime" yaml:"chime"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#chimesdkmediapipelines AwsProvider#chimesdkmediapipelines}
	Chimesdkmediapipelines *string `field:"optional" json:"chimesdkmediapipelines" yaml:"chimesdkmediapipelines"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#chimesdkvoice AwsProvider#chimesdkvoice}
	Chimesdkvoice *string `field:"optional" json:"chimesdkvoice" yaml:"chimesdkvoice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cleanrooms AwsProvider#cleanrooms}
	Cleanrooms *string `field:"optional" json:"cleanrooms" yaml:"cleanrooms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloud9 AwsProvider#cloud9}
	Cloud9 *string `field:"optional" json:"cloud9" yaml:"cloud9"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudcontrol AwsProvider#cloudcontrol}
	Cloudcontrol *string `field:"optional" json:"cloudcontrol" yaml:"cloudcontrol"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudcontrolapi AwsProvider#cloudcontrolapi}
	Cloudcontrolapi *string `field:"optional" json:"cloudcontrolapi" yaml:"cloudcontrolapi"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudformation AwsProvider#cloudformation}
	Cloudformation *string `field:"optional" json:"cloudformation" yaml:"cloudformation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudfront AwsProvider#cloudfront}
	Cloudfront *string `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudfrontkeyvaluestore AwsProvider#cloudfrontkeyvaluestore}
	Cloudfrontkeyvaluestore *string `field:"optional" json:"cloudfrontkeyvaluestore" yaml:"cloudfrontkeyvaluestore"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudhsm AwsProvider#cloudhsm}
	Cloudhsm *string `field:"optional" json:"cloudhsm" yaml:"cloudhsm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudhsmv2 AwsProvider#cloudhsmv2}
	Cloudhsmv2 *string `field:"optional" json:"cloudhsmv2" yaml:"cloudhsmv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudsearch AwsProvider#cloudsearch}
	Cloudsearch *string `field:"optional" json:"cloudsearch" yaml:"cloudsearch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudtrail AwsProvider#cloudtrail}
	Cloudtrail *string `field:"optional" json:"cloudtrail" yaml:"cloudtrail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatch AwsProvider#cloudwatch}
	Cloudwatch *string `field:"optional" json:"cloudwatch" yaml:"cloudwatch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchevents AwsProvider#cloudwatchevents}
	Cloudwatchevents *string `field:"optional" json:"cloudwatchevents" yaml:"cloudwatchevents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchevidently AwsProvider#cloudwatchevidently}
	Cloudwatchevidently *string `field:"optional" json:"cloudwatchevidently" yaml:"cloudwatchevidently"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchlog AwsProvider#cloudwatchlog}
	Cloudwatchlog *string `field:"optional" json:"cloudwatchlog" yaml:"cloudwatchlog"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchlogs AwsProvider#cloudwatchlogs}
	Cloudwatchlogs *string `field:"optional" json:"cloudwatchlogs" yaml:"cloudwatchlogs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchobservabilityaccessmanager AwsProvider#cloudwatchobservabilityaccessmanager}
	Cloudwatchobservabilityaccessmanager *string `field:"optional" json:"cloudwatchobservabilityaccessmanager" yaml:"cloudwatchobservabilityaccessmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cloudwatchrum AwsProvider#cloudwatchrum}
	Cloudwatchrum *string `field:"optional" json:"cloudwatchrum" yaml:"cloudwatchrum"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codeartifact AwsProvider#codeartifact}
	Codeartifact *string `field:"optional" json:"codeartifact" yaml:"codeartifact"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codebuild AwsProvider#codebuild}
	Codebuild *string `field:"optional" json:"codebuild" yaml:"codebuild"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codecatalyst AwsProvider#codecatalyst}
	Codecatalyst *string `field:"optional" json:"codecatalyst" yaml:"codecatalyst"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codecommit AwsProvider#codecommit}
	Codecommit *string `field:"optional" json:"codecommit" yaml:"codecommit"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codeconnections AwsProvider#codeconnections}
	Codeconnections *string `field:"optional" json:"codeconnections" yaml:"codeconnections"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codedeploy AwsProvider#codedeploy}
	Codedeploy *string `field:"optional" json:"codedeploy" yaml:"codedeploy"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codeguruprofiler AwsProvider#codeguruprofiler}
	Codeguruprofiler *string `field:"optional" json:"codeguruprofiler" yaml:"codeguruprofiler"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codegurureviewer AwsProvider#codegurureviewer}
	Codegurureviewer *string `field:"optional" json:"codegurureviewer" yaml:"codegurureviewer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codepipeline AwsProvider#codepipeline}
	Codepipeline *string `field:"optional" json:"codepipeline" yaml:"codepipeline"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codestarconnections AwsProvider#codestarconnections}
	Codestarconnections *string `field:"optional" json:"codestarconnections" yaml:"codestarconnections"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#codestarnotifications AwsProvider#codestarnotifications}
	Codestarnotifications *string `field:"optional" json:"codestarnotifications" yaml:"codestarnotifications"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cognitoidentity AwsProvider#cognitoidentity}
	Cognitoidentity *string `field:"optional" json:"cognitoidentity" yaml:"cognitoidentity"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cognitoidentityprovider AwsProvider#cognitoidentityprovider}
	Cognitoidentityprovider *string `field:"optional" json:"cognitoidentityprovider" yaml:"cognitoidentityprovider"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cognitoidp AwsProvider#cognitoidp}
	Cognitoidp *string `field:"optional" json:"cognitoidp" yaml:"cognitoidp"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#comprehend AwsProvider#comprehend}
	Comprehend *string `field:"optional" json:"comprehend" yaml:"comprehend"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#computeoptimizer AwsProvider#computeoptimizer}
	Computeoptimizer *string `field:"optional" json:"computeoptimizer" yaml:"computeoptimizer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#config AwsProvider#config}
	Config *string `field:"optional" json:"config" yaml:"config"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#configservice AwsProvider#configservice}
	Configservice *string `field:"optional" json:"configservice" yaml:"configservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#connect AwsProvider#connect}
	Connect *string `field:"optional" json:"connect" yaml:"connect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#connectcases AwsProvider#connectcases}
	Connectcases *string `field:"optional" json:"connectcases" yaml:"connectcases"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#controltower AwsProvider#controltower}
	Controltower *string `field:"optional" json:"controltower" yaml:"controltower"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#costandusagereportservice AwsProvider#costandusagereportservice}
	Costandusagereportservice *string `field:"optional" json:"costandusagereportservice" yaml:"costandusagereportservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#costexplorer AwsProvider#costexplorer}
	Costexplorer *string `field:"optional" json:"costexplorer" yaml:"costexplorer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#costoptimizationhub AwsProvider#costoptimizationhub}
	Costoptimizationhub *string `field:"optional" json:"costoptimizationhub" yaml:"costoptimizationhub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#cur AwsProvider#cur}
	Cur *string `field:"optional" json:"cur" yaml:"cur"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#customerprofiles AwsProvider#customerprofiles}
	Customerprofiles *string `field:"optional" json:"customerprofiles" yaml:"customerprofiles"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#databasemigration AwsProvider#databasemigration}
	Databasemigration *string `field:"optional" json:"databasemigration" yaml:"databasemigration"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#databasemigrationservice AwsProvider#databasemigrationservice}
	Databasemigrationservice *string `field:"optional" json:"databasemigrationservice" yaml:"databasemigrationservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#databrew AwsProvider#databrew}
	Databrew *string `field:"optional" json:"databrew" yaml:"databrew"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dataexchange AwsProvider#dataexchange}
	Dataexchange *string `field:"optional" json:"dataexchange" yaml:"dataexchange"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#datapipeline AwsProvider#datapipeline}
	Datapipeline *string `field:"optional" json:"datapipeline" yaml:"datapipeline"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#datasync AwsProvider#datasync}
	Datasync *string `field:"optional" json:"datasync" yaml:"datasync"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#datazone AwsProvider#datazone}
	Datazone *string `field:"optional" json:"datazone" yaml:"datazone"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dax AwsProvider#dax}
	Dax *string `field:"optional" json:"dax" yaml:"dax"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#deploy AwsProvider#deploy}
	Deploy *string `field:"optional" json:"deploy" yaml:"deploy"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#detective AwsProvider#detective}
	Detective *string `field:"optional" json:"detective" yaml:"detective"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#devicefarm AwsProvider#devicefarm}
	Devicefarm *string `field:"optional" json:"devicefarm" yaml:"devicefarm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#devopsguru AwsProvider#devopsguru}
	Devopsguru *string `field:"optional" json:"devopsguru" yaml:"devopsguru"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#directconnect AwsProvider#directconnect}
	Directconnect *string `field:"optional" json:"directconnect" yaml:"directconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#directoryservice AwsProvider#directoryservice}
	Directoryservice *string `field:"optional" json:"directoryservice" yaml:"directoryservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dlm AwsProvider#dlm}
	Dlm *string `field:"optional" json:"dlm" yaml:"dlm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dms AwsProvider#dms}
	Dms *string `field:"optional" json:"dms" yaml:"dms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#docdb AwsProvider#docdb}
	Docdb *string `field:"optional" json:"docdb" yaml:"docdb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#docdbelastic AwsProvider#docdbelastic}
	Docdbelastic *string `field:"optional" json:"docdbelastic" yaml:"docdbelastic"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#drs AwsProvider#drs}
	Drs *string `field:"optional" json:"drs" yaml:"drs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ds AwsProvider#ds}
	Ds *string `field:"optional" json:"ds" yaml:"ds"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dsql AwsProvider#dsql}
	Dsql *string `field:"optional" json:"dsql" yaml:"dsql"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#dynamodb AwsProvider#dynamodb}
	Dynamodb *string `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ec2 AwsProvider#ec2}
	Ec2 *string `field:"optional" json:"ec2" yaml:"ec2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ecr AwsProvider#ecr}
	Ecr *string `field:"optional" json:"ecr" yaml:"ecr"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ecrpublic AwsProvider#ecrpublic}
	Ecrpublic *string `field:"optional" json:"ecrpublic" yaml:"ecrpublic"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ecs AwsProvider#ecs}
	Ecs *string `field:"optional" json:"ecs" yaml:"ecs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#efs AwsProvider#efs}
	Efs *string `field:"optional" json:"efs" yaml:"efs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#eks AwsProvider#eks}
	Eks *string `field:"optional" json:"eks" yaml:"eks"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticache AwsProvider#elasticache}
	Elasticache *string `field:"optional" json:"elasticache" yaml:"elasticache"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticbeanstalk AwsProvider#elasticbeanstalk}
	Elasticbeanstalk *string `field:"optional" json:"elasticbeanstalk" yaml:"elasticbeanstalk"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticloadbalancing AwsProvider#elasticloadbalancing}
	Elasticloadbalancing *string `field:"optional" json:"elasticloadbalancing" yaml:"elasticloadbalancing"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticloadbalancingv2 AwsProvider#elasticloadbalancingv2}
	Elasticloadbalancingv2 *string `field:"optional" json:"elasticloadbalancingv2" yaml:"elasticloadbalancingv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticsearch AwsProvider#elasticsearch}
	Elasticsearch *string `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elasticsearchservice AwsProvider#elasticsearchservice}
	Elasticsearchservice *string `field:"optional" json:"elasticsearchservice" yaml:"elasticsearchservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elastictranscoder AwsProvider#elastictranscoder}
	Elastictranscoder *string `field:"optional" json:"elastictranscoder" yaml:"elastictranscoder"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elb AwsProvider#elb}
	Elb *string `field:"optional" json:"elb" yaml:"elb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#elbv2 AwsProvider#elbv2}
	Elbv2 *string `field:"optional" json:"elbv2" yaml:"elbv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#emr AwsProvider#emr}
	Emr *string `field:"optional" json:"emr" yaml:"emr"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#emrcontainers AwsProvider#emrcontainers}
	Emrcontainers *string `field:"optional" json:"emrcontainers" yaml:"emrcontainers"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#emrserverless AwsProvider#emrserverless}
	Emrserverless *string `field:"optional" json:"emrserverless" yaml:"emrserverless"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#es AwsProvider#es}
	Es *string `field:"optional" json:"es" yaml:"es"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#eventbridge AwsProvider#eventbridge}
	Eventbridge *string `field:"optional" json:"eventbridge" yaml:"eventbridge"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#events AwsProvider#events}
	Events *string `field:"optional" json:"events" yaml:"events"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#evidently AwsProvider#evidently}
	Evidently *string `field:"optional" json:"evidently" yaml:"evidently"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#evs AwsProvider#evs}
	Evs *string `field:"optional" json:"evs" yaml:"evs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#finspace AwsProvider#finspace}
	Finspace *string `field:"optional" json:"finspace" yaml:"finspace"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#firehose AwsProvider#firehose}
	Firehose *string `field:"optional" json:"firehose" yaml:"firehose"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#fis AwsProvider#fis}
	Fis *string `field:"optional" json:"fis" yaml:"fis"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#fms AwsProvider#fms}
	Fms *string `field:"optional" json:"fms" yaml:"fms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#fsx AwsProvider#fsx}
	Fsx *string `field:"optional" json:"fsx" yaml:"fsx"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#gamelift AwsProvider#gamelift}
	Gamelift *string `field:"optional" json:"gamelift" yaml:"gamelift"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#glacier AwsProvider#glacier}
	Glacier *string `field:"optional" json:"glacier" yaml:"glacier"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#globalaccelerator AwsProvider#globalaccelerator}
	Globalaccelerator *string `field:"optional" json:"globalaccelerator" yaml:"globalaccelerator"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#glue AwsProvider#glue}
	Glue *string `field:"optional" json:"glue" yaml:"glue"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#gluedatabrew AwsProvider#gluedatabrew}
	Gluedatabrew *string `field:"optional" json:"gluedatabrew" yaml:"gluedatabrew"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#grafana AwsProvider#grafana}
	Grafana *string `field:"optional" json:"grafana" yaml:"grafana"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#greengrass AwsProvider#greengrass}
	Greengrass *string `field:"optional" json:"greengrass" yaml:"greengrass"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#groundstation AwsProvider#groundstation}
	Groundstation *string `field:"optional" json:"groundstation" yaml:"groundstation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#guardduty AwsProvider#guardduty}
	Guardduty *string `field:"optional" json:"guardduty" yaml:"guardduty"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#healthlake AwsProvider#healthlake}
	Healthlake *string `field:"optional" json:"healthlake" yaml:"healthlake"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#iam AwsProvider#iam}
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#identitystore AwsProvider#identitystore}
	Identitystore *string `field:"optional" json:"identitystore" yaml:"identitystore"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#imagebuilder AwsProvider#imagebuilder}
	Imagebuilder *string `field:"optional" json:"imagebuilder" yaml:"imagebuilder"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#inspector AwsProvider#inspector}
	Inspector *string `field:"optional" json:"inspector" yaml:"inspector"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#inspector2 AwsProvider#inspector2}
	Inspector2 *string `field:"optional" json:"inspector2" yaml:"inspector2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#inspectorv2 AwsProvider#inspectorv2}
	Inspectorv2 *string `field:"optional" json:"inspectorv2" yaml:"inspectorv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#internetmonitor AwsProvider#internetmonitor}
	Internetmonitor *string `field:"optional" json:"internetmonitor" yaml:"internetmonitor"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#invoicing AwsProvider#invoicing}
	Invoicing *string `field:"optional" json:"invoicing" yaml:"invoicing"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#iot AwsProvider#iot}
	Iot *string `field:"optional" json:"iot" yaml:"iot"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ivs AwsProvider#ivs}
	Ivs *string `field:"optional" json:"ivs" yaml:"ivs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ivschat AwsProvider#ivschat}
	Ivschat *string `field:"optional" json:"ivschat" yaml:"ivschat"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kafka AwsProvider#kafka}
	Kafka *string `field:"optional" json:"kafka" yaml:"kafka"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kafkaconnect AwsProvider#kafkaconnect}
	Kafkaconnect *string `field:"optional" json:"kafkaconnect" yaml:"kafkaconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kendra AwsProvider#kendra}
	Kendra *string `field:"optional" json:"kendra" yaml:"kendra"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#keyspaces AwsProvider#keyspaces}
	Keyspaces *string `field:"optional" json:"keyspaces" yaml:"keyspaces"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kinesis AwsProvider#kinesis}
	Kinesis *string `field:"optional" json:"kinesis" yaml:"kinesis"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kinesisanalytics AwsProvider#kinesisanalytics}
	Kinesisanalytics *string `field:"optional" json:"kinesisanalytics" yaml:"kinesisanalytics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kinesisanalyticsv2 AwsProvider#kinesisanalyticsv2}
	Kinesisanalyticsv2 *string `field:"optional" json:"kinesisanalyticsv2" yaml:"kinesisanalyticsv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kinesisvideo AwsProvider#kinesisvideo}
	Kinesisvideo *string `field:"optional" json:"kinesisvideo" yaml:"kinesisvideo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#kms AwsProvider#kms}
	Kms *string `field:"optional" json:"kms" yaml:"kms"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lakeformation AwsProvider#lakeformation}
	Lakeformation *string `field:"optional" json:"lakeformation" yaml:"lakeformation"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lambda AwsProvider#lambda}
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#launchwizard AwsProvider#launchwizard}
	Launchwizard *string `field:"optional" json:"launchwizard" yaml:"launchwizard"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lex AwsProvider#lex}
	Lex *string `field:"optional" json:"lex" yaml:"lex"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lexmodelbuilding AwsProvider#lexmodelbuilding}
	Lexmodelbuilding *string `field:"optional" json:"lexmodelbuilding" yaml:"lexmodelbuilding"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lexmodelbuildingservice AwsProvider#lexmodelbuildingservice}
	Lexmodelbuildingservice *string `field:"optional" json:"lexmodelbuildingservice" yaml:"lexmodelbuildingservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lexmodels AwsProvider#lexmodels}
	Lexmodels *string `field:"optional" json:"lexmodels" yaml:"lexmodels"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lexmodelsv2 AwsProvider#lexmodelsv2}
	Lexmodelsv2 *string `field:"optional" json:"lexmodelsv2" yaml:"lexmodelsv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lexv2models AwsProvider#lexv2models}
	Lexv2Models *string `field:"optional" json:"lexv2Models" yaml:"lexv2Models"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#licensemanager AwsProvider#licensemanager}
	Licensemanager *string `field:"optional" json:"licensemanager" yaml:"licensemanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lightsail AwsProvider#lightsail}
	Lightsail *string `field:"optional" json:"lightsail" yaml:"lightsail"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#location AwsProvider#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#locationservice AwsProvider#locationservice}
	Locationservice *string `field:"optional" json:"locationservice" yaml:"locationservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#logs AwsProvider#logs}
	Logs *string `field:"optional" json:"logs" yaml:"logs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#lookoutmetrics AwsProvider#lookoutmetrics}
	Lookoutmetrics *string `field:"optional" json:"lookoutmetrics" yaml:"lookoutmetrics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#m2 AwsProvider#m2}
	M2 *string `field:"optional" json:"m2" yaml:"m2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#macie2 AwsProvider#macie2}
	Macie2 *string `field:"optional" json:"macie2" yaml:"macie2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#managedgrafana AwsProvider#managedgrafana}
	Managedgrafana *string `field:"optional" json:"managedgrafana" yaml:"managedgrafana"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediaconnect AwsProvider#mediaconnect}
	Mediaconnect *string `field:"optional" json:"mediaconnect" yaml:"mediaconnect"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediaconvert AwsProvider#mediaconvert}
	Mediaconvert *string `field:"optional" json:"mediaconvert" yaml:"mediaconvert"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#medialive AwsProvider#medialive}
	Medialive *string `field:"optional" json:"medialive" yaml:"medialive"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediapackage AwsProvider#mediapackage}
	Mediapackage *string `field:"optional" json:"mediapackage" yaml:"mediapackage"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediapackagev2 AwsProvider#mediapackagev2}
	Mediapackagev2 *string `field:"optional" json:"mediapackagev2" yaml:"mediapackagev2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediapackagevod AwsProvider#mediapackagevod}
	Mediapackagevod *string `field:"optional" json:"mediapackagevod" yaml:"mediapackagevod"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mediastore AwsProvider#mediastore}
	Mediastore *string `field:"optional" json:"mediastore" yaml:"mediastore"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#memorydb AwsProvider#memorydb}
	Memorydb *string `field:"optional" json:"memorydb" yaml:"memorydb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mgn AwsProvider#mgn}
	Mgn *string `field:"optional" json:"mgn" yaml:"mgn"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mq AwsProvider#mq}
	Mq *string `field:"optional" json:"mq" yaml:"mq"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#msk AwsProvider#msk}
	Msk *string `field:"optional" json:"msk" yaml:"msk"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#mwaa AwsProvider#mwaa}
	Mwaa *string `field:"optional" json:"mwaa" yaml:"mwaa"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#neptune AwsProvider#neptune}
	Neptune *string `field:"optional" json:"neptune" yaml:"neptune"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#neptunegraph AwsProvider#neptunegraph}
	Neptunegraph *string `field:"optional" json:"neptunegraph" yaml:"neptunegraph"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#networkfirewall AwsProvider#networkfirewall}
	Networkfirewall *string `field:"optional" json:"networkfirewall" yaml:"networkfirewall"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#networkmanager AwsProvider#networkmanager}
	Networkmanager *string `field:"optional" json:"networkmanager" yaml:"networkmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#networkmonitor AwsProvider#networkmonitor}
	Networkmonitor *string `field:"optional" json:"networkmonitor" yaml:"networkmonitor"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#notifications AwsProvider#notifications}
	Notifications *string `field:"optional" json:"notifications" yaml:"notifications"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#notificationscontacts AwsProvider#notificationscontacts}
	Notificationscontacts *string `field:"optional" json:"notificationscontacts" yaml:"notificationscontacts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#oam AwsProvider#oam}
	Oam *string `field:"optional" json:"oam" yaml:"oam"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#opensearch AwsProvider#opensearch}
	Opensearch *string `field:"optional" json:"opensearch" yaml:"opensearch"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#opensearchingestion AwsProvider#opensearchingestion}
	Opensearchingestion *string `field:"optional" json:"opensearchingestion" yaml:"opensearchingestion"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#opensearchserverless AwsProvider#opensearchserverless}
	Opensearchserverless *string `field:"optional" json:"opensearchserverless" yaml:"opensearchserverless"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#opensearchservice AwsProvider#opensearchservice}
	Opensearchservice *string `field:"optional" json:"opensearchservice" yaml:"opensearchservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#organizations AwsProvider#organizations}
	Organizations *string `field:"optional" json:"organizations" yaml:"organizations"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#osis AwsProvider#osis}
	Osis *string `field:"optional" json:"osis" yaml:"osis"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#outposts AwsProvider#outposts}
	Outposts *string `field:"optional" json:"outposts" yaml:"outposts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#paymentcryptography AwsProvider#paymentcryptography}
	Paymentcryptography *string `field:"optional" json:"paymentcryptography" yaml:"paymentcryptography"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pcaconnectorad AwsProvider#pcaconnectorad}
	Pcaconnectorad *string `field:"optional" json:"pcaconnectorad" yaml:"pcaconnectorad"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pcs AwsProvider#pcs}
	Pcs *string `field:"optional" json:"pcs" yaml:"pcs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pinpoint AwsProvider#pinpoint}
	Pinpoint *string `field:"optional" json:"pinpoint" yaml:"pinpoint"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pinpointsmsvoicev2 AwsProvider#pinpointsmsvoicev2}
	Pinpointsmsvoicev2 *string `field:"optional" json:"pinpointsmsvoicev2" yaml:"pinpointsmsvoicev2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pipes AwsProvider#pipes}
	Pipes *string `field:"optional" json:"pipes" yaml:"pipes"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#polly AwsProvider#polly}
	Polly *string `field:"optional" json:"polly" yaml:"polly"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#pricing AwsProvider#pricing}
	Pricing *string `field:"optional" json:"pricing" yaml:"pricing"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#prometheus AwsProvider#prometheus}
	Prometheus *string `field:"optional" json:"prometheus" yaml:"prometheus"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#prometheusservice AwsProvider#prometheusservice}
	Prometheusservice *string `field:"optional" json:"prometheusservice" yaml:"prometheusservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#qbusiness AwsProvider#qbusiness}
	Qbusiness *string `field:"optional" json:"qbusiness" yaml:"qbusiness"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#qldb AwsProvider#qldb}
	Qldb *string `field:"optional" json:"qldb" yaml:"qldb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#quicksight AwsProvider#quicksight}
	Quicksight *string `field:"optional" json:"quicksight" yaml:"quicksight"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ram AwsProvider#ram}
	Ram *string `field:"optional" json:"ram" yaml:"ram"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#rbin AwsProvider#rbin}
	Rbin *string `field:"optional" json:"rbin" yaml:"rbin"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#rds AwsProvider#rds}
	Rds *string `field:"optional" json:"rds" yaml:"rds"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#recyclebin AwsProvider#recyclebin}
	Recyclebin *string `field:"optional" json:"recyclebin" yaml:"recyclebin"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#redshift AwsProvider#redshift}
	Redshift *string `field:"optional" json:"redshift" yaml:"redshift"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#redshiftdata AwsProvider#redshiftdata}
	Redshiftdata *string `field:"optional" json:"redshiftdata" yaml:"redshiftdata"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#redshiftdataapiservice AwsProvider#redshiftdataapiservice}
	Redshiftdataapiservice *string `field:"optional" json:"redshiftdataapiservice" yaml:"redshiftdataapiservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#redshiftserverless AwsProvider#redshiftserverless}
	Redshiftserverless *string `field:"optional" json:"redshiftserverless" yaml:"redshiftserverless"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#rekognition AwsProvider#rekognition}
	Rekognition *string `field:"optional" json:"rekognition" yaml:"rekognition"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#resiliencehub AwsProvider#resiliencehub}
	Resiliencehub *string `field:"optional" json:"resiliencehub" yaml:"resiliencehub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#resourceexplorer2 AwsProvider#resourceexplorer2}
	Resourceexplorer2 *string `field:"optional" json:"resourceexplorer2" yaml:"resourceexplorer2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#resourcegroups AwsProvider#resourcegroups}
	Resourcegroups *string `field:"optional" json:"resourcegroups" yaml:"resourcegroups"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#resourcegroupstagging AwsProvider#resourcegroupstagging}
	Resourcegroupstagging *string `field:"optional" json:"resourcegroupstagging" yaml:"resourcegroupstagging"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#resourcegroupstaggingapi AwsProvider#resourcegroupstaggingapi}
	Resourcegroupstaggingapi *string `field:"optional" json:"resourcegroupstaggingapi" yaml:"resourcegroupstaggingapi"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#rolesanywhere AwsProvider#rolesanywhere}
	Rolesanywhere *string `field:"optional" json:"rolesanywhere" yaml:"rolesanywhere"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53 AwsProvider#route53}
	Route53 *string `field:"optional" json:"route53" yaml:"route53"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53domains AwsProvider#route53domains}
	Route53Domains *string `field:"optional" json:"route53Domains" yaml:"route53Domains"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53profiles AwsProvider#route53profiles}
	Route53Profiles *string `field:"optional" json:"route53Profiles" yaml:"route53Profiles"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53recoverycontrolconfig AwsProvider#route53recoverycontrolconfig}
	Route53Recoverycontrolconfig *string `field:"optional" json:"route53Recoverycontrolconfig" yaml:"route53Recoverycontrolconfig"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53recoveryreadiness AwsProvider#route53recoveryreadiness}
	Route53Recoveryreadiness *string `field:"optional" json:"route53Recoveryreadiness" yaml:"route53Recoveryreadiness"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#route53resolver AwsProvider#route53resolver}
	Route53Resolver *string `field:"optional" json:"route53Resolver" yaml:"route53Resolver"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#rum AwsProvider#rum}
	Rum *string `field:"optional" json:"rum" yaml:"rum"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#s3 AwsProvider#s3}
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#s3api AwsProvider#s3api}
	S3Api *string `field:"optional" json:"s3Api" yaml:"s3Api"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#s3control AwsProvider#s3control}
	S3Control *string `field:"optional" json:"s3Control" yaml:"s3Control"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#s3outposts AwsProvider#s3outposts}
	S3Outposts *string `field:"optional" json:"s3Outposts" yaml:"s3Outposts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#s3tables AwsProvider#s3tables}
	S3Tables *string `field:"optional" json:"s3Tables" yaml:"s3Tables"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sagemaker AwsProvider#sagemaker}
	Sagemaker *string `field:"optional" json:"sagemaker" yaml:"sagemaker"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#scheduler AwsProvider#scheduler}
	Scheduler *string `field:"optional" json:"scheduler" yaml:"scheduler"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#schemas AwsProvider#schemas}
	Schemas *string `field:"optional" json:"schemas" yaml:"schemas"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#secretsmanager AwsProvider#secretsmanager}
	Secretsmanager *string `field:"optional" json:"secretsmanager" yaml:"secretsmanager"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#securityhub AwsProvider#securityhub}
	Securityhub *string `field:"optional" json:"securityhub" yaml:"securityhub"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#securitylake AwsProvider#securitylake}
	Securitylake *string `field:"optional" json:"securitylake" yaml:"securitylake"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#serverlessapplicationrepository AwsProvider#serverlessapplicationrepository}
	Serverlessapplicationrepository *string `field:"optional" json:"serverlessapplicationrepository" yaml:"serverlessapplicationrepository"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#serverlessapprepo AwsProvider#serverlessapprepo}
	Serverlessapprepo *string `field:"optional" json:"serverlessapprepo" yaml:"serverlessapprepo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#serverlessrepo AwsProvider#serverlessrepo}
	Serverlessrepo *string `field:"optional" json:"serverlessrepo" yaml:"serverlessrepo"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#servicecatalog AwsProvider#servicecatalog}
	Servicecatalog *string `field:"optional" json:"servicecatalog" yaml:"servicecatalog"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#servicecatalogappregistry AwsProvider#servicecatalogappregistry}
	Servicecatalogappregistry *string `field:"optional" json:"servicecatalogappregistry" yaml:"servicecatalogappregistry"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#servicediscovery AwsProvider#servicediscovery}
	Servicediscovery *string `field:"optional" json:"servicediscovery" yaml:"servicediscovery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#servicequotas AwsProvider#servicequotas}
	Servicequotas *string `field:"optional" json:"servicequotas" yaml:"servicequotas"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ses AwsProvider#ses}
	Ses *string `field:"optional" json:"ses" yaml:"ses"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sesv2 AwsProvider#sesv2}
	Sesv2 *string `field:"optional" json:"sesv2" yaml:"sesv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sfn AwsProvider#sfn}
	Sfn *string `field:"optional" json:"sfn" yaml:"sfn"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#shield AwsProvider#shield}
	Shield *string `field:"optional" json:"shield" yaml:"shield"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#signer AwsProvider#signer}
	Signer *string `field:"optional" json:"signer" yaml:"signer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sns AwsProvider#sns}
	Sns *string `field:"optional" json:"sns" yaml:"sns"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sqs AwsProvider#sqs}
	Sqs *string `field:"optional" json:"sqs" yaml:"sqs"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssm AwsProvider#ssm}
	Ssm *string `field:"optional" json:"ssm" yaml:"ssm"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssmcontacts AwsProvider#ssmcontacts}
	Ssmcontacts *string `field:"optional" json:"ssmcontacts" yaml:"ssmcontacts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssmincidents AwsProvider#ssmincidents}
	Ssmincidents *string `field:"optional" json:"ssmincidents" yaml:"ssmincidents"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssmquicksetup AwsProvider#ssmquicksetup}
	Ssmquicksetup *string `field:"optional" json:"ssmquicksetup" yaml:"ssmquicksetup"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssmsap AwsProvider#ssmsap}
	Ssmsap *string `field:"optional" json:"ssmsap" yaml:"ssmsap"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sso AwsProvider#sso}
	Sso *string `field:"optional" json:"sso" yaml:"sso"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#ssoadmin AwsProvider#ssoadmin}
	Ssoadmin *string `field:"optional" json:"ssoadmin" yaml:"ssoadmin"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#stepfunctions AwsProvider#stepfunctions}
	Stepfunctions *string `field:"optional" json:"stepfunctions" yaml:"stepfunctions"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#storagegateway AwsProvider#storagegateway}
	Storagegateway *string `field:"optional" json:"storagegateway" yaml:"storagegateway"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#sts AwsProvider#sts}
	Sts *string `field:"optional" json:"sts" yaml:"sts"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#swf AwsProvider#swf}
	Swf *string `field:"optional" json:"swf" yaml:"swf"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#synthetics AwsProvider#synthetics}
	Synthetics *string `field:"optional" json:"synthetics" yaml:"synthetics"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#taxsettings AwsProvider#taxsettings}
	Taxsettings *string `field:"optional" json:"taxsettings" yaml:"taxsettings"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#timestreaminfluxdb AwsProvider#timestreaminfluxdb}
	Timestreaminfluxdb *string `field:"optional" json:"timestreaminfluxdb" yaml:"timestreaminfluxdb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#timestreamquery AwsProvider#timestreamquery}
	Timestreamquery *string `field:"optional" json:"timestreamquery" yaml:"timestreamquery"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#timestreamwrite AwsProvider#timestreamwrite}
	Timestreamwrite *string `field:"optional" json:"timestreamwrite" yaml:"timestreamwrite"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#transcribe AwsProvider#transcribe}
	Transcribe *string `field:"optional" json:"transcribe" yaml:"transcribe"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#transcribeservice AwsProvider#transcribeservice}
	Transcribeservice *string `field:"optional" json:"transcribeservice" yaml:"transcribeservice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#transfer AwsProvider#transfer}
	Transfer *string `field:"optional" json:"transfer" yaml:"transfer"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#verifiedpermissions AwsProvider#verifiedpermissions}
	Verifiedpermissions *string `field:"optional" json:"verifiedpermissions" yaml:"verifiedpermissions"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#vpclattice AwsProvider#vpclattice}
	Vpclattice *string `field:"optional" json:"vpclattice" yaml:"vpclattice"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#waf AwsProvider#waf}
	Waf *string `field:"optional" json:"waf" yaml:"waf"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#wafregional AwsProvider#wafregional}
	Wafregional *string `field:"optional" json:"wafregional" yaml:"wafregional"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#wafv2 AwsProvider#wafv2}
	Wafv2 *string `field:"optional" json:"wafv2" yaml:"wafv2"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#wellarchitected AwsProvider#wellarchitected}
	Wellarchitected *string `field:"optional" json:"wellarchitected" yaml:"wellarchitected"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#workspaces AwsProvider#workspaces}
	Workspaces *string `field:"optional" json:"workspaces" yaml:"workspaces"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#workspacesweb AwsProvider#workspacesweb}
	Workspacesweb *string `field:"optional" json:"workspacesweb" yaml:"workspacesweb"`
	// Use this to override the default service endpoint URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs#xray AwsProvider#xray}
	Xray *string `field:"optional" json:"xray" yaml:"xray"`
}

