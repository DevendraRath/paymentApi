Steps for running APIs(Locally)
  step-1 : Go to project folder and run go mod download command( It will automatically download all the dependency packages)
  step-2 : Run go run main.go(It will run our application in port 3008

 note: Taking the API key from env file for the security purpose. So No need for provide the stripe key 
       within header. if someone need to change that key then he can change it in env file. 

List of APIs:
   1) POST /api/customer
          Url:-POST localhost:3008/api/customer
          Request body :{
                                           "email":"token@gmail.com",
                                        }
           Response:{
                              "stripeCustomerID": "cus_K96xmHH1uVXPrv",
                               "success": true
                            }
     2) POST /api/payments
       url: POST localhost:3008/api/payments
       Request : {
                   "stripeCustomerID": "cus_K8sPNXxVUhjAL1",
 	             "amount": 700
                      }
     Response:{
    "paymentIntentID": "pi_3JUopOSJuSXF3ti00bRZnZf4",
    "success": true
}


3) GET /api/payments/:customer_id
  url: GET localhost:3008/api/payments/cus_K8sPNXxVUhjAL1
  Request : {
 "stripeCustomerID": "cus_K8sPNXxVUhjAL1"
}
Response:{
    "payments": [
        {
            "id": "pi_3JUopOSJuSXF3ti00bRZnZf4",
            "amount": 700
        },
        {
            "id": "pi_3JUbI1SJuSXF3ti00kc2SAn3",
            "amount": 700
        },
        {
            "id": "pi_3JUbHmSJuSXF3ti00edNyuJQ",
            "amount": 700
        },
        {
            "id": "pi_3JUbFfSJuSXF3ti00Gnvet2k",
            "amount": 700
        }
    ],
    "success": true
}
 
----------------------------------------------------- 
 
Steps for running APIs(Heroku) server

1. import the postman collection into postman and you can simply hit the API


 note: Taking the API key from env file. No need for provide within header. 
List of APIs:
   1) POST /api/customer
          Url:-POST https://blooming-falls-44557.herokuapp.com/api/customer
          Request body :{
                                           "email":"token@gmail.com",
                                        }
           Response:{
                              "stripeCustomerID": "cus_K96xmHH1uVXPrv",
                               "success": true
                            }
     2) POST /api/payments
       url: POST https://blooming-falls-44557.herokuapp.com/api/payments
       Request : {
                   "stripeCustomerID": "cus_K8sPNXxVUhjAL1",
 	             "amount": 700
                      }
     Response:{
    "paymentIntentID": "pi_3JUopOSJuSXF3ti00bRZnZf4",
    "success": true
}


3) GET /api/payments/:customer_id
  url: GET https://blooming-falls-44557.herokuapp.com/api/payments/cus_K8sPNXxVUhjAL1
  Request : {
 "stripeCustomerID": "cus_K8sPNXxVUhjAL1"
}
Response:{
    "payments": [
        {
            "id": "pi_3JUopOSJuSXF3ti00bRZnZf4",
            "amount": 700
        },
        {
            "id": "pi_3JUbI1SJuSXF3ti00kc2SAn3",
            "amount": 700
        },
        {
            "id": "pi_3JUbHmSJuSXF3ti00edNyuJQ",
            "amount": 700
        },
        {
            "id": "pi_3JUbFfSJuSXF3ti00Gnvet2k",
            "amount": 700
        }
    ],
    "success": true
}

--------------------------

it took me 2 hours to complete the assignment 
