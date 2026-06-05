import "./../styles/Login.css"


export default function Register() {
  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          <h1>Eventia</h1>
          <p>Creá tu cuenta para empezar</p>
        </div>

        <form>

          <div className="form-group">
            <label>Nombre completo</label>
            <input
              type="text"
              placeholder="Juan Pérez"
            />
          </div>

          <div className="form-group">
            <label>Email</label>
            <input
              type="email"
              placeholder="ejemplo@email.com"
            />
          </div>

          <div className="form-group">
            <label>Contraseña</label>
            <input
              type="password"
              placeholder="********"
            />
          </div>

          <div className="form-group">
            <label>Confirmar contraseña</label>
            <input
              type="password"
              placeholder="********"
            />
          </div>

          <button className="login-button">
            Crear cuenta
          </button>

        </form>

        <div className="login-links">
          <p>
            ¿Ya tenés cuenta?
            <a href="/">
              Iniciá sesión
            </a>
          </p>
        </div>

      </div>
    </div>
  )
}