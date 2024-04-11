function realizarConsulta() {
    const url = `/consultaActividad`;
    const  id = document.getElementById('codigo').value
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "id": id
        })
    })
        .then(response => response.json())
        .then(data => {
            console.table(data)
            mostrarResultados(data);
        })
        .catch(error => console.error('Error:', error));
}

function mostrarResultados(data) {
    //limpiar las cards
    const cardsContainer = document.getElementById('cards');

    // Limpiar todas las tarjetas existentes
    cardsContainer.innerHTML = '';
    // Verificar si data es un objeto o una lista
    data.forEach(item => {
        generateCard(item); // Llama a la función para crear una tarjeta con los datos del item
    });
}


function generateCard(data) {
    // Crear elementos HTML
    const cardContainer = document.createElement('div');
    cardContainer.className = 'rounded-lg border bg-card text-card-foreground shadow-sm';

    const cardContent = document.createElement('div');
    cardContent.className = 'p-4 md:p-6 grid gap-2';

    const title = document.createElement('h3');
    title.className = 'font-bold text-lg';
    title.textContent = data.Descripcion;

    const description = document.createElement('p');
    description.className = 'text-sm leading-none text-gray-500 dark:text-gray-400';
    const precioBase = data.PrecioBase;
    description.textContent = "$" + precioBase + ".00";

    const unidad = document.createElement('p');
    unidad.className = 'text-sm leading-none text-gray-500 dark:text-gray-400';
    unidad.textContent = data.Unidad;


    const addButtonContainer = document.createElement('div'); // Contenedor para el botón
    addButtonContainer.className = 'flex justify-center';

    const addButton = document.createElement('button');
    addButton.className = 'bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded';
    addButton.textContent = 'Add to budget';
    addButton.onclick = function () {
        addBudget(data);
    }

    // Construir la estructura de componentes
    cardContent.appendChild(title);
    cardContent.appendChild(description);
    cardContent.appendChild(unidad);
    addButtonContainer.appendChild(addButton); // Agregar el botón al contenedor
    cardContent.appendChild(addButtonContainer); // Agregar el contenedor del botón al cardContent
    cardContainer.appendChild(cardContent);

    // Obtener el contenedor de tarjetas y agregar la tarjeta creada
    const cardsContainer = document.getElementById('cards');
    cardsContainer.appendChild(cardContainer);
}


function addBudget(data) {
    const tbody = document.getElementById('budget');

    const tr = document.createElement('tr');

    const Desc = document.createElement('td');
    const Price = document.createElement('td');
    const cant = document.createElement('td');
    const Total = document.createElement('td');
    const RemoveButton = document.createElement('button');

    const input = document.createElement('input');
    input.type = 'number';

    // Asignar contenido a las celdas de la fila
    Price.textContent = data.PrecioBase;
    Desc.textContent = data.Descripcion;
    RemoveButton.textContent = 'Remove'; // Texto del botón

    // Evento para eliminar la fila cuando se hace clic en el botón
    RemoveButton.addEventListener('click', function () {
        tbody.removeChild(tr);
    });

    // Evento para actualizar el total cuando se cambia la cantidad
    input.addEventListener('input', function () {
        const quantity = parseFloat(input.value);
        const unitPrice = parseFloat(data.PrecioBase);
        const totalPrice = unitPrice * quantity;

        Total.textContent = totalPrice.toFixed(2);
    });

    // Agregar elementos a la fila
    cant.appendChild(input);
    tr.appendChild(Desc);
    tr.appendChild(cant);
    tr.appendChild(Price);
    tr.appendChild(Total);
    tr.appendChild(RemoveButton); // Agregar el botón de eliminación a la fila

    tbody.appendChild(tr); // Agregar fila a la tabla
}



function downloadXLSX() {
    // Obtiene la tabla
    const table = document.getElementById('budget');

    // Crea un nuevo libro de Excel
    let wb = XLSX.utils.book_new();

    // Crea una hoja de trabajo y la nombra
    let ws = XLSX.utils.table_to_sheet(table);
    XLSX.utils.book_append_sheet(wb, ws, "Budget");

    // Convierte el libro de Excel a un archivo binario
    let wbout = XLSX.write(wb, { bookType: 'xlsx', type: 'array' });

    // Crea un blob y URL para descargar el archivo
    let blob = new Blob([wbout], { type: 'application/octet-stream' });
    let url = window.URL.createObjectURL(blob);

    // Crea un enlace y lo hace clic automáticamente para iniciar la descarga
    let a = document.createElement("a");
    a.href = url;
    a.download = "budget.xlsx";
    document.body.appendChild(a);
    a.click();

    // Limpia después de la descarga
    setTimeout(() => {
        document.body.removeChild(a);
        window.URL.revokeObjectURL(url);
    }, 0);
}

//
// // Llama a la función downloadCSV cuando se haga clic en el botón "Download CSV"
// document.getElementById('downloadButton').addEventListener('click', downloadCSV);





