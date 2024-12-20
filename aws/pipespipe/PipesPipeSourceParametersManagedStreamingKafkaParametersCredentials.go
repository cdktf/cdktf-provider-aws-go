// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersManagedStreamingKafkaParametersCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/pipes_pipe#client_certificate_tls_auth PipesPipe#client_certificate_tls_auth}.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/pipes_pipe#sasl_scram_512_auth PipesPipe#sasl_scram_512_auth}.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

