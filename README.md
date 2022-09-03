# Learn Microservice - Basic CQRS
### Below concept are covered in this repository:
- API Gateway
- gRPC
- Basic CQRS pattern with rabbitMQ

### Architecture

```
                          ||=========================||
                          ||     MySQL1(Write DB)    ||
                          ||=========================||
                                        ↑
                                        | Write
||==============|| gRPC   ||=========================|| pub [product_add] ||==========||
|| API Gateway  ||------->|| Product Command Service ||------------------>|| rabbitMQ ||
||==============||        ||=========================||                   ||==========||
        |                                                                       |
        |      gRPC       ||=========================||                         |
        ----------------> || Product Query Service   ||                         |
                          ||=========================||                         |
                                        ↑                                       |
                                        |                     sub [product_add] |          
                                        | Read                                  |
                                        |                       ||==========================||
                          ||=========================||         ||     Product Listener     ||
                          ||     MySQL2(Read DB)     ||<------- ||          Service         ||
                          ||=========================||   Write ||==========================||

```
### Explanation
This example is consist of 2 API endpoint
- /product/add : when user hit add product endpoint, API gateway will reach product command service through gRPC. Product command service will insert new product data to mysql1 and publish product_add event to message broker. Product listener is already listen to product_add event, so it will consume product_add event and thus insert new product data in mysql2
- /product/search : API gateway will reach product query service through gRPC, product query service will directly read data from mysql2 and return search result to API gateway


### Deployment
1. Using docker compose
```shell
docker-compose up --build
```
2. Open your browser and access http://localhost:3000/docs/
