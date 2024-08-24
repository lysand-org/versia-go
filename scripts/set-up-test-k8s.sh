  #!/bin/sh

set -x

k3d cluster create versia-go --agents 1 -p "30000-30050:30000-30050@server:0" -p "8443:443@loadbalancer" -p "8080:80@loadbalancer" || true

helm repo add nats https://nats-io.github.io/k8s/helm/charts/ || true
helm repo update

helm install nats nats/nats \
	--set config.jetstream.enabled=true \
	--set config.cluster.enabled=true \
	--set config.cluster.replicas=2 \
	--set config.jetstream.fileStore.pvc.size=1Gi

opts=$(cat <<EOF
  --set image.tag=main
  --set nats.uri=nats://nats:4222
  --set versia.instance.address=http://localhost:8080
  `#--set versia.telemetry.sentryDSN=`
EOF
)

# shellcheck disable=SC2086
helm install versia ./chart/ $opts || helm upgrade versia ./chart/ $opts