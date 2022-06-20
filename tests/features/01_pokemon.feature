# language:es
Característica: OBTENER POKEMON

Escenario: Obtener todos los pokemones
Dado -
Cuando solicito los pokemon
Entonces obtengo una lista de los pokemon por id y nombre
Y un código de respuesta 200

Escenario: Obtener detalle de un pokemon con ID válido
Dado que tengo un ID válido
Cuando solicito un pokemon
Entonces obtengo las características del pokemon
Y un código de respuesta 200

Escenario: Obtener detalle de un pokemon con ID no existente
Dado que tengo un ID no existente
Cuando solicito un pokemon
Entonces obtengo un mensaje de error
Y un código de respuesta 404

Escenario: Crear un nuevo pokemon
Dado que tengo los datos de un nuevo pokemon
Cuando solicito su creación
Entonces obtengo las caracteristicas del pokemon creado
Y un codigo de respuesta 201

Escenario: Actualizar un pokemon
Dado que tengo un pokemon
Cuando quiero actualizar su informacion
Entonces obtengo las caracteristicas del pokemon actualizado
Y un codigo de respuesta 200

Escenario: Eliminar un pokemon
Dado que tengo un pokemon que quiero eliminar
Cuando envio su id
Entonces se elimina el pokemon
Y obtengo un codigo 204