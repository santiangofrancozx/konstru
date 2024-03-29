
function realizarConsulta() {
    const codigo = document.getElementById('codigo').value;

    fetch(`/consulta?codigo=${codigo}`)
        .then(response => response.json())
        .then(data => {
    mostrarResultados(data);
    })
    .catch(error => console.error('Error:', error));
}

    function mostrarResultados(data) {
    const tbody = document.querySelector('#tablaResultados tbody');
    tbody.innerHTML = '';

    const tr = document.createElement('tr');
    for (const key in data) {
    const td = document.createElement('td');
    td.textContent = data[key];
    tr.appendChild(td);
}
    tbody.appendChild(tr);
}
