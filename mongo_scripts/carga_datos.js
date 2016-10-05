var conn = new Mongo();
var db = conn.getDB("donemos");

var localidadesC = db.getCollection("localidades");
var provincias  = db.getCollection("provincias");

localidadesC.drop();
provincias.drop();

var json = JSON.parse(cat(`${pwd()}/mongo_scripts/provincias_con_localidades.json`));


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
