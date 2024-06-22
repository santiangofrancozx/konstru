Guía de Instalación y Configuración del Proyecto
Esta guía detalla los pasos necesarios para configurar un entorno de desarrollo con Chocolatey, Git, Golang, y MySQL, así como las dependencias para los proyectos backend y frontend.

Proyecto Front-End
Para la instalación del front-end, dirígete al repositorio:

arduino
Copy code
https://github.com/santiangofrancozx/Konstru-Front
Instalación de Chocolatey
Para instalar Chocolatey, ejecuta el siguiente script en PowerShell con permisos de administrador:

powershell
Copy code
Set-ExecutionPolicy Bypass -Scope Process -Force;
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
Instalación de Git
Una vez que Chocolatey esté instalado, instala Git con el siguiente comando:

powershell
Copy code
choco install git
Instalación de Golang
Para instalar Golang, utiliza el siguiente comando:

powershell
Copy code
choco install golang
Instalación de MySQL
Para instalar MySQL, ejecuta el siguiente comando:

powershell
Copy code
choco install mysql
Configuración del Proyecto Backend
Dirígete al directorio raíz de tu proyecto backend y ejecuta:

bash
Copy code
go mod tidy
Esto descargará y limpiará las dependencias necesarias para el proyecto Go.

Configuración del Proyecto Frontend
Dirígete al directorio raíz de tu proyecto frontend y ejecuta:

bash
Copy code
npm install
Esto instalará todas las dependencias necesarias definidas en el archivo package.json.

Entendiendo la Configuración
Para configurar el DSN (Data Source Name) de la base de datos del backend, dirígete al archivo main.go ubicado en el directorio del backend y modifica la línea:

go
Copy code
config.SetDSN("root:safraroot@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local")
Lee los comentarios para ajustar el usuario y contraseña de la conexión según sea necesario.

Para cargar datos CSV de ejemplo a la base de datos, encuentra la línea comentada en el código:

go
Copy code
// Migrates.ImportDataFromCSVDB()
Descoméntala para ejecutar la carga de estos datos de prueba.

