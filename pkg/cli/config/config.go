package config

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/sigstore/cosign/pkg/cosign"
)

type Config struct {
	PublicKey  string
	PrivateKey string
	Images     []string
	Signature  string
}

const CONFIG_FILE_NAME = "sigrun-config.json"

func Read() (*Config, error) {
	configF, err := os.Open(CONFIG_FILE_NAME)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = json.NewDecoder(configF).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func Create(conf *Config, password string) error {
	configF, err := os.Create(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}
	conf.Signature = ""

	confRaw, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	signer, err := cosign.LoadECDSAPrivateKey([]byte(conf.PrivateKey), []byte(password))
	if err != nil {
		return err
	}

	sig, err := signer.SignMessage(bytes.NewReader(confRaw))
	if err != nil {
		return err
	}

	conf.Signature = string(sig)

	err = json.NewEncoder(configF).Encode(conf)
	if err != nil {
		return err
	}

	return nil
}
