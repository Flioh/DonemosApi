var conn = new Mongo();
var db = conn.getDB("donemos");

var localidadesC = db.getCollection("localidades");
var provincias  = db.getCollection("provincias");

load(`${pwd()}/mongo_scripts/provincias_con_localidades.js`);


for (var i = 0; i < json.length; i++) {
  var provincia = json[i];

  printjson(provincia.nombre);
  var id = ObjectId();
  provincias.insert({
    _id: id,
    nombre: provincia.nombre
  });

  var localidades = provincia.ciudades;
  for (var y = 0; y < localidades.length; y++) {
    var localidad = localidades[y];
    printjson(localidad.nombre);

    localidadesC.insert({
      provinciaId: id,
      nombre: localidad.nombre
    });
  }
}
