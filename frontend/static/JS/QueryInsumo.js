function realizarConsulta() {
    const id = document.getElementById('codigo').value;
    const url = `/consultaActividad?id=${id}`;
    fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.table(data.data)
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
    if (data.hasOwnProperty('data')) {
        data.data.forEach(item => {
            generateCard(item); // Llama a la función para crear una tarjeta con los datos del item
        });
    } else {
        // Si data no es un objeto con la propiedad 'data', asumimos que es un solo objeto
        generateCard(data); // Llama a la función para crear una tarjeta con los datos del objeto
    }
}


function generateCard(data) {

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
    addButton.id = "openPopupButton"
    addButton.textContent = 'Add to budget';
    addButton.onclick = function () {
        openPopup(data.ID)
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

function generatePopup(id) {
    // Crear el contenedor del popup
    const url = `/consultaApu?id=${id}`;
    fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            const popupOverlay = document.createElement('div');
            popupOverlay.className = 'popup-overlay hidden';
            // popupOverlay.style.display = 'none';

            // Crear el contenedor del contenido del popup
            const popupContainer = document.createElement('div');
            popupContainer.className = 'popup-container rounded-lg border bg-white shadow-sm w-full max-w-4xl';

            // Crear el icono de cerrar
            const closeIcon = document.createElement('div');
            closeIcon.className = 'close-icon';
            closeIcon.innerHTML = `
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-6 w-6">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
        </svg>
    `;
            closeIcon.addEventListener('click', closePopup);

            // Crear el contenido del popup
            const popupContent = document.createElement('div');
            popupContent.className = 'flex flex-col space-y-1.5 p-6';

            const title = document.createElement('h3');
            title.className = 'whitespace-nowrap text-2xl font-semibold leading-none tracking-tight';
            title.textContent = 'Análisis de Precios Unitarios';

            const activityDescription = document.createElement('p');
            activityDescription.className = 'text-sm text-gray-500 desc';
            activityDescription.textContent = 'Actividad: ' + data.Descripcion;

            popupContent.appendChild(title);
            popupContent.appendChild(activityDescription);

            // Crear la tabla de contenido
            const tableContainer = document.createElement('div');
            tableContainer.className = 'p-6 grid gap-6';

            const tableGrid = document.createElement('div');
            tableGrid.className = 'grid grid-cols-2 gap-4';

            const labels = ['Unidad', 'Cantidad'];
            labels.forEach(labelText => {
                const labelDiv = document.createElement('div');
                labelDiv.className = 'flex flex-col';

                const label = document.createElement('label');
                label.className = 'text-sm font-medium leading-none';
                label.textContent = labelText;

                if (labelText === 'Unidad') {
                    const p = document.createElement('p');
                    p.id = 'unidad';
                    p.className = 'text-gray-500';
                    p.textContent = data.Unidad;
                    labelDiv.appendChild(label);
                    labelDiv.appendChild(p);
                } else if (labelText === 'Cantidad') {
                    const inputContainer = document.createElement('div');
                    inputContainer.className = 'flex items-center';

                    const input = document.createElement('input');
                    input.className = 'flex-1 h-10 rounded-md border border-gray-300 bg-gray-100 px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500 focus:outline-none';
                    input.type = 'number';

                    const button = document.createElement('button');
                    button.className = 'ml-2 inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium border bg-gray-100 text-gray-700 hover:bg-gray-200 h-10 px-4 py-2';
                    //button.onclick = addBudget();
                    button.innerHTML = `
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
                <path d="M5 12h14"></path>
            </svg>
            Crear memoria
        `;

                    inputContainer.appendChild(input);
                    inputContainer.appendChild(button);

                    labelDiv.appendChild(label);
                    labelDiv.appendChild(inputContainer);
                }

                tableGrid.appendChild(labelDiv);
            });


            tableContainer.appendChild(tableGrid);

            const tableWrapper = document.createElement('div');
            tableWrapper.className = 'relative w-full overflow-auto';

            const table = document.createElement('table');
            table.className = 'w-full caption-bottom text-sm';

            const thead = document.createElement('thead');
            thead.className = 'border-b';

            const theadRow = document.createElement('tr');
            ['Insumo', 'Cantidad', 'Unidad', 'Precio Unitario', 'Precio Total'].forEach(text => {
                const th = document.createElement('th');
                th.className = 'h-12 px-4 text-left align-middle font-medium text-gray-500';
                th.textContent = text;
                theadRow.appendChild(th);
            });

            thead.appendChild(theadRow);

            const tbody = document.createElement('tbody');
            data.Insumos.forEach(insumo => {
                const row = document.createElement('tr');
                ['Descripcion', 'Cantidad', 'Unidad', 'PrecioBase'].forEach(attr => {
                    const td = document.createElement('td');
                    td.className = 'px-4 py-2 text-gray-800';
                    td.textContent = insumo[attr]; // Aquí se asume que los atributos del objeto coinciden con los nombres de las cabeceras
                    row.appendChild(td);
                });
                tbody.appendChild(row);
            });

            // Agregar el tbody a la tabla existente
            table.appendChild(tbody);


            table.appendChild(thead);
            table.appendChild(tbody);
            tableWrapper.appendChild(table);

            tableContainer.appendChild(tableWrapper);

            // Botones
            const buttonContainerTop = document.createElement('div');
            buttonContainerTop.className = 'p-6 flex items-center justify-between';

            const selectChapterButton = document.createElement('button');
            selectChapterButton.className = 'flex h-10 w-full items-center justify-between rounded-md border border-gray-300 bg-gray-100 px-3 py-2 text-sm focus:ring-blue-500 focus:border-blue-500 focus:outline-none';
            selectChapterButton.innerHTML = `
        <span>Seleccionar capítulo</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="h-4 w-4">
            <path d="M6 9 12 15 18 9"></path>
        </svg>
    `;

            const newChapterButton = document.createElement('button');
            newChapterButton.className = 'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium border bg-gray-100 text-gray-700 hover:bg-gray-200 h-10 px-4 py-2';
            newChapterButton.textContent = 'Nuevo capítulo';

            buttonContainerTop.appendChild(selectChapterButton);
            buttonContainerTop.appendChild(newChapterButton);

            const buttonContainerBottom = document.createElement('div');
            buttonContainerBottom.className = 'p-6 flex items-center justify-between';

            const editButton = document.createElement('button');
            editButton.className = 'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium border bg-gray-100 text-gray-700 hover:bg-gray-200 h-10 px-4 py-2';
            editButton.textContent = 'Editar APU';

            const addButton = document.createElement('button');
            addButton.className = 'inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium bg-blue-500 text-white hover:bg-blue-600 h-10 px-4 py-2';
            addButton.textContent = 'Agregar al presupuesto';

            buttonContainerBottom.appendChild(editButton);
            buttonContainerBottom.appendChild(addButton);

            popupContainer.appendChild(closeIcon);
            popupContainer.appendChild(popupContent);
            popupContainer.appendChild(tableContainer);
            popupContainer.appendChild(buttonContainerTop);
            popupContainer.appendChild(buttonContainerBottom);

            popupOverlay.appendChild(popupContainer);

            document.body.appendChild(popupOverlay);
        })
        .catch(error => console.error('Error:', error));
}




// Función para abrir el popup
function openPopup(data) {
    generatePopup(data);
}

// Función para cerrar el popup
function closePopup() {
    const popupOverlay = document.querySelector('.popup-overlay');
    document.body.removeChild(popupOverlay);
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






