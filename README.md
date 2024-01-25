My Favorite Artist (Go Backend Engineer) - Developer

Note: I am using go-chi framework(https://medium.com/@jitenderkmr/building-scalable-micro-services-with-go-golang-and-chi-framework-6db5f2f9ad28)

I have created 3 seconadry services like
1). toptrackservice:
        a). hosted in 3001 and its take country as input parameter.
        b). it pulls the data and store in track object.
2). artistservice:
        a). hosted in 3002 and its takes artist name as input parameter.
        b). it pulls the data and store in artist object
3).suggestingservice:
        a). its hosted in 3003 and its take artist name and suggestion as input paramete.
        b). its pulls the data and store in suggestion object.

Now comes with primary services where it will communicate the seconadry services. i made very lossly couple so that it won't depend on the other services.

how to run:
1). please start all the seconadry services.
     Note: plz use "go run main.go > response" 

Now comes primary service its under src/go-microservice/main.go.
its hosted in 3000, in the home page it will guide you with other uri's.