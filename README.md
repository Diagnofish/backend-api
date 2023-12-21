# Diagnofish Backend API

## Description

Diagnofish is an application that enables fish disease detection through images sent by users. The images sent by users will be processed by machine learning, and the detection results will be used to determine what kind of disease is attacking, how to treat it, and how to prevent it.

- This API uses a **PostgreSQL** database to store user data and detection result data.
- This API uses buckets in **Cloud Storage** to store the required files.

## Endpoints 

- **user**
  - Send **POST** request to the `/user/register` endpoint for registration process
  - Send **POST** request to the `/user/login` endpoint for the login process 
  - Send **POST** request to the `/user/logout` endpoint for the logout process
- **detection**
  - Perform fish disease detection and get information about it via a **POST** request to the `/detection` endpoint.
  - Get a list of detection history by sending a **GET** request to the `/detection/history` endpoint.
  - Get details of the detection history by sending a **GET** request to the `/detection/history/:id` endpoint.
 
for more complete documentation on available endpoints, please check here. [Diagnofish API Documentation](https://documenter.getpostman.com/view/21174179/2s9YkrZeBF)


## How to Deploy to Cloud Run

1. Select **Create service** on Cloud Run
2. Select **Continuously deploy new revisions from a source repository**
3. Select **Set up with Cloud Build**
4. Choose this repository
5. Select **Create**
