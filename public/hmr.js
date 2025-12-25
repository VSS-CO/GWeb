const evtSource = new EventSource("/__hmr");
evtSource.onmessage = function(e) {
    if (e.data === "reload") {
        console.log("[GWeb HMR] Reloading page...");
        location.reload();
    }
};
