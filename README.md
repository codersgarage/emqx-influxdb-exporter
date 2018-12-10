# Golang RESTful boilerplate

##### Included

- chi (router)
- consul (read remote config)
- mysql (persistent layer)
- rabbitmq (pubsub)
- machinery (worker library)
- docker
- glide (dependency management)


##### Architecture
Implement SOLID design principle

- api : contains api routes and handlers
- app : contains database instance and scope
- cmd : commands to execute
- config : application configuration
- data : data access layer (dao)
- errors : error definitions
- log : logging utils
- models : database models
- mq : messages queue
- repos : repository to validate data coming from request
- utils : utility stuffs
- worker : worker to consume tasks
- build.sh : helper to build and run application

Distributed under MIT license.
