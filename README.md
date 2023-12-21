# Diagnofish Backend API

### Description

Diagnofish is an application that enables fish disease detection through images sent by users. The images sent by users will be processed by machine learning, and the detection results will be used to determine what kind of disease is attacking, how to treat it, and how to prevent it.

This API uses a **PostgreSQL** database to store user data and detection result data.
This API uses buckets in **Cloud Storage** to store the required files.

#### Endpoints 

- **user**
  - Send **POST** request to the `/user/register` endpoint for registration process
  - Send **POST** request to the `/user/login` endpoint for the login process 
  - Send **POST** request to the `/user/logout` endpoint for the logout process
- **detection**
  - Perform fish disease detection and get information about it via a **POST** request to the `/detection` endpoint.
  - Get a list of detection history by sending a **GET** request to the `/detection/history` endpoint.
  - Get details of the detection history by sending a **GET** request to the `/detection/history/:id` endpoint.
 
for more complete documentation on available endpoints, please check here. [Diagnofish API Documentation](https://documenter.getpostman.com/view/21174179/2s9YkrZeBF)


#### Folder Structures

📁 **repository**

This is a function that interacts with the Postgres database using GORM

📁 **service**

The service layer is used to process data according to predefined business rules.

📁 **middleware**

In the `middleware/auth.go` file there is an `Auth()` function that is used to authenticate users using JWT (JSON Web Token). This middleware serves to check whether a user who accesses a certain endpoint or route has been authenticated or not.

📁 **api**

Contains the HTTP handler for the API

#### How to Deploy to Cloud Run

2. Select **Create service** on Cloud Run
3. Select **Continuously deploy new revisions from a source repository**
4. Select **Set up with Cloud Build**
5. Choose this repository
6. Select **Create**
