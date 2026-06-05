import "./../styles/Login.css"

export default function Login() {
  return (
    <div className="login-page">

      <div className="login-container">

        <div className="login-banner">
          <div className="banner-overlay">
            <h1>Descubrí experiencias inolvidables</h1>
            <p>
  Encontrá conciertos, festivales, teatro y eventos cerca tuyo.
</p>
          </div>
        </div>

        <div className="login-card">

          <div className="login-logo">

  <h1>Eventia</h1>

  <span className="logo-tagline">
    VIVÍ LO QUE TE MUEVE
  </span>

  <p className="login-subtitle">
    Iniciá sesión para continuar
  </p>

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

            <a
  href="/dashboard"
  className="login-button"
>
  Iniciar sesión
</a>

          </form>

          <div className="login-links">
            <a href="#">
              ¿Olvidaste tu contraseña?
            </a>

            <p>
              ¿No tenés cuenta?
              <a href="/register">
                Registrate
              </a>
            </p>
          </div>

        </div>

      </div>

    </div>
  )
}