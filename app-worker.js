const cacheName = "app-" + "9c48c5c3efabe3f0ff6b132e29515b1979c500b0";

self.addEventListener("install", event => {
  console.log("installing app worker 9c48c5c3efabe3f0ff6b132e29515b1979c500b0");
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
  console.log("app worker 9c48c5c3efabe3f0ff6b132e29515b1979c500b0 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
