


function FindProxyForURL(url, host) {
  if (shExpMatch(url,"km.sankuai.com")) {
        return "PROXY 127.0.0.1:8080;DIRECT";
    }
}