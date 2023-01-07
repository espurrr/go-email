# go-email

A golang mail client made using the package [gomail](https://pkg.go.dev/gopkg.in/mail.v2@v2.3.1). (using SMTP) 

## Pre-requisites

Rename the .env.example file to **.env**. Set the email configurations in the file and save it.  !Important : Make sure your gmail account's *less secure app access* option is ON.

## Run the client

### By running the go app  

Run the following command to start the go application

      go run .

The application will be up on http://localhost:8081/
(You can change the server port in [main.go](https://gitlab.com/vip16/goemailclient/-/blob/main/main.go), if port 8081 is occupied)

### By running the dockerfile

Run the following command to build the dockerfile

     docker build -t goemailclient .

Now run the docker image

     docker run -it -p 8080:8081 goemailclient

The application will be up on http://localhost:8080/


## Test the API

With Postman or your preferred API client, send a POST request with the following JSON code as its body.

    {
		"address": "<enter_receivers_email_address_here>",
		"subject": "<enter_email_subject_here>",
		"body": "<enter_email_body_here>"
	}

The receiver should get the email the minute you send the request.  
If an error occurs; follow through the logs printed onto the console.
