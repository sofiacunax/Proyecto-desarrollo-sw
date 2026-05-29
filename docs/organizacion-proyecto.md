# Organización del Proyecto

## Nombre del proyecto

Sistema de Gestión de Eventos y Entradas

## Objetivo

Desarrollar una aplicación web tipo Ticketek que permita a usuarios clientes explorar eventos, comprar entradas, ver sus entradas, cancelarlas o transferirlas; y a usuarios administradores gestionar eventos y consultar reportes de ventas u ocupación.

## Tecnologías principales

* Backend: Go
* ORM: GORM
* Base de datos: MySQL
* Frontend: React
* Autenticación: JWT
* Contenedores: Docker y Docker Compose
* Control de versiones: GitHub con GitFlow

## Estructura del repositorio

```txt
Proyecto-desarrollo-sw/
├── backend/
├── frontend/
├── docs/
├── docker-compose.yml
├── .env.example
├── README.md
└── .gitignore
```

## Ramas de Git

* main: versión estable del proyecto.
* develop: rama de integración del equipo.
* feature/nombre-tarea: ramas individuales para cada funcionalidad.

## Flujo de trabajo

1. Cada integrante parte desde develop.
2. Cada tarea se desarrolla en una rama feature.
3. Al terminar, se sube la rama a GitHub.
4. Se crea un Pull Request hacia develop.
5. Otra integrante revisa el código antes de mergear.
6. Cuando develop esté estable, se integra a main.

## Funcionalidades principales

### Cliente

* Ver catálogo de eventos.
* Filtrar eventos.
* Ver detalle de un evento.
* Comprar entrada.
* Ver mis entradas.
* Cancelar entrada.
* Transferir entrada a otro usuario.

### Administrador

* Crear evento.
* Editar evento.
* Cancelar o eliminar evento.
* Consultar reporte de ventas u ocupación.

## Entidades iniciales

* Usuario
* Evento
* Entrada

## Próximos pasos

1. Definir atributos completos de cada entidad.
2. Definir relaciones entre entidades.
3. Diseñar endpoints REST.
4. Crear estructura interna del backend.
5. Crear estructura inicial del frontend.
6. Configurar Docker Compose.
