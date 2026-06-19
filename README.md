# Eventia - Sistema de Gestión de Eventos y Entradas

Proyecto integrador de Desarrollo de Software de la Universidad Católica de Córdoba. Eventia es una aplicación web para explorar eventos, comprar y administrar entradas, transferirlas a otros usuarios y puntuar eventos. También dispone de operaciones protegidas para que los administradores gestionen el catálogo.

## Integrantes

- Sofía Acuña
- Anna Garello Bertorello
- Leticia Vega Pereira

## Funcionalidades principales

### Clientes

- Registro e inicio de sesión con JWT.
- Consulta y búsqueda de eventos.
- Compra de una o más entradas.
- Consulta, cancelación y transferencia de entradas propias.
- Puntuación de eventos y consulta del ranking.

### Administradores

- Creación de eventos.
- Actualización de eventos existentes.
- Eliminación de eventos.

## Tecnologías utilizadas

### Backend

- Go 1.26.1.
- Gin para la API HTTP.
- `database/sql` y el driver de MySQL.
- JWT para autenticación y autorización.
- bcrypt para el hash de contraseñas.

### Frontend

- React 19.
- React Router.
- Vite 8.
- ESLint.

### Persistencia y testing

- MySQL.
- Paquete estándar `testing` de Go.
- Dobles SQL incluidos en los tests, sin requerir una base de datos activa para ejecutar la suite.

## Estructura del proyecto

```text
Proyecto-desarrollo-sw/
├── backend/
│   ├── controllers/   # Recepción de solicitudes y respuestas HTTP
│   ├── services/      # Reglas de negocio
│   ├── dao/           # Acceso a MySQL
│   ├── models/        # Entidades del dominio
│   ├── dtos/          # Objetos de entrada y salida de la API
│   ├── utils/         # JWT y middlewares
│   ├── db/            # Conexión a MySQL
│   ├── database.sql   # Creación inicial del esquema
│   ├── go.mod
│   └── main.go
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── components/
│   │   ├── context/
│   │   ├── pages/
│   │   ├── services/
│   │   └── styles/
│   ├── package.json
│   └── vite.config.js
├── docs/              # Requisitos, modelo de datos y decisiones de diseño
└── README.md
```

El backend sigue una arquitectura por capas:

```text
Controller -> Service -> DAO -> MySQL
```

## Requisitos previos

Antes de comenzar, instalar:

- Git.
- Go 1.26.1 o una versión compatible con la indicada en `backend/go.mod`.
- Node.js compatible con Vite 8 y npm.
- MySQL Server 8 o compatible.

Comprobar las instalaciones con:

```bash
git --version
go version
node --version
npm --version
mysql --version
```

## Instalación desde cero

Clonar el repositorio y entrar al directorio:

```bash
git clone <URL_DEL_REPOSITORIO>
cd Proyecto-desarrollo-sw
```

## Configuración de la base de datos

1. Iniciar MySQL.
2. Abrir MySQL Workbench o el cliente de línea de comandos.
3. Ejecutar el script [`backend/database.sql`](backend/database.sql). El script crea la base `proyecto_desarrollo_sw` y las tablas `usuarios`, `eventos`, `puntuaciones` y `entradas`.
4. Configurar en `backend/db/mysql.go` el DSN con las credenciales locales de MySQL.

El formato utilizado por el driver es:

```text
usuario:contraseña@tcp(localhost:3306)/proyecto_desarrollo_sw
```

No se deben subir credenciales reales al repositorio. En una instalación local, reemplazar usuario y contraseña por los correspondientes al entorno de desarrollo.

También puede cargarse el esquema desde una terminal:

```bash
mysql -u root -p < backend/database.sql
```

## Ejecución del backend

Desde la raíz del repositorio:

```bash
cd backend
go run .
```

La API queda disponible en `http://localhost:8080`. El backend debe poder conectarse a MySQL durante el inicio.

## Ejecución del frontend

En otra terminal, desde la raíz del repositorio:

```bash
cd frontend
npm install
npm run dev
```

Vite muestra en la terminal la URL del frontend, normalmente `http://localhost:5173`. El backend permite solicitudes CORS desde esa dirección.

## Tests

Los tests unitarios cubren prioritariamente controllers y services. Los dobles de base de datos se implementan dentro de archivos `*_test.go`, por lo que la suite no necesita MySQL activo.

Para ejecutar todos los tests:

```bash
cd backend
go test ./...
```

## Cobertura

Para ejecutar la suite y mostrar la cobertura por paquete:

```bash
cd backend
go test ./... -cover
```

Para generar un perfil y consultar la cobertura detallada por función:

```bash
cd backend
go test -coverprofile=coverage ./...
go tool cover -func=coverage
```

Para abrir el reporte HTML interactivo:

```bash
go tool cover -html=coverage
```

El archivo `coverage` es un artefacto local de medición y no debe incluirse en commits.

## Consigna de testing

La materia exige implementar tests que alcancen como mínimo **40% de cobertura sobre las capas de controllers y services**.

Estado actual verificado:

- **Controllers: 84.4%**
- **Services: 96.3%**
- **Requisito mínimo: 40%**

Por lo tanto, el proyecto supera actualmente el requisito de cobertura. Los tests contemplan flujos exitosos, validaciones, recursos inexistentes y errores de persistencia para `AuthController`, `EventoController`, `EntradaController`, `PuntuacionController`, `AuthService`, `EventoService`, `EntradaService` y `PuntuacionService`.

## Documentación adicional

La carpeta [`docs/`](docs/) contiene:

- Requisitos funcionales.
- Diseño de endpoints.
- Modelo de datos y DER.
- Decisiones de diseño.
- Organización y flujo de ramas del proyecto.
