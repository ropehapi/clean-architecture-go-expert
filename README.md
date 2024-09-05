//TODO: Escrever a documentação disso

sudo docker compose up -d <br>
go run main.go wire_generated.go <br>
REST: http://localhost:8000
gRPC: ~evans -r repl <br>
GraphQL: http://localhost:8080 <br>
RabbitMQ: http://localhost:15672 <br>

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