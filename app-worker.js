const cacheName = "app-" + "fd8d760cc8dc27844f42fe6682e1f6b4850ca2ae";

self.addEventListener("install", event => {
  console.log("installing app worker fd8d760cc8dc27844f42fe6682e1f6b4850ca2ae");
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
  console.log("app worker fd8d760cc8dc27844f42fe6682e1f6b4850ca2ae is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
