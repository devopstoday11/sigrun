<h1 align="left">SigRun 
<a href="https://www.producthunt.com/posts/awesome-github-profiles?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-github-profiles" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=277987&theme=light" alt="Awesome GitHub Profiles - Best curated list of developers readme, updated every 15 min | Product Hunt" style="width: 200px; height: 44px;" width="200" height="44" /></a></h1>


Sign your artifacts source code or container images using Sigstore chain of tools & Known Container Image Build tools, Save the Signatures you want to use within your Infra, and Validate &amp; Control the deployments to allow only the known Signatures.
> What's with the Name (in case if you are curious)?
> You can think of multiple ways. It has a flexible interpretation, like Signatures for Runtime or Runtime Signatures. Whatever you want to imagine! :smiley: 
#

### Install
```
go install cmd/sigrun/kubectl-sigrun.go
```

### Usage
```
kubectl sigrun --help
```

##### Purpose:
To make it easy to use SigStore chain of tools. Make the Supply Chain Security for Software adoption easy. 
#
##### Usage feasibility:
Local, CI/CD pipelines, K8s Clusters, VMs. 
#
#### Features:
- Using Sigstore tools in your Infra for Air-Gap offline usage via your CI/CD Pipeline
- Sign your artifacts
- Private & Public key-pair
- Keyless
- Save your artifacts signatures to certain storage
- Save your container image signatures to certain storage
- Validate Signatures using Storage location of Signatures
- Control deployments to allow only known Signatures using our Custom Admission Controller or OPA/Kyverno/Gatekeeper
- Vault Integration to save Keys
- CI/CD Tools integration
- Integration with tools like Buildpacks, Buildah, Source2Image, Kaniko, Skaffold, Docker Build, Podman, etc. 
- OIDC
- Vulnerability Scanning of your container images
- Integrate with Non-Profit SigStore public service


#
