// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersSelfManagedKafkaParametersCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/pipes_pipe#basic_auth PipesPipe#basic_auth}.
	BasicAuth *string `field:"required" json:"basicAuth" yaml:"basicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/pipes_pipe#client_certificate_tls_auth PipesPipe#client_certificate_tls_auth}.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/pipes_pipe#sasl_scram_256_auth PipesPipe#sasl_scram_256_auth}.
	SaslScram256Auth *string `field:"optional" json:"saslScram256Auth" yaml:"saslScram256Auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/pipes_pipe#sasl_scram_512_auth PipesPipe#sasl_scram_512_auth}.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

