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

