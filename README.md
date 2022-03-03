# Api Hyper Media


# Stock Investments
#### Easy way to control my investments

## Technologies
- `Language` [Golang](https://go.dev/)
    - `Framework`  [Gin](https://gin-gonic.com/)
- `Database` [MySql](https://www.mysql.com/)
    - `ORM` [Prisma2.0](https://www.prisma.io/)
- `Deploy` [MySql](https://www.mysql.com/)
    - `Deploy Database` [Planet Scale](https://planetscale.com/)
    - `Aplication` [Docker](https://www.docker.com/)
    - `Docker` [Heroku](https://devcenter.heroku.com/)


## Functional Requirements
    Login system
       - Personal Login 
    Stock management 
        - store position(buy,sell) [Add,Edit,Delete]
        - centralized wallet
        - real time stocks prices
        - get news related to the stocks that the user bought
        - get wallet from famous investors
        - calculate profit or lost for each position

## Non Functional Requirements
    Scalability
       - Load Balancer - Round-Robin
    Security 
        - JWT - BEARER Token
    Usability
        - subscriver that updates in real time the stocks ??
    Design
        - siren

# Description
### Stock Hyper Media Api

**Creating a [hyper media](https://api.gov.au/standards/national_api_standards/hypermedia.html) api using [siren](https://github.com/kevinswiber/siren) specifications.**

The Api is for investors that will register all their positions and the brokers in what they were made. 
For that reason the user will need to be registered in the api to insure privacy.
In the api they will be able to know the current price of every stock,register as many positions as they want in any broker.
Though the api they will be able to create a subscriber to see updating the price off any stock in real time
