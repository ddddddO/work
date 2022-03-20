(function(dc, dt){
    var s = dc.createElement("script");
    s.src = "https://cdn.jsdelivr.net/gh/ddddddO/work@master/bookmarklet/copy_date/copy_date.js";
    dc.body.appendChild(s);
    copyDate(dc, dt);
    dc.body.removeChild(s);
}(document, new Date()));
