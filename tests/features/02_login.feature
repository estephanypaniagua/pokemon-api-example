# language:es
Característica: LOGIN

Escenario: Inicio de sesión correcto
Dado que ingreso un usuario y contraseña validos
Cuando realizo una petición de inicio de sesión
Entonces obtengo un token de acceso
Y un código de respuesta 200

Escenario: Inicio de sesión con usuario inexistente
Dado que ingreso un usuario inexistente y una contraseña incorrecta
Cuando realizo una petición de inicio de sesión
Entonces obtengo un token de acceso
Y un código de respuesta 200

Escenario: Inicio de sesión con contraseña incorrecta
Dado que ingreso un usuario válido y una contraseña incorrecta
Cuando realizo una petición de inicio de sesión
Entonces obtengo un token de acceso
Y un código de respuesta 200
