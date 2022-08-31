docker build product_q_srv -t registry.dev.svc.cluster.local:62119/product_q:0.0.1
docker push registry.dev.svc.cluster.local:62119/product_q:0.0.1

docker build product_c_srv -t registry.dev.svc.cluster.local:62119/product_c:0.0.1
docker push registry.dev.svc.cluster.local:62119/product_c:0.0.1

docker build api_gateway -t registry.dev.svc.cluster.local:62119/api_gateway:0.0.1
docker push registry.dev.svc.cluster.local:62119/api_gateway:0.0.1

docker build product_listener_srv -t registry.dev.svc.cluster.local:62119/product_listener:0.0.1
docker push registry.dev.svc.cluster.local:62119/product_listener:0.0.1