# Sistema de Gestión de Eventos y Entradas

## Integrantes

* Sofía Acuña
* Anna Garello Bertorello
* Leticia Vega Pereira

## Materia

Desarrollo de Software - UCC

## Descripción

Este proyecto consiste en el desarrollo de una aplicación web para la gestión de eventos y venta de entradas.

El sistema permitirá a los usuarios explorar eventos disponibles, consultar información detallada, adquirir entradas, administrar sus compras y transferir entradas a otros usuarios.

Además, los administradores podrán gestionar eventos y consultar reportes de ventas y ocupación.

Como funcionalidad adicional, se incorporará un sistema de puntuación mediante estrellas que permitirá a los usuarios valorar eventos y visualizar un ranking de los mejor calificados.

---

## Estado del Proyecto

Actualmente el proyecto se encuentra en fase de análisis y diseño.

Se encuentran definidos:

* Organización del proyecto.
* Modelo de datos.
* Requisitos funcionales.
* Diseño de endpoints.
* Decisiones de diseño iniciales.

---

## Tecnologías Utilizadas

### Backend

* Go
* Gin
* GORM

### Frontend

* React

### Base de Datos

* MySQL

### Seguridad

* JWT

### Infraestructura

* Docker
* Docker Compose

### Testing

* Go Testing
* Testify

---

## Funcionalidades Principales

### Cliente

* Registro e inicio de sesión.
* Exploración de eventos.
* Consulta de detalle de eventos.
* Compra de entradas.
* Consulta de entradas adquiridas.
* Cancelación de entradas.
* Transferencia de entradas.
* Puntuación de eventos.

### Administrador

* Creación de eventos.
* Modificación de eventos.
* Cancelación de eventos.
* Consulta de reportes de ventas.
* Consulta de reportes de ocupación.

---

## Funcionalidad Adicional

### Sistema de Puntuación

Los usuarios podrán calificar eventos utilizando una escala de 0 a 5 estrellas.

El sistema calculará el promedio de puntuaciones y permitirá visualizar un ranking de eventos ordenado según su valoración.

---

## Estructura del Proyecto

Proyecto-desarrollo-sw/
│
├── backend/
├── frontend/
├── docs/
│   ├── organizacion-proyecto.md
│   ├── modelo-datos.md
│   ├── decisiones-diseno.md
│   ├── endpoints.md
│   └── requisitos-funcionales.md
│
├── README.md
└── .gitignore


## Documentación

La documentación técnica y funcional se encuentra disponible dentro de la carpeta:

```txt
docs/
```

---

## Próximos Pasos

* Diseño del DER (Diagrama Entidad-Relación).
* Implementación del backend en Go.
* Desarrollo del frontend en React.
* Configuración de Docker Compose.
* Implementación de testing automatizado.
