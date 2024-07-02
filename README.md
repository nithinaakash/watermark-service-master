# Watermark-Service

## Problem Statement
A global publishing company specializing in books and journals seeks to develop a service to watermark their documents. The company publishes books on various topics such as business, science, and media, while journals do not include specific topics. Each document (whether a book or journal) has a title, author, and a watermark property. An empty watermark property indicates that the document has not been watermarked yet.

## Service Requirements

The watermark service needs to be asynchronous. For any given document, the service should return a ticket that can be used to poll the status of processing (e.g., Status: Started, Pending, Finished). Once watermarking is complete, the document can be retrieved using the ticket. The watermark of a book or journal is identified by setting the watermark property of the object. For a book, the watermark includes the properties content, title, author, and topic. The journal watermark includes the content, title, and author.

### Examples of Watermarks

#### Book:

{
  "content": "book",
  "title": "The Dark Code",
  "author": "Bruce Wayne",
  "topic": "Science"
}

## Tasks

### Watermark Service:

Implement a microservice for watermarking documents and returning a ticket ID.
Implement user authorization API (Users: SuperAdmin, Default).

### Document Retrieval Service:

Implement a microservice for retrieving a document by its ticket ID once the watermarking is finished.

### Status Retrieval API:

Implement an API for retrieving the status of a document by its ticket ID.

### Implementation Details
Use Golang/gRPC stack and Postgress DB.
Provide sufficient unit tests to ensure the functionality of the service.
Include logging output to monitor various asynchronous watermark processes identified by a unique ticket ID.
Testing
Create a test script wrapper in the root directory to run tests.
Ensure the tests cover at least 10 books and 10 journals.
Development Environment
Set up a local development environment using Kubernetes ( Docker for Desktop, OKD, K3s, microk8s).

## Nice to Have
Create test scenarios for the following typical microservice patterns:

Ambassador API Gateway (Facade Pattern) <br>
Circuit Breaker <br>
Mirroring <br>
Canary Deployments <br>
Alerting with Grafana <br>
