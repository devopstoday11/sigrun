module github.com/devopstoday11/sigrun

go 1.16

require (
	github.com/docker/distribution v2.8.2+incompatible
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-git/go-git/v5 v5.2.0
	github.com/google/go-containerregistry v0.6.0
	github.com/gopherjs/gopherjs v0.0.0-20190328170749-bb2674552d8f // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sigstore/cosign v1.0.1
	github.com/sigstore/fulcio v0.1.1
	github.com/sigstore/rekor v0.3.0
	github.com/sigstore/sigstore v0.0.0-20210729211320-56a91f560f44
	github.com/spf13/cobra v1.2.1
	github.com/tidwall/pretty v1.2.0
	k8s.io/api v0.22.0
	k8s.io/apimachinery v0.22.0
	k8s.io/cli-runtime v0.22.0
	k8s.io/client-go v0.22.0
)

replace github.com/gorilla/rpc v1.2.0+incompatible => github.com/gorilla/rpc v1.2.0
