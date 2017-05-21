Abstract
--------
Once upon a time there was a hero less perfect than me call Gary Stu. Like all other RPG game hero of that time the not-so-young-anymore Gary decided to slay Dumbfuck Dragon. He started [go](https://golang.org/)ing out on a journey

Like all the journey ... it's started with a [hello](https://github.com/ZeroneVu/go.journey/tree/master/hello) :)
Before starting any journey ... let's [rest](https://github.com/ZeroneVu/go.journey/tree/master/simplerest) :)
Like old time RPG characters & movies, @tavern is the base of all quests :) ... Before building a tavern ... let's say that we have a minibar, drinks are served now :)

Copy from this [article](https://thenewstack.io/make-a-restful-json-api-go/) for it simplicity. Following its implementation [github](https://github.com/corylanou/tns-restful-json-api)

Then after that I've read this [article](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81) for the idea about how to add middleware in a cool way to source base of a RESTful service. And a very well implementation of that idea could be found in this [github](https://github.com/alecholmez/http-server)

I wanna create a very basic skeleton for simple RESTful service which matching my business requirements ... so the journey keep going further ...

Architecture
------------
![enter image description here](http://res.cloudinary.com/zeronevu/image/upload/v1495387671/minibar.svg)

Vision
------
Well ... it's human nature to want more vs my architecture design concept (taking out as much as possible until nothing left)

 - What happen if the service is not just RESTful anymore. For instance WAMP (web application message protocol such as [Turnpike](https://github.com/jcelliott/turnpike)) ...
 - I may want to turn this thing into even-base with many pub/sub services could run asynchronously ...
 - With long pooling techniques and the likes ...
 - Until next time.

Running the example
-------------------

To run this exmaple, from the root of this project:
```sh
go run ./*.go
```
