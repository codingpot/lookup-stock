const cacheName = "app-" + "d6e60dcae5f665b803c637c7aa8bbedb98e6afd7";

self.addEventListener("install", event => {
  console.log("installing app worker d6e60dcae5f665b803c637c7aa8bbedb98e6afd7");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/lookup-stock",
          "/lookup-stock/app.css",
          "/lookup-stock/app.js",
          "/lookup-stock/manifest.webmanifest",
          "/lookup-stock/wasm_exec.js",
          "/lookup-stock/web/app.wasm",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
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
  console.log("app worker d6e60dcae5f665b803c637c7aa8bbedb98e6afd7 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
