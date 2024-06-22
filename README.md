# Guía de Instalación y Configuración del Proyecto

Esta guía detalla los pasos necesarios para configurar un entorno de desarrollo con Chocolatey, Git, Golang, y MySQL, así como las dependencias para los proyectos backend y frontend.

## Proyecto Front-End

Para la instalacion del front vaya a el repositorio:
```
    https://github.com/santiangofrancozx/Konstru-Front
```

## Instalación de Chocolatey

Para instalar Chocolatey, ejecuta el siguiente script en PowerShell con permisos de administrador:

```powershell
powershellCopy code
Set-ExecutionPolicy Bypass -Scope Process -Force;
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

```

## Instalación de Git

Una vez que Chocolatey esté instalado, instala Git con el siguiente comando:

```powershell
powershellCopy code
choco install git

```

## Instalación de Golang

Para instalar Golang, utiliza el siguiente comando:

```powershell
powershellCopy code
choco install golang

```

## Instalación de MySQL

Para instalar MySQL, ejecuta el siguiente comando:

```powershell
powershellCopy code
choco install mysql

```

## Configuración del Proyecto Backend

Dirígete al directorio raíz de tu proyecto backend y ejecuta:

```bash
bashCopy code
go mod tidy

```

Esto descargará y limpiará las dependencias necesarias para el proyecto Go.

## Configuración del Proyecto Frontend

Dirígete al directorio raíz de tu proyecto frontend y ejecuta:

```bash
bashCopy code
npm install

```

Esto instalará todas las dependencias necesarias definidas en el archivo `package.json`.
