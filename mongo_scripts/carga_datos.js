var conn = new Mongo();
var db = conn.getDB("donemos");

var localidadesC = db.getCollection("localidades");
var provincias  = db.getCollection("provincias");
var bancos = db.getCollection("bancos")

localidadesC.drop();
provincias.drop();
bancos.drop();

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

print("Provincias y ciudades cargadas\n")

var json = JSON.parse(cat(`${pwd()}/mongo_scripts/bancos.json`));

for (var i = 0; i < json.length; i++) {
  var banco = json[i];
  printjson(banco.institucion);
  
  var id = ObjectId();
  bancos.insert({
    _id: id,
    provinciaId: ObjectId(banco.provincia),
    ciudad: banco.nombreCiudad,
    institucion: banco.institucion,
    direccion: banco.direccion,
    telefono: banco.telefono,
    horario: banco.horario,
    lat: banco.lat,
    lon: banco.lon
  });
}
