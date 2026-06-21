import "./../styles/Login.css"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { login } from "../services/authService"

export default function Login() {
  const [email, setEmail] = useState("")
const [password, setPassword] = useState("")

const navigate = useNavigate()
const handleLogin = async (e) => {
  e.preventDefault()

  try {
    const data = await login(email, password)

    localStorage.setItem("token", data.token)

    navigate("/dashboard")
  } catch (error) {
    alert(error.message)
  }
}
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

          <form onSubmit={handleLogin}>

            <div className="form-group">
              <label>Email</label>
              <input
  type="email"
  placeholder="ejemplo@email.com"
  value={email}
  onChange={(e) => setEmail(e.target.value)}
/>
            </div>

            <div className="form-group">
              <label>Contraseña</label>
              <input
  type="password"
  placeholder="********"
  value={password}
  onChange={(e) => setPassword(e.target.value)}
/>
            </div>

            <button
  type="submit"
  className="login-button"
>
  Iniciar sesión
</button>

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