const version = 4;

self.addEventListener('install', function(event) {
    console.log("SW v%s Installed at", version, new Date().toLocaleDateString());
    self.skipWaiting();
});


self.addEventListener('activate', function(event) {
    console.log("SW v%s Activated at", version, new Date().toLocaleDateString());
});

self.addEventListener('fetch', function(event) {
    if(!navigator.onLine) {
        event.respondWith(new Response('<h1> Offline </h1>', { headers: {'Content-Type': 'text/html'}}))
    } else {
        console.log(event.request.url);
        event.respondWith(fetch(event.request));
    }
});