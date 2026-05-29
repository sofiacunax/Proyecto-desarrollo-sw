# Diseño de API REST - Sistema de Gestión de Eventos y Entradas

## Objetivo

Este documento define los endpoints principales que utilizarán el frontend y el backend para comunicarse.

La API permitirá gestionar usuarios, eventos, entradas, compras, cancelaciones, transferencias y reportes.

## Base URL

```txt
/api/v1
```

---

# Autenticación

## Registrar usuario

```http
POST /api/v1/auth/register
```

Permite crear una cuenta de usuario cliente.

### Datos requeridos

* nombre
* apellido
* email
* password

---

## Iniciar sesión

```http
POST /api/v1/auth/login
```

Permite autenticar un usuario y obtener un token JWT.

### Datos requeridos

* email
* password

### Respuesta esperada

* token JWT
* datos básicos del usuario
* rol del usuario

---

# Eventos

## Listar eventos

```http
GET /api/v1/events
```

Permite obtener el catálogo de eventos.

No requiere autenticación.

### Filtros posibles

* categoría
* disponibilidad
* fecha
* texto de búsqueda

---

## Ver detalle de evento

```http
GET /api/v1/events/{id}
```

Permite obtener toda la información de un evento específico.

No requiere autenticación.

---

## Crear evento

```http
POST /api/v1/events
```

Permite crear un nuevo evento.

Requiere autenticación y rol ADMIN.

### Datos requeridos

* título
* descripción
* fecha
* horario
* duración
* ubicación
* cupo
* precio
* categoría
* imagen_url

---

## Actualizar evento

```http
PUT /api/v1/events/{id}
```

Permite modificar los datos de un evento existente.

Requiere autenticación y rol ADMIN.

---

## Cancelar evento

```http
DELETE /api/v1/events/{id}
```

Permite cancelar o dar de baja un evento.

Requiere autenticación y rol ADMIN.

---

# Entradas

## Comprar entrada

```http
POST /api/v1/tickets
```

Permite comprar una o más entradas generales para un evento.

Requiere autenticación.

### Datos requeridos

* evento_id
* cantidad

### Validaciones

* El usuario debe estar autenticado.
* El evento debe existir.
* El evento debe estar activo.
* Debe haber cupo disponible.
* No se manejan asientos numerados.

---

## Ver mis entradas

```http
GET /api/v1/tickets/my
```

Permite obtener todas las entradas adquiridas por el usuario autenticado.

Requiere autenticación.

---

## Ver detalle de entrada

```http
GET /api/v1/tickets/{id}
```

Permite obtener información de una entrada específica.

Requiere autenticación.

### Validaciones

* El usuario solo puede ver entradas propias.
* Un administrador puede consultar entradas para gestión o reporte.

---

## Cancelar entrada

```http
PUT /api/v1/tickets/{id}/cancel
```

Permite cancelar una entrada comprada.

Requiere autenticación.

### Validaciones

* La entrada debe existir.
* La entrada debe pertenecer al usuario autenticado.
* La entrada debe estar activa.
* Al cancelarse, se libera el cupo del evento.

---

## Transferir entrada

```http
PUT /api/v1/tickets/{id}/transfer
```

Permite transferir una entrada a otro usuario registrado.

Requiere autenticación.

### Datos requeridos

* nuevo_usuario_id

### Validaciones

* La entrada debe existir.
* La entrada debe pertenecer al usuario autenticado.
* La entrada debe estar activa.
* El usuario destino debe existir.
* La entrada cambia de titular de forma íntegra.

---

# Reportes de Administrador

## Reporte de ocupación de evento

```http
GET /api/v1/admin/events/{id}/occupancy
```

Permite obtener métricas de ocupación de un evento.

Requiere autenticación y rol ADMIN.

### Información esperada

* capacidad total
* entradas vendidas
* entradas disponibles
* porcentaje de ocupación

---

## Reporte de ventas de evento

```http
GET /api/v1/admin/events/{id}/sales
```

Permite obtener métricas de ventas de un evento.

Requiere autenticación y rol ADMIN.

### Información esperada

* cantidad de entradas vendidas
* recaudación total
* listado de usuarios compradores

# Puntuaciones

## Puntuar evento

```http
POST /api/v1/events/{id}/ratings

---

# Códigos de estado esperados

## 200 OK

La operación se realizó correctamente.

## 201 Created

El recurso fue creado correctamente.

## 400 Bad Request

Los datos enviados son inválidos o incompletos.

## 401 Unauthorized

El usuario no está autenticado o el token es inválido.

## 403 Forbidden

El usuario está autenticado, pero no tiene permisos suficientes.

## 404 Not Found

El recurso solicitado no existe.

## 409 Conflict

La operación no puede realizarse por una regla de negocio.

Ejemplos:

* evento sin cupo disponible
* entrada ya cancelada
* evento cancelado

## 500 Internal Server Error

Error inesperado del servidor.

---

# Consideraciones

* El frontend nunca accederá directamente a la base de datos.
* Toda operación sensible deberá validarse en el backend.
* Las operaciones de compra, cancelación y transferencia deberán ser transaccionales.
* La API no contempla asientos numerados.
* Las entradas son generales y se controlan mediante cupo del evento.
