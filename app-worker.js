const cacheName = "app-" + "08e92a91a762fc0ca0a5e17e5c74c0b8cca4edae";

self.addEventListener("install", event => {
  console.log("installing app worker 08e92a91a762fc0ca0a5e17e5c74c0b8cca4edae");
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
  console.log("app worker 08e92a91a762fc0ca0a5e17e5c74c0b8cca4edae is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
