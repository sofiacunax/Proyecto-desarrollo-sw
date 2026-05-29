# Decisiones de Diseño Iniciales

> Este documento registra las decisiones iniciales del proyecto. Estas decisiones pueden modificarse durante el desarrollo mediante consenso del equipo.

# 1. Roles del sistema

El sistema contará con dos tipos de usuarios:

## Cliente

Puede:

* Registrarse e iniciar sesión.
* Consultar eventos disponibles.
* Comprar entradas.
* Ver sus entradas adquiridas.
* Cancelar entradas según las reglas definidas.
* Transferir entradas a otro usuario.

## Administrador

Puede:

* Crear eventos.
* Modificar eventos existentes.
* Cancelar eventos.
* Consultar reportes de ventas y ocupación.

---

# 2. Compra de entradas

Un usuario podrá adquirir más de una entrada para un mismo evento.

Justificación:

Permite contemplar compras grupales y simplifica la experiencia de uso.

---

# 3. Capacidad de eventos

Cada evento tendrá una capacidad máxima definida.

El sistema no permitirá vender entradas cuando se alcance dicho límite.

La disponibilidad se calculará como:

capacidad_total - entradas_vendidas

---

# 4. Estados de eventos

Los eventos podrán encontrarse en alguno de los siguientes estados:

* ACTIVO
* CANCELADO
* FINALIZADO

### Reglas

* Solo los eventos activos podrán vender entradas.
* Los eventos cancelados permanecerán visibles para mantener el historial.
* Los eventos finalizados serán únicamente informativos.

---

# 5. Estados de entradas

Las entradas podrán encontrarse en alguno de los siguientes estados:

* ACTIVA
* CANCELADA
* TRANSFERIDA

Esto permitirá conservar trazabilidad de operaciones realizadas por los usuarios.

---

# 6. Transferencia de entradas

Se permitirá transferir una entrada a otro usuario registrado.

La transferencia:

* Mantendrá el historial de la entrada.
* Cambiará el propietario de la misma.
* Quedará registrada en el sistema.

---

# 7. Cancelación de entradas

Un usuario podrá cancelar únicamente sus propias entradas.

La cancelación:

* Liberará un lugar para el evento.
* Actualizará el estado de la entrada a CANCELADA.

---

# 8. Seguridad

La autenticación se realizará mediante JWT.

Las contraseñas no se almacenarán en texto plano.

Se almacenarán utilizando algoritmos de hash seguros.

---

# 9. Gestión de eventos

Los administradores podrán:

* Crear nuevos eventos.
* Editar información de eventos existentes.
* Cancelar eventos.

Los eventos cancelados no podrán recibir nuevas compras.

---

# 10. Reportes

El sistema ofrecerá reportes básicos para administradores:

* Cantidad de entradas vendidas por evento.
* Porcentaje de ocupación.
* Cantidad de entradas disponibles.

---

# 11. Persistencia

La información se almacenará en MySQL utilizando GORM como ORM.

Todas las operaciones críticas que involucren múltiples cambios en la base de datos deberán ejecutarse dentro de transacciones.

---

# 12. Arquitectura

El backend seguirá una arquitectura por capas:

Controller → Service → Repository → Base de Datos

Las reglas de negocio residirán en la capa Service.

La capa Repository tendrá únicamente responsabilidades de persistencia.

---

# 13. Dockerización

El entorno se ejecutará mediante Docker Compose.

Los servicios iniciales serán:

* Backend (Go)
* Base de datos MySQL
* Frontend React

La configuración sensible se almacenará mediante variables de entorno.

# 14. Puntuación de eventos

Como funcionalidad adicional del grupo, se permitirá que los usuarios puntúen eventos con una valoración de 0 a 5 estrellas.

Reglas:

- Solo usuarios autenticados podrán puntuar eventos.
- La puntuación mínima será 0.
- La puntuación máxima será 5.
- Cada usuario podrá puntuar una vez por evento.
- El sistema calculará el promedio de puntuaciones de cada evento.
- Existirá una vista con los eventos mejor rankeados.