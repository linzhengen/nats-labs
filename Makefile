deploy-nats:
	helm repo add nats https://nats-io.github.io/k8s/helm/charts/ --force-update
	helm repo update nats
	helm upgrade --install nats-js nats/nats \
	--create-namespace \
	--namespace nats-js \
	--set nats.jetstream.enabled=true

port-forward:
	kubectl -n nats-js port-forward service/nats-js 4222:4222


install-cli:
	go install github.com/nats-io/natscli/nats@latest

add-nats-cxt:
	nats context add nats --server localhost:4222 --description "NATS minikube" --select