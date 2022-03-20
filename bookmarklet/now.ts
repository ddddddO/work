(function(){
    const date = new Date();
    alert(date.toLocaleString());

    const text = date.toLocaleString();
    const pre = document.createElement('pre');

    pre.style.userSelect = 'auto';
    pre.textContent = text;

    document.body.appendChild(pre);
    document.getSelection().selectAllChildren(pre);
    const result = document.execCommand('copy');

    document.body.removeChild(pre);
    return result;
  })()
