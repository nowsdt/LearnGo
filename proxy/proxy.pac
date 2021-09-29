


function FindProxyForURL(url, host) {

  //alert("url:" + url);
  //alert("host" +host);
  //alert(host.indexOf("pixel.sankuai.com"));
  let result = wrap(url, host);
  alert(url + ":" + result);
  alert(host + ":" + result);

  return result;
}

function wrap(url, host) {

  if (shExpMatch(host,"*api/operationHistory/log/*")) {
    return "PROXY 127.0.0.1:81200";
  }

  if (host.indexOf("pixel.sankuai.com") != -1) {
    return "DIRECT";
  }

  if (shExpMatch(host,"*sankuai.com*")) {
    return "PROXY 127.0.0.1:8888";
  }
  return "DIRECT";
}