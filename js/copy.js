function copyName(id) {
    let name = document.getElementById(id.id).innerText
    navigator.clipboard.writeText(name)
    alert("Copied name: " + name);
}
