###  Repository points to the right database connection, create table, migrations, and offer all the operations to that database (which could be an ORM) 
### Service offers all allowed operations to a controller
### Controller builds and receives the request

# Overview
This project follows a layered architecture to manage Restfull API operations using Go. The design focuses on clean separation of concerns, making it easy to extend, test, and maintain.

`UserController -> UserService -> UserRepository (Interface) -> MongoUserRepository (Implementation)`

# Layers Description
1. Controller Layer
Component: UserController

Responsibility: Handles HTTP requests, calls UserService, and prepares HTTP responses.

Details:
This layer does not contain any business logic or database code.

2. Service Layer
Component: UserService

Responsibility: Contains business logic for user operations (e.g., validation, processing).

Details:
Calls the UserRepository interface for database interactions.

3. Repository Layer
Component: UserRepository (interface)

Responsibility: Defines the contract for data access operations (CRUD).

Details:
The service layer depends on this interface, not on specific implementations.

4. Data Access Layer
Component: MongoUserRepository (MongoDB implementation)

Responsibility: Actual interaction with MongoDB (using MongoDB Go Driver).

Details:
Implements all methods declared by UserRepository.

# Testing Strategy
Service Layer Testing: Use mock implementations of UserRepository for unit tests.

Repository Layer Testing: Use either a test database or a mock MongoDB server like MongoDB Memory Server.

# Requirements
Go 1.20+

MongoDB

MongoDB Go Driver (go.mongodb.org/mongo-driver/mongo)

## TODOS
- Create a SproutContext where all configurations will live in
- Create Bean "annotation": automatically assigned to the context and multual single instance reference among Beans
