module github.com/devopstoday11/sigrun

go 1.16

require (
	github.com/docker/distribution v2.8.1+incompatible
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-git/go-git/v5 v5.2.0
	github.com/google/go-containerregistry v0.14.0
	github.com/pkg/errors v0.9.1
	github.com/sigstore/cosign v1.0.1
	github.com/sigstore/fulcio v0.1.1
	github.com/sigstore/rekor v1.2.0
	github.com/sigstore/sigstore v1.6.4
	github.com/spf13/cobra v1.7.0
	github.com/tidwall/pretty v1.2.0
	k8s.io/api v0.26.1
	k8s.io/apimachinery v0.26.1
	k8s.io/cli-runtime v0.22.0
	k8s.io/client-go v0.26.1
)

replace github.com/gorilla/rpc v1.2.0+incompatible => github.com/gorilla/rpc v1.2.0
