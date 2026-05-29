# Modelo de Datos

## Usuario

Representa a una persona registrada en el sistema.

### Atributos

* id
* nombre
* apellido
* email
* password_hash
* rol
* fecha_creacion

### Roles

* CLIENTE
* ADMIN

---

## Evento

Representa un evento disponible para la venta de entradas.

### Atributos

* id
* titulo
* descripcion
* fecha
* ubicacion
* capacidad
* precio
* categoria
* imagen_url
* estado

### Estados

* ACTIVO
* CANCELADO
* FINALIZADO

---

## Entrada

Representa una entrada comprada por un usuario para un evento.

### Atributos

* id
* usuario_id
* evento_id
* fecha_compra
* estado

### Estados

* ACTIVA
* CANCELADA
* TRANSFERIDA

---

# Relaciones

## Usuario → Entrada

Un usuario puede tener muchas entradas.

Usuario 1 ---- N Entrada

## Evento → Entrada

Un evento puede tener muchas entradas.

Evento 1 ---- N Entrada

## Entrada

Cada entrada pertenece a:

* un usuario
* un evento


## Puntuación

Representa la calificación que un usuario realiza sobre un evento.

### Atributos

- id
- usuario_id
- evento_id
- valor
- comentario
- fecha_creacion

### Valores permitidos

- valor mínimo: 0
- valor máximo: 5

---

## Usuario → Puntuación

Un usuario puede realizar muchas puntuaciones.

Usuario 1 ---- N Puntuación