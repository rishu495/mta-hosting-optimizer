# mta-hosting-optimizer

1. Service flow starts from main.go file

2. API exposed : localhost:8000/inactiveIPAddressHostname

3. Expected Response on X= 1 from this API -->

{
    "result": [
        "mta-prod-1",
        "mta-prod-3"
    ],
    "status": "OK",
    "error": null
}


4. Defined the Sample Date Set with the help of a function written in util.go

5. Main business logic written in usecase file

6. Common functions written in util.go

7. Fields defined in model.go file

8. ENV file not pushed on git. 

9. To make this work, follow below steps -->
  --> CLone this repo
  --> Make .env file and write X=1 and save 
  --> Run go mod vendor and go mod tidy
  --> From root folder, go build and go run app.go
  --> You are ready to test above provided API on Postman
  
 
 
 
 Thanks for reading! :)
