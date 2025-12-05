// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer


type CloudwatchLogTransformerTransformerConfig struct {
	// add_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#add_keys CloudwatchLogTransformer#add_keys}
	AddKeys interface{} `field:"optional" json:"addKeys" yaml:"addKeys"`
	// copy_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#copy_value CloudwatchLogTransformer#copy_value}
	CopyValue interface{} `field:"optional" json:"copyValue" yaml:"copyValue"`
	// csv block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#csv CloudwatchLogTransformer#csv}
	Csv interface{} `field:"optional" json:"csv" yaml:"csv"`
	// date_time_converter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#date_time_converter CloudwatchLogTransformer#date_time_converter}
	DateTimeConverter interface{} `field:"optional" json:"dateTimeConverter" yaml:"dateTimeConverter"`
	// delete_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#delete_keys CloudwatchLogTransformer#delete_keys}
	DeleteKeys interface{} `field:"optional" json:"deleteKeys" yaml:"deleteKeys"`
	// grok block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#grok CloudwatchLogTransformer#grok}
	Grok interface{} `field:"optional" json:"grok" yaml:"grok"`
	// list_to_map block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#list_to_map CloudwatchLogTransformer#list_to_map}
	ListToMap interface{} `field:"optional" json:"listToMap" yaml:"listToMap"`
	// lower_case_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#lower_case_string CloudwatchLogTransformer#lower_case_string}
	LowerCaseString interface{} `field:"optional" json:"lowerCaseString" yaml:"lowerCaseString"`
	// move_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#move_keys CloudwatchLogTransformer#move_keys}
	MoveKeys interface{} `field:"optional" json:"moveKeys" yaml:"moveKeys"`
	// parse_cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_cloudfront CloudwatchLogTransformer#parse_cloudfront}
	ParseCloudfront interface{} `field:"optional" json:"parseCloudfront" yaml:"parseCloudfront"`
	// parse_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_json CloudwatchLogTransformer#parse_json}
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// parse_key_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_key_value CloudwatchLogTransformer#parse_key_value}
	ParseKeyValue interface{} `field:"optional" json:"parseKeyValue" yaml:"parseKeyValue"`
	// parse_postgres block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_postgres CloudwatchLogTransformer#parse_postgres}
	ParsePostgres interface{} `field:"optional" json:"parsePostgres" yaml:"parsePostgres"`
	// parse_route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_route53 CloudwatchLogTransformer#parse_route53}
	ParseRoute53 interface{} `field:"optional" json:"parseRoute53" yaml:"parseRoute53"`
	// parse_to_ocsf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_to_ocsf CloudwatchLogTransformer#parse_to_ocsf}
	ParseToOcsf interface{} `field:"optional" json:"parseToOcsf" yaml:"parseToOcsf"`
	// parse_vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_vpc CloudwatchLogTransformer#parse_vpc}
	ParseVpc interface{} `field:"optional" json:"parseVpc" yaml:"parseVpc"`
	// parse_waf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#parse_waf CloudwatchLogTransformer#parse_waf}
	ParseWaf interface{} `field:"optional" json:"parseWaf" yaml:"parseWaf"`
	// rename_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#rename_keys CloudwatchLogTransformer#rename_keys}
	RenameKeys interface{} `field:"optional" json:"renameKeys" yaml:"renameKeys"`
	// split_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#split_string CloudwatchLogTransformer#split_string}
	SplitString interface{} `field:"optional" json:"splitString" yaml:"splitString"`
	// substitute_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#substitute_string CloudwatchLogTransformer#substitute_string}
	SubstituteString interface{} `field:"optional" json:"substituteString" yaml:"substituteString"`
	// trim_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#trim_string CloudwatchLogTransformer#trim_string}
	TrimString interface{} `field:"optional" json:"trimString" yaml:"trimString"`
	// type_converter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#type_converter CloudwatchLogTransformer#type_converter}
	TypeConverter interface{} `field:"optional" json:"typeConverter" yaml:"typeConverter"`
	// upper_case_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/cloudwatch_log_transformer#upper_case_string CloudwatchLogTransformer#upper_case_string}
	UpperCaseString interface{} `field:"optional" json:"upperCaseString" yaml:"upperCaseString"`
}

