package ivsplaybackkeypair


type IvsPlaybackKeyPairTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_playback_key_pair#create IvsPlaybackKeyPair#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/ivs_playback_key_pair#delete IvsPlaybackKeyPair#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

