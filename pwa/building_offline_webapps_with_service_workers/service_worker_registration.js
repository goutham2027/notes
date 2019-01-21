if('serviceWorker' in navigator) {
    navigator.serviceWorker
    .register('/scripts/sw.js')
    .then(r => console.log("SW registered"))
    .catch(err => console.error('There is a problem', err));
}