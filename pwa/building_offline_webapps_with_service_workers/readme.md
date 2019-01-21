Pluralsight course
Building offline webapps with service workers - https://app.pluralsight.com/library/courses/building-offline-web-apps-service-worker/table-of-contents
Nik Molnar - nikcodes.com

### Course Overview
Topics
- Aynchrony in JS
- Making Network requests
- Register, install and activate a service worker
- Cache API in service worker spec
- Common caching patterns

Goal of this course: You'll be able to enhance an existing web application with service workers

Prerequisites:
Basic JS and web development

### Chapter 2: Why Service Worker?

Why offline matters?
Lie Fi

Offlinefirst.org
Treat network as a progressive enhancement. Assume you don't have it, use when you can.

The importance of Offline: git

AppCache - Application Cache was introduced in HTML5 for offline.
Downsides of AppCache
- Files always come from AppCache, even if you're online
- The Appcache only updates if the content of the manifest itself has changed
- The AppCache is an additional cache, not a browser default
- Never ever ever far-future cache the manifest.

Jake Archibald one of the main editors of the spec of service workers

Service Workers:
Newest of the line of web workers that were introduced in the HTML5 timeframe.
- Dedicated Worker
- Shared Worker
- Service Worker

The way a worker works is way different than the standard web application. In the standard web application,
the HTML, CSS and the JS, they are all running in one process in your browser. What a worker allows us to is to move
the JS or a portion of it into its own process and then we can communicate from the other JS that's running in the
browser process over to that worker - that's a dedicated worker.

A shared worker, allows us to use one worker thread with many different processes, scripts, and tabs.

A service worker:

Progressive enhancement means that we can always offer a baseline experience, and then make it better
and better as user's devices and their capabilities get better.

"An escalator can never break - it can only become stairs.

You would never see an "Escalator Temporarily out of order" sign, just 'Escalator Temporarily Stairs.
Sorry for the convenience. We apologize for the fact that you can still get up there." - Mitch Headberg

#### How Service Workers Work?
Service workers work on top of the principle - progressive enhancement.

How web works when it has a network connection?
A request is made from the browser to the server, server does some processing and sends a result back.

But when we lose the n/w connection, any request we try to make to the server doesn't get there, and
this is where service workers come in. Service workers run in a separate process and allow all network
traffic to go through them first.

Service workers can take a request, process it and send the response back  without ever going to the
server or leaving the device. When online, service workers take a request, they can pass it off to the
server, the server processes, returns the response, and the service worker returns the response back
to the browser for rendering. It might take that response, cache it saving it for later. It might
synthesize its own response. It can really do whatever it wants to.

### Fundamentals of JS Promises

#### Threading
JS has always run as a single-threaded application model. There is one thread of execution, and any
function or method or unit of work that you execute happens one after the other. This model has some
pros and cons.

The biggest benefit of the model is, it is easy to reason the flow of execution in the application
and how state is changing over time.

The oppositve of single threaded processing model is multi threaded
processing model. In this model each thread still executes one unit of
work at a time, but because there are multiple threads we can spread
those units of work across them, and run them in paralle. This model
gives better performance. The challenge when all threads are chanding
the same state or variable.

#### Asynchrony
When we program asynchronously we get some of the benefits of
multithreading without the need to create another thread or a worker,
and deal with the coordination.

In synchronous execution, we know one unit of work must execute before
the next unit of work can be followed. This makes sense for the CPU, in
many cases the CPU is spent waiting eg: I/O or network.

In asynchronous world, we can do the upfront CPU work for the first task
and then we call out whatever I/O task needs to get done, after the
I/O task is done it will tell us - call us back - thus the term
callback. This is also where the word promis comes from. I promise I
will let you know when things are done.

#### Asynchrony in Javascript
1) The first technique to deal with asynchrony is Events.
```javascript
var ajax = someHelper();
ajax.success = function(response) {
// do something with the data
};
ajax.get('/my-soda-api');
```

2) Callbacks.
```javascript
var ajax = someHelper();
ajax.get('/my-sode-api', function(response) {
// do something with the data
});
```

Event is executed over and over again, where as callback is executed
once. The challenge with Callback is composability. Nested callbacks
inside of nested callbacks - leads to callback nightmare.

3) Promises
```javascript
var ajax = someHelper();
ajax.get('/my-sode-api')
.then(function(response) {
// do something with the data
});
```

When compared with callbacks, it looks promises add a little value. But
it gives composability and reusability. Promises also open doors to
async/await pattern.

#### What is a promise?
CommonJS - Goal of CommonJS was to build up the JS ecosystem for web
servers, desktop and command line apps as well as in browser apps. This
group also made proposals for Promises.

Promises/A - A promise is defined as an object that has a function as
the value for the property then: then(fulfilledHandler, errorHandler,
progressHandler). The spec is so simple it left to the implentor how
then is implemented. This resuled in many promise libraries.

Promises/A+ - Actual full detailed spec browsers implement.

Promises work like a state machine. They begin with Pending and then
they move either to Fulfilled with value or Rejected with Reason and
settled.

Pending -> Fulfilled/Rejected -> Settled

A promise can only succeed or fail once, it doesn't repeat its action.
It can't swithc from a rejected state to a fulfilled state.

Promise remebers it's state.

```
var ajax = someHelper();
ajax.get('/my-soda-api')
.then(onFulfilled, onRejected);
```

When we implement onFulfilled or onRejected callbacks, it's important we
do one of three things:
1) Return another promise
2) Return a value
3) Throw an error

`.catch` is only if error.

// array of promises
Promise.all([
  ajax.get('a-api'),
  ajax.get('b-api'),
  ajax.get('c-api'),
])
.then(allFulfilled, firstRejected);

If all promises fulfilled then values are passed.

If anyone of the promises rejects the first rejection that we get will
be passed to this then handler.

Promise.race([
  ajax.get('a-api'),
  ajax.get('b-api'),
  ajax.get('c-api'),
])
.then(firstFulfilled, firstRejected);

the first promise that fulfills will get passed, or the first one to
reject gets passed to firstRejected callback.

#### Creating Promises
* Battery
* StorageQuota
* Font Loading
* Web MIDI
* Streams
* Fetch
* Service Worker
* Cache

```javascript
function example() {
 if (successful) { return value}
 else { throw error }
}
```

```javascript
navigator.geolocation.getCurrentPosition(function(position){
  console.log(position.coords.latitude);
}, function(error){
  console.log(error.message);
});
```
https://codepen.io/goutham2027-1478716461/pen/REqqQv?editors=1010#0


### Chapter-4 Introducing the Fetch API

#### History of Ajax
Ajax - Asychrnous Javascript and XML.

XMLHttpRequest

why we need fetch api? lets look at xmlhttprequest
```javascript
var request = new XMLHttpRequest();

request.onreadystatechange = function() {
  if (request.readyState === 4) {
    if (request.status >= 200 && request.status <= 299)
      console.dir(JSON.parse(request.responseText));
      else
      console.error('there was an error');
  };
}

request.open('GET', 'https://api.example.com', true);
request.send();
```

Fetch - Browser API
* Modern networking API
* Built on promises
* A living standard - spec is updated when new features are needed.
* Available on window and worker objects

```javascript
fetch('https://api.example.com')
.then(response => {
  if(response.ok)
    return response;

    throw new Error('there was an error');
})
.then(response => response.json())
.then(response => console.dir(response) || response)
.catch(console.error);
```

Fetch Concepts
It centeres around 4 main concepts.
1) Request
  * URL
  * Method
  * Headers
  * Body
  * Context - Why the request is being made?
  * Referrer
  * ReferrerPolicy
  * Mode - Mode for the request. Network environment. Same origin or
    Cross Origin request
  * Credentials - cookies should be passed along with the request or
    not.
  * Redirect - Follow redirects manually or automatically
  * Integrity
  * Cache - the way caching works.
2) Response
  * URL
  * Headers
  * Body
  * Status
  * StatusText ( 200 == Ok, 404 = File not found)
  * Ok
  * Type - if the request is normal same origin, cross origin, opaque
    response.
  * Redirected - Boolean property.
  * UseFinalUrl - This url is the final url
3) Headers
Dictionary or structure of key/value pairs
  * append()
  * set()
  * delete()
  * get() # first value of header
  * getAll() # multi value header
  * has() # Headers.has('content-length')
  * entries()
  * keys() # gives an iterator of keys
  * values()
4) Body
  * arrayBuffer()
  * blob()
  * formData()
  * json()
  * text()
  * bodyUsed
  all body methods can be used only once. to use again we should use the
  following methods.
  * Request.clone()
  * Response.clone()

fetchAPI Get - https://codepen.io/goutham2027-1478716461/pen/Jwewbx?editors=1011
fetchAPIPost - https://codepen.io/goutham2027-1478716461/pen/VqVqPo?editors=1011

### Chapter 5: First Service Worker
isserviceworkerready - https://jakearchibald.github.io/isserviceworkerready/

Path to service workers
Registration of the service worker - Let the browser know that the website has a service worker
and that is should go and download.

Installation - Get the service worker up and running

Usage - Make the service worker do useful stuff

#### Registration
To register a service worker with the browser

```javascript
if('serviceWorker' in navigator) {
    navigator.serviceWorker
    .register('/scripts/sw.js')
    .catch(err => console.error('There is a problem', err));
}
```

#### Installation
states of the installation of brand new service worker
state: no service worker
action: install -> installing state
event.waitUntil() -> Error state
action: activate -> activated state
event.waitUntil() -> Error state

update existing/active service worker
active service worker
install -> installing -> waiting
                      -> error
activate -> activated state

#### Using a service worker
From activated state service worker can sit in idle. If it is idle for too long, browser will terminate to save resoruces.