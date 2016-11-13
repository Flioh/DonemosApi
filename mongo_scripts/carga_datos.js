var conn = new Mongo();
var db = conn.getDB("donemos");

var localidadesC = db.getCollection("localidades");
var provinciasC = db.getCollection("provincias");
var bancosC = db.getCollection("bancos");

localidadesC.drop();
provinciasC.drop();
bancosC.drop();

var json = JSON.parse(cat(`${pwd()}/mongo_scripts/datos.json`));

for (var i = 0; i < json.length; i++) {
  var provincia = json[i];

  var cantidadLocalidades = provincia.ciudades && provincia.ciudades.length ? provincia.ciudades.length : 0;
  var cantidadBancos = provincia.bancos && provincia.bancos.length ? provincia.bancos.length : 0;
  printjson(provincia.nombre + ' | Localidades: ' + cantidadLocalidades + ' | Bancos de Sangre: ' + cantidadBancos);
  
  // Carga de provincias
  var provinciaId = ObjectId();
  
  provinciasC.insert({
    _id: provinciaId,
    nombre: provincia.nombre
  });

  // Carga de localidades
  var localidades = provincia.ciudades;
  
  for (var y = 0; y < localidades.length; y++) {
    var localidad = localidades[y];
    
    localidadesC.insert({
      provinciaId: provinciaId,
      nombre: localidad.nombre
    });
  }

  // Carga de bancos de sangre  
  var bancos = provincia.bancos;

  if(provincia.bancos) {
    for(var y = 0; y < bancos.length; y++) {
      var banco = bancos[y];

      // Obtenemos la lat y la lon del string coordenadas
      banco.lat = getLat(banco.coordenadas);
      banco.lon = getLon(banco.coordenadas);

      var id = ObjectId();
      bancosC.insert({
        _id: id,
        provinciaId: provinciaId,
        ciudad: procesarString(banco.nombreCiudad),
        institucion: procesarString(banco.institucion),
        direccion: procesarString(banco.direccion),
        telefono: banco.telefono,
        horario: banco.horario,
        loc: { type: "Point", coordinates: [ banco.lon, banco.lat ] },
        lat: banco.lat,
        lon: banco.lon
      });
    }

    bancosC.createIndex( { loc: "2dsphere" } );
  }
}

// Método que convierte un string a formato Oracion y reemplaza algunas abreviaciones o letras
function procesarString(texto){
  return texto.split(' ').map(e=> e.charAt(0).toUpperCase() + e.toLowerCase().slice(1)).join(' ')
                            .replace(" De ", " de ")
                            .replace(" A ", " a ")
                            .replace("Y", "y");
}

// Método que obtiene la latitud de un string de coordenadas
function getLat(coordenadas) {
  return coordenadas.replace(" ", "").split(",")[0];
}

// Método que obtiene la longitud de un string de coordenadas
function getLon(coordenadas) {
  return coordenadas.replace(" ", "").split(",")[1];
}
