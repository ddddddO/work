function copyDate(dc, dt){
    var txt = dt.toLocaleString(), pre = dc.createElement('pre');
    console.log(txt);

    pre.style.userSelect = 'auto';
    pre.textContent = txt;

    dc.body.appendChild(pre);
    dc.getSelection().selectAllChildren(pre);
    dc.execCommand('copy');
    dc.body.removeChild(pre);
};
