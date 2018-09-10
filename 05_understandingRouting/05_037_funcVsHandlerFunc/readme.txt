type handler: which is an interface it's got to serve HVM method where the response writer and a pointer to a request

PS1 know that type handler interface piece to listen and serve takes a handler.

Got that boom piece three we want to do routing the the way we do routing is with a mux a multiplexer or serve mux.

So we can do news server mux gives us back a type pointer to a server mux with pointer to serve mux

Yes hallelujahs got HP response right a pointer to a request type serve Mock's any value of type served

Mock's implements the handler interface suite.

So when we create a server mocs we can pass that in listen and serve because listeners are once a handler

Yes type pointer to serve Mock's also has several other methods attached to it to them handle and handle

Falk handle takes a handler.

We saw how that works handle fog takes a function with the parameters response writer pointer to request

that is not type handler that is a function that's its own type.

It's a function with the response writer and a pointer to a request.

Ok cool so we can use Handle fonk and handle fonk is going to take a route and then a function.

We're going to pass a function in there.

Some people even put it right there which looks too much like javascript for my flavor.

My preference.

Right but you pass a function in.

That's got the parameters response writer pointer to a request.

Cool.

We can also use default serve mux as next piece in the package list and serve pasand nil for the handler.

The program will automatically the standard library will automatically use the default server marks

and we can attach a handle or a handle func.

We could use those methods to attach routes.

Right.

And so handle or handle on handle fonk is by far the most commonly used and handle fog.

We again pass in the route and then a function with the response writer and pointer to request as parameters.

Cool Doug.

Got it solid.

No all of that.

And then you go and you look at the standard library and you're like OK handle fog takes a function