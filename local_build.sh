REGISTRY_HOST=registry.dev.svc.cluster.local
REGISTRY_PORT=62119

docker build product_q_srv -t $REGISTRY_HOST:$REGISTRY_PORT/product_q:0.0.1
docker push $REGISTRY_HOST:$REGISTRY_PORT/product_q:0.0.1

docker build product_c_srv -t $REGISTRY_HOST:REGISTRY_PORT/product_c:0.0.1
docker push $REGISTRY_HOST:$REGISTRY_PORT/product_c:0.0.1

docker build api_gateway -t $REGISTRY_HOST:$REGISTRY_PORT/api_gateway:0.0.1
docker push $REGISTRY_HOST:$REGISTRY_PORT/api_gateway:0.0.1

docker build product_listener_srv -t $REGISTRY_HOST:$REGISTRY_PORT/product_listener:0.0.1
docker push $REGISTRY_HOST:$REGISTRY_PORT/product_listener:0.0.1