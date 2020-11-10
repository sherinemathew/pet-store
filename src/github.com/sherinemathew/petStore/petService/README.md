# Pet Store
Created By Sherine Mary Mathew.

## API requirements
The API requirements have been set and can be found at https://petstore.swagger.io.

### Functionality/Features
- Implemented one endpoint - Find pet by Id (/pet/{petId})

#### Description
At application startup, 50 fake pets are inserted into BoltDB (key-value store). petId is from 1-50 which can be queried from the microservice.

#### Below is an example
Sample Request -> curl http://localhost:6004/pet/12
Sample Response -> {"id":12,"category":{"Id":0,"Name":""},"name":"Pet_11","status":"available"}

### Technologies Used
This project has been developed in Go.

### Nice to have (future enhancements)- Planned to do the below if more time was available.

* Implement API gateway.
* Implement endpoints mentioned in the swagger.
* Write Unit and Integration tests.
* Implement a database like MySQL/PostgreSQL
* Dockerize the microservice.
