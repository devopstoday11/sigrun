package config

type KeylessConfig struct {
	Name        string
	Mode        string
	ChainNo     int64
	Maintainers []string
	Images      []string
	Signature   string
}

func (conf *KeylessConfig) GetVerificationInfo() *VerificationInfo {
	panic("implement me")
}

func (conf *KeylessConfig) VerifySuccessorConfig(config Config) error {
	panic("implement me")
}

func (conf *KeylessConfig) GetSignature() string {
	panic("implement me")
}

func (conf *KeylessConfig) InitializeRepository() error {
	panic("implement me")
}

func (conf *KeylessConfig) SignImages() error {
	panic("implement me")
}

func (conf *KeylessConfig) CommitRepositoryUpdate() error {
	panic("implement me")
}

func (conf *KeylessConfig) GetChainNo() int64 {
	panic("implement me")
}

func (conf *KeylessConfig) Sign(bytes []byte) (string, error) {
	panic("implement me")
}

func (conf *KeylessConfig) SignDoc() ([]byte, error) {
	panic("implement me")
}

func (conf *KeylessConfig) Validate() error {
	panic("implement me")
}