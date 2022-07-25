# microservices-golang-grpc-k8s
docker build -t gcr.io/diesel-command-355401/microservices-golang-grpc-k8s:v0.0.2 .
docker push gcr.io/diesel-command-355401/microservices-golang-grpc-k8s:v0.0.2

kubectl create ns greeting
kubectl apply -f k8s/deploy.yaml
kubectl -n greeting get pods
kubectl -n greeting logs -f greeting-6d5f95b844-p8k5p
kubectl -n greeting port-forward greeting-6d5f95b844-p8k5p 50051:50051



