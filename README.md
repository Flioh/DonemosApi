##Donemos REST API

Back-end para la app Donemos.

###Como usar

1. Cloná.

2. Ejecutá `mongod`.

3. Usá [Gin](https://github.com/codegangsta/gin) para live-reload.

### API

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /solicitud| []Solicitud | Todas las solicitudes|
| GET       | /solicitud/{número_de_página} | []Solicitud | Solicitudes correspondiente a `número_de_página`
| GET       | /solicitud/filtrar/{provinciaID}/{localidadID}/{grupoID}/{factorID} | []Solicitud | Busca solicitudes que tengan los atributos dados, pasar `null` en un atributo para ignorar ese atributo.
| GET  | /solicitud/{número_de_página}/filtrar/{provinciaID}/{localidadID}/{grupoID}/{factorID} | []Solicitud | Paginación para el filtro solicitado.
| POST | /solicitud | En Error: 422, Creada: 201 | Solicita agregar la solicitud enviada a la base de datos.
| GET | /solicitud/{solicitudID} | Solicitud | La solicitud solicitada por id.
| PUT | /solicitud/{solicitudID} | Solicitud | Reemplaza la solicitud con id `solicitudID` por la enviada.
| DELETE | /solicitud/{solicitudID} | Eliminado: 202, No Encontrado: 204 | Elimina la solicitud con id `solicitudID`

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /provincia| []Provincia | Todas las provincias|

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /localidad/{provinciaID}| []Localidad | Todas las localidades de la provincia con id `provinciaID`|
