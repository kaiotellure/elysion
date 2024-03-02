function success(proxy, message) {
  proxy.$data.notifications.push({
    success: true, message
  });
}

function error(proxy, message) {
  proxy.$data.notifications.push({
    error: true, message
  });
}

// proxy is an alpine data proxy, body is proxy.data stringified
// if successful inserts a notification into proxy.$data.notifications
// if failed does the same but notification.type = "error"
function put(proxy, url = "") {
    return fetch(url, {
      method: "put",
      body: JSON.stringify(proxy.data),
      headers: {
        "Content-Type": "application/json"
      }
    }).then(async response => {
      const text = await response.text();
      response.ok ? success(proxy, text) : error(proxy, response.status);
    }).catch(reason => {
      console.error(reason);
      error(proxy, reason);
    });
}

function blockwhile(element, promise) {
  promise.then(r => element.style.pointerEvents = "");
  element.style.pointerEvents = "none";
}
