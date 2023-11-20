<div align="center">
   <img href="https://avatars.githubusercontent.com/u/78228526?v=4" alt="Noven" height="168" title="Noven" src="https://avatars.githubusercontent.com/u/78228526?v=4" />
   <h1>B2C E-commerce API</h1>
</div>

<p align="center">
  <a href="#about">📃 About</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#technologies">💻 Technologies</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#features">✨ Features</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#getting_started">🚀 Getting Started</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#usage">👨‍💻 Usage</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#roadmap">🗺️ Roadmap</a>&nbsp;&nbsp;|&nbsp;&nbsp;
  <a href="#author">👤 Author</a>
</p>

## 📄 About <a name = "about"></a>

B2C E-commerce API is a Golang project featuring GraphQL with gqlgen and Gin for Web Framework. 

The project adopts a well-organized controller-service-repository structure and supports various functionalities, including user registration, Google authentication, cart management, order creation, payment processing, and more. Additionally, it integrates Asaas for enhanced functionality.

One standout feature is its multitenant design. This means the project is crafted to effortlessly support various databases with different structures. This flexibility allows it to adapt and scale based on the specific needs or preferences of users and businesses.

## 💻 Tools <a name="tools"></a>

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [gqlgen](https://gqlgen.com)
- [Asaas](https://www.asaas.com/)
- [Gin](https://gin-gonic.com)

## ✨ Features <a name="features"></a>

- **Authentication and Authorization:** 
User registration with standard credentials or Google authentication.
- **Cart Management:** 
Users can add products, create orders, and track payment status.
- **Product Management:** 
Admins can add and categorize products for an enhanced shopping experience.
- **Order Processing:** 
Creation of orders and secure payment processing with detailed status tracking.
- **Payment Integration:** 
Seamless integration with Asaas for efficient and reliable payment processing.
- **User Lifecycle Enhancements:** 
Password recovery, email verification, and more to enhance user experience.
- **Multitenancy:** 
Designed for seamless scalability, the project supports multiple databases with varying structures.

## 🚀 Getting Started <a name = "getting_started"></a>

### Prerequisites

- **Go**
- **Docker**
- **make**

### Installing

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/ecommerce-api.git
    cd ecommerce-api
    ```

2. Install dependencies:

    ```bash
    make deps
    ```

## 👨‍💻 Usage <a name = "usage"></a>

### GraphQL code generation

1. Generate GraphQL code for new gql/schema:

    ```bash
    make gqlgen
    ```

### Manage Docker containers

1. Start containers:

    ```bash
    make docker-up
    ```

2. Stop containers:

    ```bash
    make docker-down
    ```

### Configuration

The project utilizes a configuration file named `config.yml` to manage environment-specific settings.

```yaml
env: dev
dev:
  port: 3000
  tenants:
    tenant1: scheme://user:password@host:port/db
    tenant2: ...
```

## 🗺️ Roadmap <a name="roadmap"></a>

- **Testing Suite:** Integrate a comprehensive testing suite for robust code validation.
- **Expanded Database Support:** Extend database support beyond PostgreSQL and MySQL to include additional databases.
- **Enhanced Security:** Implement advanced security measures, such as JWT for user authentication and authorization.
- **Notification System:** Integrate a notification system to keep users informed about order updates and promotions.
- **Order Tracking:** Integrate a user-friendly order tracking system, providing real-time updates on order status and location.


## 👤 Author <a name = "author"></a>

<div align="center">
  <table>
    <tr>
      <td align="center">
        <a href="http://github.com/nobeluc/">
          <img src="https://avatars.githubusercontent.com/u/78228526?v=4" width="75px;" alt="Nobeluc"/>
          <br />
          <sub>
            <b>Jean Rocha</b>
          </sub>
        </a>
        </td>
    </tr>
  </table>

[![Youtube Badge](https://img.shields.io/badge/-Noven-FF0000?style=flat-square&labelColor=FF0000&logo=youtube&logoColor=white&link=https://www.youtube.com/channel/UC4IOuH99CdKBPydv7CW8Tdg)](https://www.youtube.com/channel/UCgg16Rkdisznb1au1QgqjPA)
[![Linkedin Badge](https://img.shields.io/badge/-Jean%20Rocha-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/lucrocha2/)](https://www.linkedin.com/in/lucrocha2/)
</div>
