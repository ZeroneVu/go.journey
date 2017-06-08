Abstract
--------
Once upon a time there was a hero less perfect than me call Gary Stu. Like all other RPG game hero of that time the not-so-young-anymore Gary decided to slay Dumbfuck Dragon. He started [go](https://golang.org/)ing out on a journey

Like all the journey ... it's started with a hello :).

 - Hello to how variable scope between packages and accessing between public and private asserts. Our poor Stu wasted whole day to find out that public assert can be accessed with capital letter on the first character for example:

    ```go
	type PublicMe struct {
		Value string
	}
    ```
Could be a public (among packets) assert struct vs with how private (among packets) struct

    ```go
	type privateme struct {
		value string
	}
    ```

 - Hello to how golang express lambda in a very neat way by following this [thread](https://stackoverflow.com/questions/31362044/anonymous-interface-implementation-in-golang) on stackoverflow by this [answer](https://stackoverflow.com/a/42182987/2540986). By learning that Stu was lead to dealing with unknown type, number of function parameters in golang, digging deep down by this [thread](https://stackoverflow.com/questions/34048861/can-golang-take-a-function-with-unknown-number-of-parameters-as-an-argument-to-a), on an very interesting [answer](https://stackoverflow.com/a/34062738/2540986). It's seem that best practice that we **SHOULD NOT** try to fight type system over a static type programming language like go. Best practice that even in lambda expression we should guarantee the static form of input (in parameters) and output (return statement). Of course we could overcome all obstacles by [reflect](https://blog.golang.org/laws-of-reflection), slicing and such but fabulous always come with a price so it's **not recommended**.

Architecture
------------

Vision
------
The possibilities are endless :). What the heck vision could have from a hello ... ha ha ha ha. This more or less a dummy codebase of mine
...

Running the example
-------------------

To run this exmaple, from the root of this project:
```sh
go run ./*.go
```
