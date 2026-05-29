# Requisitos Funcionales

## Introducción

El sistema de Gestión de Eventos y Entradas permitirá a usuarios clientes explorar eventos, adquirir entradas y administrar sus compras, mientras que los usuarios administradores podrán gestionar el catálogo de eventos y consultar métricas del sistema.

Además, el proyecto incorpora una funcionalidad adicional de puntuación de eventos mediante un sistema de valoración de 0 a 5 estrellas.

---

# RF01 - Registro de Usuario

El sistema deberá permitir el registro de nuevos usuarios clientes mediante correo electrónico y contraseña.

---

# RF02 - Inicio de Sesión

El sistema deberá permitir que un usuario autenticado acceda mediante credenciales válidas.

El sistema generará un token JWT para la gestión de sesión.

---

# RF03 - Exploración del Catálogo de Eventos

El sistema deberá permitir visualizar el listado de eventos disponibles.

El catálogo deberá ofrecer mecanismos de búsqueda y filtrado.

---

# RF04 - Consulta de Detalle de Evento

El sistema deberá permitir visualizar toda la información asociada a un evento específico.

La información incluirá:

* título
* descripción
* fecha
* horario
* duración
* ubicación
* categoría
* precio
* cupo disponible

---

# RF05 - Compra de Entradas

El sistema deberá permitir que un usuario autenticado adquiera una o más entradas para un evento.

El sistema deberá validar:

* existencia del evento
* disponibilidad de cupo
* estado del evento

---

# RF06 - Consulta de Mis Entradas

El sistema deberá permitir que cada usuario consulte exclusivamente las entradas que le pertenecen.

---

# RF07 - Cancelación de Entradas

El sistema deberá permitir cancelar una entrada previamente adquirida.

La cancelación deberá liberar el cupo correspondiente del evento.

---

# RF08 - Transferencia de Entradas

El sistema deberá permitir transferir la titularidad de una entrada a otro usuario registrado.

La operación deberá conservar la integridad de los datos.

---

# RF09 - Creación de Eventos

El sistema deberá permitir a los administradores registrar nuevos eventos.

---

# RF10 - Modificación de Eventos

El sistema deberá permitir a los administradores actualizar la información de eventos existentes.

---

# RF11 - Cancelación de Eventos

El sistema deberá permitir a los administradores cancelar eventos existentes.

Los eventos cancelados no deberán aceptar nuevas compras.

---

# RF12 - Reporte de Ocupación

El sistema deberá permitir a los administradores consultar métricas de ocupación de un evento.

La información deberá incluir:

* capacidad total
* entradas emitidas
* entradas disponibles
* porcentaje de ocupación

---

# RF13 - Reporte de Ventas

El sistema deberá permitir a los administradores consultar información de ventas de un evento.

La información deberá incluir:

* cantidad de entradas vendidas
* recaudación total
* usuarios compradores

---

# RF14 - Puntuación de Eventos

El sistema deberá permitir que usuarios autenticados califiquen eventos mediante una valoración de 0 a 5 estrellas.

Reglas:

* Cada usuario podrá puntuar un evento una sola vez.
* La puntuación mínima será 0.
* La puntuación máxima será 5.

---

# RF15 - Consulta de Eventos Mejor Valorados

El sistema deberá permitir visualizar un ranking de eventos ordenados según su puntuación promedio.

La información deberá mostrarse de mayor a menor valoración.

---

# Restricciones Generales

* Todas las operaciones sensibles requerirán autenticación mediante JWT.
* Las contraseñas deberán almacenarse utilizando mecanismos de hashing.
* La persistencia se realizará mediante MySQL y GORM.
* El sistema deberá implementar control de permisos para usuarios Cliente y Administrador.
