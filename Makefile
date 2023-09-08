
# Image URL to use all building/pushing image targets
ONMETAL_API_NET_IMG ?= onmetal-api-net:latest
APINETLET_IMG ?= apinetlet:latest
METALNETLET_IMG ?= metalnetlet:latest
KIND_CLUSTER_NAME ?= kind

# ENVTEST_K8S_VERSION refers to the version of kubebuilder assets to be downloaded by envtest binary.
ENVTEST_K8S_VERSION = 1.27.1

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

SSH_KEY ?= ${HOME}/.ssh/id_rsa

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate rbac objects.
	# onmetal-api-net
	$(CONTROLLER_GEN) rbac:roleName=manager-role paths="./internal/controllers/..." output:rbac:artifacts:config=config/onmetal-api-net/rbac

	# apinetlet
	$(CONTROLLER_GEN) rbac:roleName=manager-role paths="./apinetlet/controllers/..." output:rbac:artifacts:config=config/apinetlet/rbac
	CONTROLLER_GEN=$(CONTROLLER_GEN) ./hack/cluster-controller-gen.sh cluster=apinet rbac:roleName=apinet-role paths="./apinetlet/controllers/..." output:rbac:artifacts:config=config/apinetlet/apinet-rbac

	# metalnetlet
	$(CONTROLLER_GEN) rbac:roleName=manager-role paths="./metalnetlet/controllers/..." output:rbac:artifacts:config=config/metalnetlet/rbac
	CONTROLLER_GEN=$(CONTROLLER_GEN) ./hack/cluster-controller-gen.sh cluster=metalnet rbac:roleName=metalnet-role paths="./metalnetlet/controllers/..." output:rbac:artifacts:config=config/metalnetlet/metalnet-rbac

	# Promote *let roles.
	./hack/promote-let-role.sh config/apinetlet/apinet-rbac/role.yaml config/onmetal-api-net/rbac/apinetlet_role.yaml apinet.api.onmetal.de:system:apinetlets
	./hack/promote-let-role.sh config/metalnetlet/metalnet-rbac/role.yaml config/onmetal-api-net/rbac/metalnetlet_role.yaml apinet.api.onmetal.de:system:metalnetlet

.PHONY: generate
generate: vgopath models-schema deepcopy-gen client-gen lister-gen informer-gen defaulter-gen conversion-gen openapi-gen applyconfiguration-gen
	VGOPATH=$(VGOPATH) \
	MODELS_SCHEMA=$(MODELS_SCHEMA) \
	DEEPCOPY_GEN=$(DEEPCOPY_GEN) \
	CLIENT_GEN=$(CLIENT_GEN) \
	LISTER_GEN=$(LISTER_GEN) \
	INFORMER_GEN=$(INFORMER_GEN) \
	DEFAULTER_GEN=$(DEFAULTER_GEN) \
	CONVERSION_GEN=$(CONVERSION_GEN) \
	OPENAPI_GEN=$(OPENAPI_GEN) \
	APPLYCONFIGURATION_GEN=$(APPLYCONFIGURATION_GEN) \
	./hack/update-codegen.sh

.PHONY: add-license
add-license: addlicense ## Add license headers to all go files.
	find . -name '*.go' -exec $(ADDLICENSE) -c 'OnMetal authors' {} +

.PHONY: fmt
fmt: goimports ## Run goimports against code.
	$(GOIMPORTS) -w .

.PHONY: check-license
check-license: addlicense ## Check that every file has a license header present.
	find . -name '*.go' -exec $(ADDLICENSE) -check -c 'OnMetal authors' {} +

.PHONY: lint
lint: ## Run golangci-lint against code.
	golangci-lint run ./...

.PHONY: clean
clean: ## Clean any artifacts that can be regenerated.
	rm -rf client-go/applyconfigurations
	rm -rf client-go/informers
	rm -rf client-go/listers
	rm -rf client-go/onmetalapi
	rm -rf client-go/openapi

.PHONY: check
check: generate manifests add-license fmt lint test ## Lint and run tests.

ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
.PHONY: test
test: envtest generate fmt check-license ## Run tests.
	KUBEBUILDER_ASSETS="$(shell $(ENVTEST) use $(ENVTEST_K8S_VERSION) -p path)" go test ./... -coverprofile cover.out

##@ Build

.PHONY: build-onmetal-api-net
build-onmetal-api-net: generate fmt addlicense lint ## Build onmetal-api-net binary.
	go build -o bin/manager ./onmetal-api-net/main.go

.PHONY: build-apinetlet
build-apinetlet: generate fmt addlicense lint ## Build apinetlet.
	go build -o bin/apinetlet ./apinetlet/main.go

.PHONY: build-metalnetlet
build-metalnetlet: generate fmt addlicense lint ## Build metalnetlet.
	go build -o bin/metalnetlet ./metalnetlet/main.go

.PHONY: build
build: build-onmetal-api-net build-apinetlet build-metalnetlet ## Build onmetal-api-net, apinetlet, metalnetlet.

.PHONY: run-onmetal-api-net
run-onmetal-api-net: manifests generate fmt lint ## Run a onmetal-api-net from your host.
	go run ./onmetal-api-net/main.go

.PHONY: run-apinetlet
run-apinetlet: manifests generate fmt lint ## Run apinetlet from your host.
	go run ./apinetlet/main.go

.PHONY: run-metalnetlet
run-metalnetlet: manifests generate fmt lint ## Run metalnetlet from your host.
	go run ./metalnetlet/main.go

.PHONY: docker-build-onmetal-api-net
docker-build-onmetal-api-net: ## Build onmetal-api-net image with the manager.
	docker build --ssh default=${SSH_KEY} --target onmetal-api-net-manager -t ${ONMETAL_API_NET_IMG} .

.PHONY: docker-build-apinetlet
docker-build-apinetlet: ## Build apinetlet image with the manager.
	docker build --ssh default=${SSH_KEY} --target apinetlet-manager -t ${APINETLET_IMG} .

.PHONY: docker-build-metalnetlet
docker-build-metalnetlet: ## Build metalnetlet image with the manager.
	docker build --ssh default=${SSH_KEY} --target metalnetlet-manager -t ${METALNETLET_IMG} .

.PHONY: docker-build
docker-build: docker-build-onmetal-api-net docker-build-apinetlet docker-build-metalnetlet ## Build docker images.

.PHONY: docker-push-onmetal-api-net
docker-push-onmetal-api-net: ## Push onmetal-api-net image.
	docker push ${ONMETAL_API_NET_IMG}

.PHONY: docker-push-apinetlet
docker-push-apinetlet: ## Push apinetlet image.
	docker push ${APINETLET_IMG}

.PHONY: docker-push-metalnetlet
docker-push-metalnetlet: ## Push metalnetlet image.
	docker push ${METALNETLET_IMG}

.PHONY: docker-push
docker-push: docker-push-onmetal-api-net docker-build-apinetlet docker-build-metalnetlet ## Push onmetal-api-net, apinetlet, metalnetlet image.

##@ Deployment

.PHONY: install-onmetal-api-net
install-onmetal-api-net: manifests ## Install onmetal-api-net CRDs into the K8s cluster specified in ~/.kube/config.
	kubectl apply -k config/onmetal-api-net/crd

.PHONY: uninstall-onmetal-api-net
uninstall-onmetal-api-net: manifests ## Uninstall onmetal-api-net CRDs from the K8s cluster specified in ~/.kube/config.
	kubectl delete-k config/onmetal-api-net/crd

.PHONY: install-apinetlet
install-apinetlet: manifests ## Install apinetlet CRDs into the K8s cluster specified in ~/.kube/config.
	kubectl apply -k config/apinetlet/crd

.PHONY: uninstall-apinetlet
uninstall-apinetlet: manifests ## Uninstall apinetlet CRDs from the K8s cluster specified in ~/.kube/config.
	kubectl delete-k config/apinetlet/crd

.PHONY: install-metalnetlet
install-metalnetlet: manifests ## Install metalnetlet CRDs into the K8s cluster specified in ~/.kube/config.
	kubectl apply -k config/metalnetlet/crd

.PHONY: uninstall-metalnetlet
uninstall-metalnetlet: manifests ## Uninstall metalnetlet CRDs from the K8s cluster specified in ~/.kube/config.
	kubectl delete-k config/metalnetlet/crd

.PHONY: install
install: install-onmetal-api-net install-apinetlet install-metalnetlet ## Uninstall onmetal-api-net, apinetlet, metalnetlet.

.PHONY: uninstall
uninstall: uninstall-onmetal-api-net uninstall-apinetlet uninstall-metalnetlet ## Uninstall onmetal-api-net, apinetlet, metalnetlet.

.PHONY: deploy-onmetal-api-net
deploy-onmetal-api-net: manifests kustomize ## Deploy onmetal-api-net controller to the K8s cluster specified in ~/.kube/config.
	cd config/onmetal-api-net/manager && $(KUSTOMIZE) edit set image controller=${ONMETAL_API_NET_IMG}
	kubectl apply -k config/onmetal-api-net/default

.PHONY: deploy-apinetlet
deploy-apinetlet: manifests kustomize ## Deploy apinetlet controller to the K8s cluster specified in ~/.kube/config.
	cd config/apinetlet/manager && $(KUSTOMIZE) edit set image controller=${APINETLET_IMG}
	kubectl apply -k config/apinetlet/default

.PHONY: deploy-metalnetlet
deploy-metalnetlet: manifests kustomize ## Deploy metalnetlet controller to the K8s cluster specified in ~/.kube/config.
	cd config/metalnetlet/manager && $(KUSTOMIZE) edit set image controller=${METALNETLET_IMG}
	kubectl apply -k config/metalnetlet/default

.PHONY: deploy
deploy: deploy-onmetal-api-net deploy-apinetlet deploy-metalnetlet ## Deploy onmetal-api-net, apinetlet, metalnetlet

.PHONY: undeploy-onmetal-api-net
undeploy-onmetal-api-net: ## Undeploy onmetal-api-net controller from the K8s cluster specified in ~/.kube/config.
	kubectl delete -k config/onmetal-api-net

.PHONY: undeploy-apinetlet
undeploy-apinetlet: ## Undeploy apinetlet controller from the K8s cluster specified in ~/.kube/config.
	kubectl delete -k config/apinetlet

.PHONY: undeploy-metalnetlet
undeploy-metalnetlet: ## Undeploy metalnetlet controller from the K8s cluster specified in ~/.kube/config.
	kubectl delete -k config/metalnetlet

.PHONY: undeploy
undeploy: undeploy-onmetal-api-net undeploy-apinetlet undeploy-metalnetlet ## Undeploy onmetal-api-net, apinetlet, metalnetlet controller from the K8s cluster specified in ~/.kube/config.

##@ Kind Deployment plumbing

.PHONY: kind-load-onmetal-api-net
kind-load-onmetal-api-net: docker-build-onmetal-api-net ## Load onmetal-api-net image to kind cluster.
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${ONMETAL_API_NET_IMG}

.PHONY: kind-load-apinetlet
kind-load-apinetlet: docker-build-apinetlet ## Load apinetlet image to kind cluster.
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${APINETLET_IMG}

.PHONY: kind-load-metalnetlet
kind-load-metalnetlet: docker-build-metalnetlet ## Load metalnetlet image to kind cluster.
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${METALNETLET_IMG}

.PHONY: kind-load
kind-load: kind-load-onmetal-api-net kind-load-apinetlet kind-load-metalnetlet ## Load onmetal-api-net, apinetlet, metalnetlet image to kind cluster.

.PHONY: kind-restart-onmetal-api-net
kind-restart-onmetal-api-net: ## Restarts the onmetal-api-net controller manager.
	kubectl -n onmetal-api-net-system delete rs -l control-plane=controller-manager

.PHONY: kind-restart-apinetlet
kind-restart-apinetlet: ## Restarts the apinetlet controller manager.
	kubectl -n apinetlet-system delete rs -l control-plane=controller-manager

.PHONY: kind-restart-metalnetlet
kind-restart-metalnetlet: ## Restarts the metalnetlet controller manager.
	kubectl -n metalnetlet-system delete rs -l control-plane=controller-manager

.PHONY: kind-restart
kind-restart: kind-restart-onmetal-api-net kind-restart-apinetlet kind-restart-metalnetlet ## Restarts the onmetal-api-net, apinetlet, metalnetlet controller manager.

.PHONY: kind-build-load-restart-onmetal-api-net
kind-build-load-restart-onmetal-api-net: docker-build-onmetal-api-net kind-load-onmetal-api-net kind-restart-onmetal-api-net ## Build, load and restart onmetal-api-net.

.PHONY: kind-build-load-restart-apinetlet
kind-build-load-restart-apinetlet: docker-build-apinetlet kind-load-apinetlet kind-restart-apinetlet ## Build, load and restart apinetlet.

.PHONY: kind-build-load-restart-metalnetlet
kind-build-load-restart-metalnetlet: docker-build-metalnetlet kind-load-metalnetlet kind-restart-metalnetlet ## Build, load and restart metalnetlet.

.PHONY: kind-build-load-restart
kind-build-load-restart: kind-build-load-restart-onmetal-api-net kind-build-load-restart-apinetlet kind-build-load-restart-metalnetlet

.PHONY: kind-apply-onmetal-api-net
kind-apply-onmetal-api-net: manifests ## Apply onmetal-api-net to the cluster specified in ~/.kube/config.
	kubectl apply -k config/onmetal-api-net/kind

.PHONY: kind-apply-apinetlet
kind-apply-apinetlet: manifests ## Apply apinetlet to the cluster specified in ~/.kube/config.
	kubectl apply -k config/apinetlet/kind

.PHONY: kind-apply-metalnetlet
kind-apply-metalnetlet: manifests ## Apply metalnetlet to the cluster specified in ~/.kube/config.
	kubectl apply -k config/metalnetlet/kind

.PHONY: kind-apply
kind-apply: kind-apply-onmetal-api-net kind-apply-apinetlet kind-apply-metalnetlet ## Apply onmetal-api-net, apinetlet, metalnetlet to the cluster specified in ~/.kube/config.

.PHONY: kind-deploy
kind-deploy: kind-build-load-restart kind-apply ## Build, load and restart onmetal-api-net and apinetlet and apply them.

.PHONY: kind-delete-onmetal-api-net
kind-delete-onmetal-api-net: ## Delete onmetal-api-net from the cluster specified in ~/.kube/config.
	kubectl delete -k config/onmetal-api-net/kind

.PHONY: kind-delete-apinetlet
kind-delete-apinetlet: ## Delete apinetlet from the cluster specified in ~/.kube/config.
	kubectl delete -k config/apinetlet/kind

.PHONY: kind-delete-metalnetlet
kind-delete-metalnetlet: ## Delete metalnetlet from the cluster specified in ~/.kube/config.
	kubectl delete -k config/metalnetlet/kind

.PHONY: kind-delete
kind-delete: kind-delete-onmetal-api-net kind-delete-apinetlet kind-delete-metalnetlet ## Delete onmetal-api-net, apinetlet, metalnetlet from the cluster specified in ~/.kube/config.

##@ Tools

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
KUSTOMIZE ?= $(LOCALBIN)/kustomize
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest
OPENAPI_EXTRACTOR ?= $(LOCALBIN)/openapi-extractor
DEEPCOPY_GEN ?= $(LOCALBIN)/deepcopy-gen
CLIENT_GEN ?= $(LOCALBIN)/client-gen
LISTER_GEN ?= $(LOCALBIN)/lister-gen
INFORMER_GEN ?= $(LOCALBIN)/informer-gen
DEFAULTER_GEN ?= $(LOCALBIN)/defaulter-gen
CONVERSION_GEN ?= $(LOCALBIN)/conversion-gen
OPENAPI_GEN ?= $(LOCALBIN)/openapi-gen
APPLYCONFIGURATION_GEN ?= $(LOCALBIN)/applyconfiguration-gen
VGOPATH ?= $(LOCALBIN)/vgopath
GEN_CRD_API_REFERENCE_DOCS ?= $(LOCALBIN)/gen-crd-api-reference-docs
ADDLICENSE ?= $(LOCALBIN)/addlicense
MODELS_SCHEMA ?= $(LOCALBIN)/models-schema
GOIMPORTS ?= $(LOCALBIN)/goimports

## Tool Versions
KUSTOMIZE_VERSION ?= v5.0.0
CODE_GENERATOR_VERSION ?= v0.27.2
VGOPATH_VERSION ?= v0.1.1
CONTROLLER_TOOLS_VERSION ?= v0.11.3
VGOPATH_VERSION ?= v0.0.2
GEN_CRD_API_REFERENCE_DOCS_VERSION ?= v0.3.0
ADDLICENSE_VERSION ?= v1.1.0
GOIMPORTS_VERSION ?= v0.5.0

KUSTOMIZE_INSTALL_SCRIPT ?= "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"
.PHONY: kustomize
kustomize: $(KUSTOMIZE) ## Download kustomize locally if necessary.
$(KUSTOMIZE): $(LOCALBIN)
	@if test -x $(LOCALBIN)/kustomize && ! $(LOCALBIN)/kustomize version | grep -q $(KUSTOMIZE_VERSION); then \
		echo "$(LOCALBIN)/kustomize version is not expected $(KUSTOMIZE_VERSION). Removing it before installing."; \
		rm -rf $(LOCALBIN)/kustomize; \
	fi
	test -s $(LOCALBIN)/kustomize || { curl -Ss $(KUSTOMIZE_INSTALL_SCRIPT) | bash -s -- $(subst v,,$(KUSTOMIZE_VERSION)) $(LOCALBIN); }

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/controller-gen && $(LOCALBIN)/controller-gen --version | grep -q $(CONTROLLER_TOOLS_VERSION) || \
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: deepcopy-gen
deepcopy-gen: $(DEEPCOPY_GEN) ## Download deepcopy-gen locally if necessary.
$(DEEPCOPY_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/deepcopy-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/deepcopy-gen@$(CODE_GENERATOR_VERSION)

.PHONY: client-gen
client-gen: $(CLIENT_GEN) ## Download client-gen locally if necessary.
$(CLIENT_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/client-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/client-gen@$(CODE_GENERATOR_VERSION)

.PHONY: lister-gen
lister-gen: $(LISTER_GEN) ## Download lister-gen locally if necessary.
$(LISTER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/lister-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/lister-gen@$(CODE_GENERATOR_VERSION)

.PHONY: informer-gen
informer-gen: $(INFORMER_GEN) ## Download informer-gen locally if necessary.
$(INFORMER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/informer-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/informer-gen@$(CODE_GENERATOR_VERSION)

.PHONY: defaulter-gen
defaulter-gen: $(DEFAULTER_GEN) ## Download defaulter-gen locally if necessary.
$(DEFAULTER_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/defaulter-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/defaulter-gen@$(CODE_GENERATOR_VERSION)

.PHONY: conversion-gen
conversion-gen: $(CONVERSION_GEN) ## Download conversion-gen locally if necessary.
$(CONVERSION_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/conversion-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/conversion-gen@$(CODE_GENERATOR_VERSION)

.PHONY: openapi-gen
openapi-gen: $(OPENAPI_GEN) ## Download openapi-gen locally if necessary.
$(OPENAPI_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/openapi-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/openapi-gen@$(CODE_GENERATOR_VERSION)

.PHONY: applyconfiguration-gen
applyconfiguration-gen: $(APPLYCONFIGURATION_GEN) ## Download applyconfiguration-gen locally if necessary.
$(APPLYCONFIGURATION_GEN): $(LOCALBIN)
	test -s $(LOCALBIN)/applyconfiguration-gen || GOBIN=$(LOCALBIN) go install k8s.io/code-generator/cmd/applyconfiguration-gen@$(CODE_GENERATOR_VERSION)

.PHONY: vgopath
vgopath: $(VGOPATH) ## Download vgopath locally if necessary.
.PHONY: $(VGOPATH)
$(VGOPATH): $(LOCALBIN)
	@if test -x $(LOCALBIN)/vgopath && ! $(LOCALBIN)/vgopath version | grep -q $(VGOPATH_VERSION); then \
		echo "$(LOCALBIN)/vgopath version is not expected $(VGOPATH_VERSION). Removing it before installing."; \
		rm -rf $(LOCALBIN)/vgopath; \
	fi
	test -s $(LOCALBIN)/vgopath || GOBIN=$(LOCALBIN) go install github.com/onmetal/vgopath@$(VGOPATH_VERSION)

.PHONY: envtest
envtest: $(ENVTEST) ## Download envtest-setup locally if necessary.
$(ENVTEST): $(LOCALBIN)
	test -s $(LOCALBIN)/setup-envtest || GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-runtime/tools/setup-envtest@latest

.PHONY: openapi-extractor
openapi-extractor: $(OPENAPI_EXTRACTOR) ## Download openapi-extractor locally if necessary.
$(OPENAPI_EXTRACTOR): $(LOCALBIN)
	test -s $(LOCALBIN)/openapi-extractor || GOBIN=$(LOCALBIN) go install github.com/onmetal/openapi-extractor/cmd/openapi-extractor@latest

.PHONY: gen-crd-api-reference-docs
gen-crd-api-reference-docs: $(GEN_CRD_API_REFERENCE_DOCS) ## Download gen-crd-api-reference-docs locally if necessary.
$(GEN_CRD_API_REFERENCE_DOCS): $(LOCALBIN)
	test -s $(LOCALBIN)/gen-crd-api-reference-docs || GOBIN=$(LOCALBIN) go install github.com/ahmetb/gen-crd-api-reference-docs@$(GEN_CRD_API_REFERENCE_DOCS_VERSION)

.PHONY: addlicense
addlicense: $(ADDLICENSE) ## Download addlicense locally if necessary.
$(ADDLICENSE): $(LOCALBIN)
	test -s $(LOCALBIN)/addlicense || GOBIN=$(LOCALBIN) go install github.com/google/addlicense@$(ADDLICENSE_VERSION)

.PHONY: models-schema
models-schema: $(MODELS_SCHEMA) ## Install models-schema locally if necessary.
$(MODELS_SCHEMA): $(LOCALBIN)
	test -s $(LOCALBIN)/models-schema || GOBIN=$(LOCALBIN) go install github.com/onmetal/onmetal-api-net/models-schema

.PHONY: goimports
goimports: $(GOIMPORTS) ## Download goimports locally if necessary.
$(GOIMPORTS): $(LOCALBIN)
	test -s $(LOCALBIN)/goimports || GOBIN=$(LOCALBIN) go install golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)
