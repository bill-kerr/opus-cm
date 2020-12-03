<!-- prettier-ignore-start -->

# **Opus CM**
Opus CM is an open source construction project management platform built on asynchronous microservices. Opus utilizes [NATS Streaming Server](https://docs.nats.io/) as a central event bus for message-based communication between services and [Kubernetes](https://https://kubernetes.io/) for container orchestration.
<br />

## **Services**
Opus CM is organized around multiple services. Each "microservice" is responsible for a small portion of the system and communicates with other services via a pub/sub pattern, using a central event bus as a message broker. Below is the list of services that Opus CM is built on.

| Service | Description |
| :------ | :---------- |
| Users Service | The Users Service controls user registration, password management, and role management. |
| Organizations Service | The Organizations Service controls all functionality for creating, reading, updating, and deleting organizations. |
| Notifications Service | The Notifications Service is responsible for dispatching all notifications to users. For now, notifications are only dispatched as emails. |
| Submittals Service | The Submittals Service controls all CRUD operations for submittals. |
<br />

## **Authentication & Roles**
Registration and authentication of users is handled by [Firebase](https://firebase.google.com/). Each service will utilize the Firebase Admin client to authenticate incoming requests. Additionally, a custom claim will be embedded in the JWT to define the user's role.

### **Roles**
The following table lists the roles available in Opus CM.

| Role | Key | Description |
| :--- | :-- | :---------- |
| System Admin | SYS_ADMIN | System Admins have permissions to perform any operation. They are the superusers of the application.
| Project Admin | PRJ_ADMIN | Project Admins are assigned to one or more projects and have full permissions on a project level.
| Project User | PRJ_USER | Project Users are assigned to one or more projects and have permissions and roles as defined by Project Admins.
<br />

## **Events**
Events consist of a subject and a payload. Publishers dispatch events with a subject and subscribers to that subject receive the event and its payload. Subscribers belong to a "queue group" which tells the event bus that only one of the subscribers in a given queue group is to receive the event. All payload data is transfered as a single string and must be sent as and parsed to JSON data structures.

### **Event List**
Below is a table of events, with their subjects and payload structures.

| Subject | Payload |
| :------ | :------ |
| user:created | ``` User ``` |
| user:updated | ``` User ``` |
| user:role_changed | ``` { id: uuid, role: Role } ``` |
| organization:created | ``` Organization ``` |
| organization:updated | ``` Organization ``` |
| organization:deleted | ``` { id: uuid } ``` |
<br />

### **Payload Types**
It is necessary to define common data structures that each service can understand. Below are definitions for data structures that may appear as payloads in an event.

| Type | Structure |
| :-------- | :-------- |
| User | ``` { id: uuid, email: string, role: Role } ``` |
| Organization | ``` { id: uuid, name: string } ``` |
</br>

## **HTTP Responses**
Responses returned to users are to include the Content-Type header ```application/json```. Keys in the JSON data must be named using camel case.

### **Object Key**
Each object in the JSON response should have an "object" key, which will inform the consumer about what type of data to expect. List or array responses should have a key of "list" and an additional key of "data", which contains the array or list of objects.

### **Errors**
Like successful HTTP responses, returned errors must conform to a common definition. Below is a list of errors and their default messages. In most cases, the default error message should not be used, and a more specific message should be provided.

| Error Name | Status Code | Details |
| :--------- | :---------- | :------ |
| Bad Request Error | 400 | Request was poorly formatted. |
| Validation Error** | 400 | A request validation error occurred. |
| Unauthorized Error | 401 | An authorization token was not provided or is invalid. |
| Insufficient Permissions Error | 403 | You do not have the requisite permissions to perform this operation. |
| Not Found Error | 404 | The requested resource was not found. |
| Internal Server Error | 500 | An unknown error occurred. |
** Validation errors may contain a list of errors under the "details" key.

## **Environmental Variables**
Opus CM requires some global environmental variables to be exposed to each service. Below is the list of required environmental variables. These variables are exposed to each service through Kubernetes secrets. Specific services may require additional environmental variables.

| Secret Name | Type | Key | Description |
| :---------- | :--- | :-- | :---------- |
| firebase-config | file | n/a | This secret points to a volume containing the Google application credentials required for initializing the authenticaion capabilities of Firebase. |
| pg-user | literal | PG_USER | This secret contains the username to be used for all PostgreSQL database connections across the application. |
| pg-password | literal | PG_PASSWORD | This secret contains the password to be used for all PostgreSQL database connections across the application. |
| nats-cluster-id | literal | NATS_CLUSTER_ID | This secret contains the cluster name used by the NATS Streaming Server event bus. All "clients" must connect to this cluster. |

<!-- prettier-ignore-end -->
