import "./../styles/Login.css"

export default function Login() {
  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          <h1>Eventia</h1>
          <p>Iniciá sesión para continuar</p>
        </div>

        <form>

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

          <button className="login-button">
            Iniciar sesión
          </button>

        </form>

        <div className="login-links">
          <a href="#">¿Olvidaste tu contraseña?</a>

          <p>
            ¿No tenés cuenta?{" "}
            <a href="/register">
              Registrate
            </a>
          </p>
        </div>

      </div>
    </div>
  )
}