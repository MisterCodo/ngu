const cacheName = "app-" + "343ad0baa8e599d3041a41377d6f0276b454f71f";

self.addEventListener("install", event => {
  console.log("installing app worker 343ad0baa8e599d3041a41377d6f0276b454f71f");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/ngu",
        "/ngu/app.css",
        "/ngu/app.js",
        "/ngu/manifest.webmanifest",
        "/ngu/wasm_exec.js",
        "/ngu/web/app.wasm",
        "https://storage.googleapis.com/murlok-github/icon-192.png",
        "https://storage.googleapis.com/murlok-github/icon-512.png",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 343ad0baa8e599d3041a41377d6f0276b454f71f is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
