const version = 'v2';

// self.addEventListener('install', function(event) {
//     console.log("SW v%s Installed at", version, new Date().toLocaleDateString());
//     self.skipWaiting();
// });


// self.addEventListener('activate', function(event) {
//     console.log("SW v%s Activated at", version, new Date().toLocaleDateString());
// });

// self.addEventListener('fetch', function(event) {
//     if(!navigator.onLine) {
//         event.respondWith(new Response('<h1> Offline </h1>', { headers: {'Content-Type': 'text/html'}}))
//     } else {
//         console.log(event.request.url);
//         event.respondWith(fetch(event.request));
//     }
// });

self.addEventListener('install', function(event) {
    event.waitUntil(
        caches.open(version)
        .then(function(cache) {
            return cache.addAll(['/scripts/sw.js'])
        })
    )
});

self.addEventListener('fetch', function(event) {
    event.respondWith(
        caches.match(event.request)
        .then(function(res) {
            if(res)
                return res;

            if(!navigator.onLine)
                return caches.match(new Request('/scripts/sw.js'))

            return fetch(event.request);
        })
    )
})


// self.addEventListener('activate', function(event) {
//     console.log("SW v%s Activated at", version, new Date().toLocaleDateString());
// });

// self.addEventListener('fetch', function(event) {
//     if(!navigator.onLine) {
//         event.respondWith(new Response('<h1> Offline </h1>', { headers: {'Content-Type': 'text/html'}}))
//     } else {
//         console.log(event.request.url);
//         event.respondWith(fetch(event.request));
//     }
// });
