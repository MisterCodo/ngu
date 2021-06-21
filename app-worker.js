const cacheName = "app-" + "3aeed1ebd2ed34ff3c0b5adb43c8046102c98502";

self.addEventListener("install", event => {
  console.log("installing app worker 3aeed1ebd2ed34ff3c0b5adb43c8046102c98502");
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
  console.log("app worker 3aeed1ebd2ed34ff3c0b5adb43c8046102c98502 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
