# Go-Microservices-UsersAPI


   APPLICATION is
   
    USERS_API------->MySql db
    
    DB details:
    Users schema
      Users table(Id,first_name,last_name,email,date_created)
    
    API documentation:
    
    1)GET specific user:
        Endpoint: "/users/<user_id>"
        Response:
          {
            "id": 1,
            "first_name": "Lwis",
            "last_name": "Ally",
            "email": "heeloee@gmail",
            "date_created": "2020-10-05T12:16:23Z"
          }
          
    2) POST user: Creates new user in db
         Endpoint: "/users"
        JSON Body:
          {
            "first_name": "Joey",
            "last_name": "Tribiani",
            "email": "joey.tribiani@gmail"
           }
         Response: Status 201 if created successfully & returns the created user
         
         {
            "id": 5,
            "first_name": "Joey",
            "last_name": "Tribiani",
            "email": "joey.tribiani@gmail",
            "date_created": "2020-10-05T15:37:04Z"
         }
         
           
    3)PUT user: Modifies existing user
      Endpoint: "/user/<user_id>"
      For ex,PUT "/user/5"
      JSON body:
        {
            "first_name": "Joey",
            "last_name": "Alan",
            "email": "joey.alan@gmail"
        }
        Response: Returns the updated JSON object
        {
            "id": 5,
            "first_name": "Joey",
            "last_name": "Alan",
            "email": "joey.alan@gmail",
            "date_created": "2020-10-05T15:37:04Z"
        }
      
      4)DELETE user:
        Endpoint: "/user/<user_id>"
        For ex,DELETE "/user/5"
        
        Response:
            { 
              "status": "deleted"
            }
