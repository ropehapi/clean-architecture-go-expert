# Desafio Clean Architecture Go Expert
Esse projeto foi desenvolvido como a solução do desafio final do curso Go expert.

## Dev
Para rodar o projeto execute os seguintes comandos, responsáveis por subir os containeres
do banco de dados (já executando as migrations), o RabbitMQ e executar a aplicação.

> sudo docker compose up -d

> make run

Com a aplicação no ar, aqui vão as instruções para consumir a aplicação em suas diferentes interfaces.

### REST
No endereço `http://localhost:8000`, você pode fazer as requisições com um cliente HTTP como o postman, ou se preferir, 
utilizar as chamadas descritas no arquivo `requests.http` dentro do diretório `api`.

### GraphQL
No endereço `http://localhost:8080`, você pode rodar as seguintes querys afim de manipular a aplicação:

```
mutation CreateOrder {
createOrder(input:{
id: "American Big black",
Price: 10.0,
Tax: 2
}) {
id,
Price,
Tax,
FinalPrice
}
}

query queryOrders {
orders {
id,
Price,
Tax,
FinalPrice
}
}
```

### gRPC
Para consumir a aplicação via um cliente gRPC, basta rodar em um terminal com o evans instalado os seguintes comandos:
> evans -r repl
 
> package pb
 
> service <nome_da_service>

> call <nome_da_chamada>

### RabbitMQ
Para acessar o painel do RabbitMQ basta acessar o seguinte endereço: `http://localhost:15672`
