// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrcontainersjobtemplate


type EmrcontainersJobTemplateJobTemplateDataJobDriver struct {
	// spark_sql_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/emrcontainers_job_template#spark_sql_job_driver EmrcontainersJobTemplate#spark_sql_job_driver}
	SparkSqlJobDriver *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSqlJobDriver `field:"optional" json:"sparkSqlJobDriver" yaml:"sparkSqlJobDriver"`
	// spark_submit_job_driver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/emrcontainers_job_template#spark_submit_job_driver EmrcontainersJobTemplate#spark_submit_job_driver}
	SparkSubmitJobDriver *EmrcontainersJobTemplateJobTemplateDataJobDriverSparkSubmitJobDriver `field:"optional" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}

