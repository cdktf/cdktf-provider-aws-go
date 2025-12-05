// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/emrcontainers_job_template#entry_point EmrcontainersJobTemplate#entry_point}.
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/emrcontainers_job_template#spark_sql_parameters EmrcontainersJobTemplate#spark_sql_parameters}.
	SparkSqlParameters *string `field:"optional" json:"sparkSqlParameters" yaml:"sparkSqlParameters"`
}

