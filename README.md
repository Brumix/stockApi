# Stocks Api

# Stock Investments
#### Easy way to control my investments

## Technologies

- `Language` [Golang](https://go.dev/)
  - `Framework`  [Gin](https://gin-gonic.com/)
- `Database` [Postgres](https://www.postgresql.org/)
  - `ORM` [Gorm](https://gorm.io/)
- `Deploy`
  - `Aplication` [Docker](https://www.docker.com/)
  - `Docker` [AWS](https://aws.amazon.com/pt/)


## Functional Requirements
    Login system
       - Personal Login 
    Stock management 
        - store position(buy,sell) [Add,Edit,Delete]
        - centralized wallet
        - real time stocks prices
        - calculate profit or lost for each position

## Non Functional Requirements
    Scalability
       - Load Balancer - Round-Robin
    Security 
        - JWT - BEARER Token - Cookies
    Usability
        - subscriver that updates in real time the stocks ??
    Design
        - Microservices

# Description
### Stock Hyper Media Api

The Api is for investors that will register all their positions and the brokers in what they were made. 
For that reason the user will need to be registered in the api to insure privacy.
In the api they will be able to know the current price of every stock,register as many positions as they want in any broker.
Though the api they will be able to create a subscriber to see updating the price off any stock in real time

## Other Mircoservices
[Auth Micro service](https://github.com/Brumix/authServer)

[Stock Node Micro service](https://github.com/Brumix/stockMicroServico)

