function realizarConsulta() {
    const codigo = document.getElementById('codigo').value;
    const tipoItem = document.getElementById('tipoItem').value;
    const url = `/consulta?codigo=${codigo}&tipoItem=${tipoItem}`;
    fetch(url)
        .then(response => response.json())
        .then(data => {
            mostrarResultados(data);
        })
        .catch(error => console.error('Error:', error));
}

function mostrarResultados(data) {
    const tbody = document.querySelector('#tablaResultados tbody');
    tbody.innerHTML = '';

    // Verificar si data es un objeto o una lista
    if (Array.isArray(data)) {
        // Si data es una lista de objetos
        data.forEach(item => {
            const tr = document.createElement('tr');
            for (const key in item) {
                const td = document.createElement('td');
                td.textContent = item[key];
                tr.appendChild(td);
            }
            tbody.appendChild(tr);
        });
    } else {
        // Si data es un solo objeto
        const tr = document.createElement('tr');
        for (const key in data) {
            const td = document.createElement('td');
            td.textContent = data[key];
            tr.appendChild(td);
        }
        tbody.appendChild(tr);
    }
}

var inputCodigo = document.getElementById("codigo");
inputCodigo.addEventListener("keydown", function(event) {
    // Verifica si la tecla presionada es "Enter"
    if (event.keyCode === 13) {
        event.preventDefault(); // Evita que el formulario se envíe
        realizarConsulta(); // Llama a la función para realizar la búsqueda
    }
});

