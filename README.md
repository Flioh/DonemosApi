##Donemos REST API

Back-end para la app Donemos.

###Como usar

0. Asegurate de tener el [Gopath configurado adecuadamente](https://golang.org/doc/code.html#Workspaces).

1. Cloná.

2. Ejecutá `go get`

3. Verifica que usas las variables de entorno requeridas, documentadas abajo.

4. Ejecutá `mongod`.

5. Usá [Gin](https://github.com/codegangsta/gin) para live-reload o simplemente `go run main.go` para ejecutar el servidor.


#### Variables de entorno

**MONGO_URL**: La URL a la db de mongo, si no se proporciona se usa `mongodb://localhost`.

**AUTH0_CLIENT_SECRET**: Llave para verificar JsonWebTokens.

**BUGSNAG_API_KEY**: API Key de Bugsnag



### API

#### Solicitudes de dadores

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /solicitud| []Solicitud | Todas las solicitudes|
| GET       | /solicitud/usuario/{usuarioID} | []Solicitud | Solicitues del usuario dado. |
| GET       | /solicitud/{número_de_página} | []Solicitud | Solicitudes correspondiente a `número_de_página`
| GET       | /solicitud/filtrar/{provinciaID}/{localidadID}/{grupoID}/{factorID} | []Solicitud | Busca solicitudes que tengan los atributos dados, pasar `null` en un atributo para ignorar ese atributo. Nota: Grupo y Factor siempre van en pares, si se recibe uno pero no el otro, se ignora el grupo/factor recibido.
| GET  | /solicitud/{número_de_página}/filtrar/{provinciaID}/{localidadID}/{grupoID}/{factorID} | []Solicitud | Paginación para el filtro solicitado.
| POST | /solicitud | En Error: 422, Creada: 201 | Solicita agregar la solicitud enviada a la base de datos.
| GET | /solicitud/{solicitudID} | Solicitud | La solicitud solicitada por id.
| PUT | /solicitud/{solicitudID} | Solicitud | Reemplaza la solicitud con id `solicitudID` por la enviada.
| DELETE | /solicitud/{solicitudID} | Eliminado: 202, No Encontrado: 204 | Elimina la solicitud con id `solicitudID`

#### Bancos de sangre

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /banco/{provinciaId}| []Banco | Todos los bancos dentro de provincia dada|
| GET       | /banco/{lat}/{lon}/{rango} | []Banco | Todos los bancos dentro de `rango` metros desde coordenadas dadas. |

#### Provincias

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /provincia| []Provincia | Todas las provincias|

#### Localidades

| Método    | Patrón    |   Resultado | Descripción |
|:----------|:----------|------------:|-------------|
| GET       | /localidad/{provinciaID}| []Localidad | Todas las localidades de la provincia con id `provinciaID`|
