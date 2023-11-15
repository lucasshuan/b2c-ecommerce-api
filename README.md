
```
ecommerce-api
├─ 📁cmd
│  └─ 📄main.go
├─ 📁docker
│  ├─ 📄company-01.docker-compose.yml
│  └─ 📄company-02.docker-compose.yml
├─ 📁internal
│  ├─ 📁database
│  │  └─ 📄gorm.go
│  ├─ 📁gql
│  │  ├─ 📁resolvers
│  │  │  ├─ 📄cart.resolvers.go
│  │  │  ├─ 📄category.resolvers.go
│  │  │  ├─ 📄order.resolvers.go
│  │  │  ├─ 📄product.resolvers.go
│  │  │  ├─ 📄resolver.go
│  │  │  ├─ 📄schema.resolvers.go
│  │  │  └─ 📄user.resolvers.go
│  │  ├─ 📁schema
│  │  │  ├─ 📄cart.graphqls
│  │  │  ├─ 📄category.graphqls
│  │  │  ├─ 📄order.graphqls
│  │  │  ├─ 📄product.graphqls
│  │  │  ├─ 📄schema.graphqls
│  │  │  └─ 📄user.graphqls
│  │  ├─ 📄generated.go
│  │  ├─ 📄models_gen.go
│  │  ├─ 📄resolver.go
│  │  └─ 📄user.resolvers.go
│  ├─ 📁middleware
│  │  ├─ 📄auth_middleware.go
│  │  └─ 📄tenant_middleware.go
│  ├─ 📁model
│  │  ├─ 📄cart.go
│  │  ├─ 📄cart_product.go
│  │  ├─ 📄category.go
│  │  ├─ 📄order.go
│  │  ├─ 📄product.go
│  │  └─ 📄user.go
│  ├─ 📁repository
│  │  ├─ 📄cart.go
│  │  ├─ 📄cart_product.go
│  │  ├─ 📄category.go
│  │  ├─ 📄order.go
│  │  ├─ 📄product.go
│  │  └─ 📄user.go
│  ├─ 📁server
│  │  └─ 📄server.go
│  ├─ 📁service
│  └─ 📁util
│     └─ 📄bcrypt.go
├─ 📄.env
├─ 📄.gitignore
├─ 📄docker-compose.yml
├─ 📄go.mod
├─ 📄go.work
├─ 📄go.work.sum
├─ 📄gqlgen.yml
└─ 📄tools.go
```