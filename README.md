#Docker Mysql
    docker run --name mysql -e MYSQL_ROOT_PASSWORD=docker -d -p 3306:3306  mysql:latest


export GO111MODULE=off

#gerenciador de dependencias
    go get -u github.com/golang/dep/cmd/dep
    dep init
    #https://imasters.com.br/back-end/gerenciando-dependencias-em-golang

go test -v

#documentação
https://documenter.getpostman.com/view/3003865/UVkiTync
------------------------------------------------------------
#https://github.com/jojoarianto/go-ddd-api
#https://go.dev/tour/basics/11
#https://yashodgayashan.medium.com/factory-design-pattern-in-golang-d2fc92223ee2