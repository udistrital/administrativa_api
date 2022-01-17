# administrativa crud api

Api REST que comunica con la base de datos respectiva con el fin de gestionar latos relacionados a los sistemas Argo y Agora principalmente

## Instalación

Para instalar el proyecto de debe relizar lo siguientes pasos:

Ejecutar desde la terminal 'go get repositorio':

```shell
go get github.com/udistrital/administrativa_crud_api
```

## Variables de Entorno

```sh
# 1. Copiar plantilla .env
cp template.env [COPIA].env

# 2. Editar .env
nano [COPIA].env
# ... puede ser con nano u otro editor de texto plano

# 3. Cargar variables del .env
source [COPIA].env
# Hacerlo cada que cambie la [COPIA].env o al abrir un nuevo terminal
```

## Ejecución del proyecto

- Ejecutar:

```shell
bee run
```

- O si se quiere ejecutar el swager:

```shell
bee run -downdoc=true -gendoc=true
```

## Licencia

This file is part of cumplidos-cliente.

cumplidos-cliente is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
