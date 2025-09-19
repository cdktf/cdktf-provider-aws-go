// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectArtifacts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#artifact_identifier CodebuildProject#artifact_identifier}.
	ArtifactIdentifier *string `field:"optional" json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#bucket_owner_access CodebuildProject#bucket_owner_access}.
	BucketOwnerAccess *string `field:"optional" json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#encryption_disabled CodebuildProject#encryption_disabled}.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#location CodebuildProject#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#name CodebuildProject#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#namespace_type CodebuildProject#namespace_type}.
	NamespaceType *string `field:"optional" json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#override_artifact_name CodebuildProject#override_artifact_name}.
	OverrideArtifactName interface{} `field:"optional" json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#packaging CodebuildProject#packaging}.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/codebuild_project#path CodebuildProject#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

