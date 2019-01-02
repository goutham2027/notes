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
