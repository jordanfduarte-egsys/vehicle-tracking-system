# API - Vehicle Tracking System

O objetivo do projeto é receber a localização do veículo com a sua velocidade atual. Caso
a velocidade atual esteja acima do cadastrado, é enviado uma notificação para os sistemas
cadastrados através de goroutines e http requests.

## Instalação e Execução

Primeiro, certifique-se de ter configurado \$GOPATH.

```bash
# Download do projeto
go get github.com/jordanfduarte/vehicle-tracking-system

# Isso pode demorar alguns minutos
```

Defina o ambiente do projeto e execute

```bash
# vá para o diretorio do projeto
cd $GOPATH/src/jordanfduarte/vehicle-tracking-system

# criar o container da ultima versão do mysql com o Docker
docker-compose up

# ou utilizar o comando ambos fazem a mesma coisa
docker run --name mysql -e MYSQL_ROOT_PASSWORD=docker -d -p 3306:3306  mysql:latest

# utilize o comando para força Go a se comportar da maneira $GOPATH, mesmo fora do $GOPATH.
export GO111MODULE=off

# executa o projeto
go run main.go

# Endpoint da api:
http://localhost:8000/

# Executar inicialmente antes de qual quer request a url de migração de dados
http://localhost:8000/migration
```

## Testes
```bash

# para realizar os testes dos endpoints exeutar o comando
go test -v
```

## Estrutura de pastas do projeto
- Action
  - Implementa os Handler dos endpoints
    - database.go (DatabaseAction)
    - fleet_alert.go (AlertsGetAction, AlertsPostAction)
    - fleet.go (FleetsGetAction, FleetsPostAction)
    - index.go (IndexAction)
    - migration.go (MigrationAction)
    - respond.go (Respond, Error, JSON) - Responsável pelo prepago do retorno JSON
    - vehicle_position.go (PositionsGetAction, PositionsPostAction)
    - vehicle.go (VehiclesGetAction, VehiclesPostAction)
- Application
  - Escrever lógica de negócios
    - fleet_alert.go (RemoveFleetAlertAll, GetAllFleetAlerts, AddFleetAlert)
    - fleet.go (AddFleet, RemoveFleetAll, GetAllFleets, GetRowFleet)
    - vehicle_position.go (RemoveVehiclePositionAll, GetAllPositionsByVehicles, AddPositionVehicle)
    - vehicle.go (RemoveVehicleAll, AddVehicle, GetAllVehicles)
- Domain
  - Entity struct que representam o mapeamento para o modelo de dados
    - factory.go - Factory de criação de objetos dos modelos de dados
    - fleet_alert.go
    - fleet.go
    - site_goroutine.go - Implementação do goroutine
    - vehicle_position.go
    - vehicle.go
  - Repository para infraestrutura
- Infrastructure
  - Implementa a interface do repositório
    - fleet_alert_repository.go
    - fleet_repository.go
    - vehicle_position_repository.go
    - vehicle_repository.go
- Interfaces
  - Configuração das rotas e chamadas dos Handlers

## URL ENDPOINT

#### /migration

- `GET` : Migra os dados iniciais da base de dados (Cria as tabelas e cria alguns registros iniciais)

#### /index

- `GET` : Index da api

#### /database

- `DELETE` : Limpa toda a base de dados inclusive os registros iniciais cadastrados no endpoint `/migration`

#### /fleets

- `GET` : Lista todas as frotas
- `POST` : Cria uma frota

#### /fleets/{id}/alerts

- `{id}` : ID da frota
- `GET` : Lista todas os alertas de uma frota)
- `POST` : Cria uma alerta para frota

#### /vehicles

- `GET` : Lista todos os veículos
- `POST`: Cria uma veículo

#### /vehicles/{id}/positions

- `{id}` : ID do veículo
- `GET` : lista todos as posições de um veículo
- `POST` : Salva a posição do veículo

### Link da documentação dos endpoints

> Postman https://documenter.getpostman.com/view/3003865/UVkiTync


Listagem de frotas, URL GET `/fleets`
```bash
curl --location --request GET 'localhost:8000/fleets'
```

Criação de frota, URL PUT `/fleets`
```bash
curl --location --request POST 'localhost:8000/fleets' \
--data-raw '{ "name": "Veículos de perseguição", "max_speed": 10 }'
```


Endpoint para recuperar todos os alertas cadastrados em uma frota, URL GET `/fleets/1/alerts`

```bash
curl --location --request GET 'localhost:8000/fleets/1/alerts'
```

Adicionar um novo alerta para uma frota, URL POST `/fleets/1/alerts`

```bash
curl --location --request POST 'localhost:8000/fleets/1/alerts' \
--data-raw '{ "webhook": "http://host:8080/dsds/dsds" }'
```

Recupera todos os veículos, URL GET `/vehicles`

```bash
curl --location --request GET 'localhost:8000/vehicles'
```

Adiciona um novo veículo, URL POST `/vehicles`

```bash
curl --location --request POST 'localhost:8000/vehicles' \
--data-raw '{"fleet_id": 2, "name": "veículo 3", "max_speed": null}'
```

Recupera as posições de um veículo, URL GET `/vehicles/{id}/positions`
```bash
curl --location --request GET 'localhost:8000/vehicles/1/positions' \
--data-raw ''
```

Adiciona uma nova posição para um veículo, URL POST `/vehicles/{id}/positions`
```bash
curl --location --request POST 'localhost:8000/vehicles/1/positions' \
--data-raw '{ "timestamp": "ISO-8601", "latitude": 0, "longitude": 0, "current_speed": 1000 }'
```

Cria a estrutura inicial do banco de dados, URL GET `/migration`
```bash
curl --location --request GET 'localhost:8000/migration'
```

Limpeza da base de dados, URL DELETE `/database`
```bash
curl --location --request DELETE 'localhost:8000/database'
```

## FLuxograma da API

![fluxograma1](https://raw.githubusercontent.com/jordanfduarte/vehicle-tracking-system/master/assets/fluxograma-1.png)

![fluxograma2](https://raw.githubusercontent.com/jordanfduarte/vehicle-tracking-system/master/assets/fluxograma-2.png)

## Diagrama do banco de dados

![diagrama](https://raw.githubusercontent.com/jordanfduarte/vehicle-tracking-system/master/assets/diagrama-db.png)


## Lista de pendências de itens do produto (Backlog)

- [x] **Obrigatório:** Banco de dados Mysql
- [x] **Obrigatório:** Criação dos endpoints
  - [x] DDD
  - [x] TDD
  - [x] Factory para criação de objetos
  - [x] WorkerPool
- [x] **Obrigatório:** Testes
- [ ] **Opcional:** Deploy
- [x] **Opcional:** Fluxograma

## Alteração da porta de execução do projeto

- Acessar o arquivo /interfaces/handler.go
- Trocar a porta padrão de 8000 para outra xxxx


## Referências & Bibliotecas & Dicas

- DDD Skeleton : https://github.com/takashabe/go-ddd-sample
- Httprouter : https://github.com/julienschmidt/httprouter
- GORM Documentation : https://gorm.io/
- Toml : https://github.com/BurntSushi/toml
- Docker : https://hub.docker.com/_/mysql
- Tests : https://medium.com/@sheimyrahman/golang-go-e-tdd-para-iniciantes-2418b6eadd92
- GO111MODULE=off : https://maelvls.dev/go111module-everywhere/
