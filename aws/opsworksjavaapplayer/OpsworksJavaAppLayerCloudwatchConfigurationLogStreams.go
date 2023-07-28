package opsworksjavaapplayer


type OpsworksJavaAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#file OpsworksJavaAppLayer#file}.
	File *string `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#log_group_name OpsworksJavaAppLayer#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#batch_count OpsworksJavaAppLayer#batch_count}.
	BatchCount *float64 `field:"optional" json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#batch_size OpsworksJavaAppLayer#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#buffer_duration OpsworksJavaAppLayer#buffer_duration}.
	BufferDuration *float64 `field:"optional" json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#datetime_format OpsworksJavaAppLayer#datetime_format}.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#encoding OpsworksJavaAppLayer#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#file_fingerprint_lines OpsworksJavaAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `field:"optional" json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#initial_position OpsworksJavaAppLayer#initial_position}.
	InitialPosition *string `field:"optional" json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#multiline_start_pattern OpsworksJavaAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `field:"optional" json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/opsworks_java_app_layer#time_zone OpsworksJavaAppLayer#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

