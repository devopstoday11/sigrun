package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/devopstoday11/sigrun/pkg/util"
	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
)

const (
	CONFIG_MODE_KEYPAIR = "keypair"
	CONFIG_MODE_KEYLESS = "keyless"
)

func NewKeypairConfig(name, pubKey, privKey string, images []string) *KeyPair {
	return &KeyPair{
		Name:       name,
		Mode:       CONFIG_MODE_KEYPAIR,
		ChainNo:    0,
		PublicKey:  pubKey,
		PrivateKey: privKey,
		Images:     images,
		Signature:  "",
	}
}

func NewKeylessConfig(name string, maintainers, images []string) *Keyless {
	return &Keyless{
		Name:        name,
		Mode:        CONFIG_MODE_KEYLESS,
		ChainNo:     0,
		Maintainers: maintainers,
		Images:      images,
		Signature:   "",
	}
}

type Config interface {
	InitializeRepository() error
	SignImages() error
	CommitRepositoryUpdate() error
	GetChainNo() int64
	Sign([]byte) (string, error)
	SignDoc() ([]byte, error)
	Validate() error
	GetVerificationInfo() *VerificationInfo
	VerifySuccessorConfig(Config) error
	GetSignature() string
	VerifyImage(image string) error
}

type VerificationInfo struct {
	Name        string
	Mode        string
	ChainNo     int64
	PublicKey   string
	Maintainers []string
	Images      []string
}

func ReadRepositoryConfig() (Config, error) {
	encodedConfig, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		return nil, err
	}

	return parseConfig(encodedConfig)
}

func GetGUID(path string) (string, error) {
	genesisConfPath := strings.Replace(path, FILE_NAME, ".sigrun/0.json", -1)

	resp, err := http.Get(genesisConfPath)
	if err != nil {
		return "", err
	}

	return util.SHA256Hash(resp.Body)
}

// TODO should be repo urls, currentl config file urls
func ReadRepos(repoUrls ...string) (map[string]Config, error) {
	pathToConfig := make(map[string]Config)
	for _, path := range repoUrls {
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}

		confRaw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		pathToConfig[path], err = parseConfig(confRaw)
		if err != nil {
			return nil, err
		}
	}

	return pathToConfig, nil
}

func GetVerificationConfigFromVerificationInfo(info *VerificationInfo) Config {
	if info.Mode == CONFIG_MODE_KEYLESS {
		return &Keyless{
			Name:        info.Name,
			Mode:        info.Mode,
			ChainNo:     info.ChainNo,
			Maintainers: info.Maintainers,
			Images:      info.Images,
			Signature:   "",
		}
	} else {
		return &KeyPair{
			Name:       info.Name,
			Mode:       info.Mode,
			ChainNo:    info.ChainNo,
			PublicKey:  info.PublicKey,
			PrivateKey: "",
			Images:     info.Images,
			Signature:  "",
		}
	}
}

func VerifyChain(repoPath string, oldConf, newConf Config) error {
	var err error

	currentChainNo := oldConf.GetChainNo() + 1
	prevConf := oldConf
	var currConf Config
	for currentChainNo <= newConf.GetChainNo() {
		currPath := strings.Replace(repoPath, FILE_NAME, ".sigrun/"+fmt.Sprint(currentChainNo)+".json", -1)
		confMap, err := ReadRepos(currPath)
		if err != nil {
			return err
		}
		currConf = confMap[currPath]

		err = prevConf.VerifySuccessorConfig(currConf)
		if err != nil {
			return err
		}
		prevConf = currConf
		currentChainNo = currConf.GetChainNo() + 1
	}

	isSame, err := isSame(currConf, newConf)
	if err != nil {
		return err
	}

	if !isSame {
		return fmt.Errorf("chain head is not the same as config file")
	}

	return nil
}

func NormalizeImageName(image string) (string, error) {
	imgInfo, err := newImageInfo(image)
	if err != nil {
		return "", err
	}

	return imgInfo.String(), nil
}

func newImageInfo(image string) (*ImageInfo, error) {
	image = addDefaultDomain(image)
	ref, err := reference.Parse(image)
	if err != nil {
		return nil, errors.Wrapf(err, "bad image: %s", image)
	}

	var registry, path, name, tag, digest string
	if named, ok := ref.(reference.Named); ok {
		registry = reference.Domain(named)
		path = reference.Path(named)
		name = path[strings.LastIndex(path, "/")+1:]
	}

	if tagged, ok := ref.(reference.Tagged); ok {
		tag = tagged.Tag()
	}

	if digested, ok := ref.(reference.Digested); ok {
		digest = digested.Digest().String()
	}

	// set default tag - the domain is set via addDefaultDomain before parsing
	if tag == "" {
		tag = "latest"
	}

	return &ImageInfo{
		Registry: registry,
		Name:     name,
		Path:     path,
		Tag:      tag,
		Digest:   digest,
	}, nil
}

type ImageInfo struct {

	// Registry is the URL address of the image registry e.g. `docker.io`
	Registry string `json:"registry,omitempty"`

	// Name is the image name portion e.g. `busybox`
	Name string `json:"name"`

	// Path is the repository path and image name e.g. `some-repository/busybox`
	Path string `json:"path"`

	// Tag is the image tag e.g. `v2`
	Tag string `json:"tag,omitempty"`

	// Digest is the image digest portion e.g. `sha256:128c6e3534b842a2eec139999b8ce8aa9a2af9907e2b9269550809d18cd832a3`
	Digest string `json:"digest,omitempty"`
}

func (i *ImageInfo) String() string {
	image := i.Registry + "/" + i.Path + ":" + i.Tag
	if i.Digest != "" {
		image = image + "@" + i.Digest
	}

	return image
}

func addDefaultDomain(name string) string {
	i := strings.IndexRune(name, '/')
	if i == -1 || (!strings.ContainsAny(name[:i], ".:") && name[:i] != "localhost" && strings.ToLower(name[:i]) == name[:i]) {
		return "docker.io/" + name
	}

	return name
}
