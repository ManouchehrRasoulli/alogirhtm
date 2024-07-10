const http = "http";
const https = "https";
const www = "www";

function domainName(url) {
    console.log(url);

    let sp = [];

    let us = url.split("://")
    for (let i = 0; i < us.length; i++) {
        uss = us[i].split(".");
        for (let j = 0; j < uss.length; j++) {
            sp.push(uss[j]);
        }
    }

    console.log(sp);
    for (let i = 0; i < sp.length; i++) {
        if (sp[i] !== http && sp[i] !== https && sp[i] !== www) {
            return sp[i];
        }
    }

    return "";
}

console.log(domainName("http://google.com"));
console.log(domainName("www.xakep.ru"));
console.log(domainName("http://google.co.jp"));
console.log(domainName("https://youtube.com"));