let http = require("http");

if (typeof localStorage === "undefined" || localStorage === null) {
  var LocalStorage = require("node-localstorage").LocalStorage;
  localStorage = new LocalStorage("./visits");
}

if (localStorage.getItem("count") === null) {
  localStorage.setItem("count", 0);
}

const server = http
  .createServer(function(req, res) {
    if (req.url === "/favicon.ico") {
      return;
    }

    let userCount = localStorage.getItem("count");
    userCount++;
    localStorage.setItem("count", userCount);
    res.writeHead(200, { "Content-Type": "text/plain" });
    res.write("Czesc!\n");
    res.write(
      `Dales sie nabrac! Ta strona jest przykladem phishingu realizowanego na zajecia na Politechnice Poznanskiej \n`
    );
    res.write("Mielismy juz " + userCount + " wizyt!\n");
    res.end();
  })
  .listen(8888);
