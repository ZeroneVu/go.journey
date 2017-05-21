Abstract
--------
Once upon a time there was a hero less perfect than me call Gary Stu. Like all other RPG game hero of that time the not-so-young-anymore Gary decided to slay Dumbfuck Dragon. He started [go](https://golang.org/)ing out on a journey

Like all the journey ... it's started with a hello :)
Before starting any journey ... let's rest :)

Copy from this [article](https://thenewstack.io/make-a-restful-json-api-go/) for it simplicity. Following its implementation [github](https://github.com/corylanou/tns-restful-json-api)

I wanna create a very basic skeleton for simple RESTful service which matching my business requirements ... so the journey keep going further ...

Architecture
------------
![enter image description here](http://res.cloudinary.com/zeronevu/image/upload/v1495387671/simplerest.svg)

Vision
------
Simple is best and strait forward RESTful service seem cool and easy to understand but how about ...

 - Pre to route all traffics into handler ... how about all auth services (authentication vs authorization). ??!!
 - Post to route all traffics into handler ... what happen with the logging services or specific wrapping handing requests ??!!
 - Ideas on implementing middleware such as Mongo, Redis and the like ...

Running the example
-------------------

To run this exmaple, from the root of this project:
```sh
go run ./*.go
```
