const cacheName = "app-" + "6192819c4fd8f0bf6b4df43eb1956d95caa4fc3a";

self.addEventListener("install", event => {
  console.log("installing app worker 6192819c4fd8f0bf6b4df43eb1956d95caa4fc3a");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "",
        "/ngu",
        "/ngu/app.css",
        "/ngu/app.js",
        "/ngu/manifest.webmanifest",
        "/ngu/wasm_exec.js",
        "/ngu/web/SpeedKnight.png",
        "/ngu/web/app.wasm",
        
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
  console.log("app worker 6192819c4fd8f0bf6b4df43eb1956d95caa4fc3a is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
