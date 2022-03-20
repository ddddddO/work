(function(dc, dt){
    var s = dc.createElement("script");
    s.src = "copy_date.js";
    dc.body.appendChild(s);
    copy_date(dc, dt);
    dc.body.removeChild(s);
}(document, new Date()));
