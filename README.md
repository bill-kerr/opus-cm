<!-- prettier-ignore-start -->

# **Opus CM**

Opus CM is an open source construction project management platform built on asynchronous microservices. Opus utilizes [NATS Streaming Server](https://docs.nats.io/) as a central event bus for message-based communication between services.

## **Services**

Opus CM is organized around multiple services. Each "microservice" is responsible for a small portion of the system and communicates with other services via a pub/sub pattern, using a central event bus as a message broker. Below is the list of services that Opus CM is built on.

| Service | Description |
| ------- | ----------- |
| Users Service | The Users Service controls user registration, password management, and role management.
| Organizations Service | The Organizations Service controls all functionality for creating, reading, updating, and deleting organizations.
| Notifications Service | The Notifications Service is responsible for dispatching all notifications to users. For now, notifications are only dispatched as emails.
| Submittals Service | The Submittals Service controls all CRUD operations for submittals. |

## **Events**

Events consist of a subject and a payload. Publishers dispatch events with a subject and subscribers to that subject receive the event and its payload. Subscribers belong to a "queue group" which tells the event bus that only one of the subscribers in a given queue group is to receive the event. All payload data is transfered as a single string and must be sent as and parsed to JSON data structures.

### **Payload Types**

It is necessary to define common data structures that each service can understand. Below are definitions for data structures that may appear as payloads in an event.

| Type | Structure |
| --------- | --------- |
| User | ``` { id: uuid, email: string, role: Role } ```
| Role | ``` SYS_ADMIN, PRJ_ADMIN, PRJ_USER (default) ```

### **Event List**

Below is a table of events, with their subjects and payload structures.

| Subject | Payload |
| ------- | ------- |
| user:created | ``` User ```
| user:updated | ``` User ```
| user:role_changed | ``` { id: uuid, role: Role } ```

## **HTTP Responses**

Responses returned to users are to be labelled with the Content-Type ```application/json```. Response shapes across services must conform to the same naming conventions and basic data structure. Below is a list of response types and the required shape of the response data.

## **Errors**

Like successful HTTP responses, returned errors must conform to a common definition. Below is the list of errors and their data shapes.

| Error | Status Code | Data |
| ----- | ----------- | ---- |
| Internal Server Error | 500 | <code>{<br />&nbsp;&nbsp;object: 'error',<br />&nbsp;&nbsp;name: 'Internal server error', <br />&nbsp;&nbsp;details: 'An unknown error occurred.'<br />}</code>


<!-- prettier-ignore-end -->
