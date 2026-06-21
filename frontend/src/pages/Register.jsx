import "./../styles/Login.css"
import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { register } from "../services/authService"


export default function Register() {
  const [nombre, setNombre] = useState("")
const [email, setEmail] = useState("")
const [password, setPassword] = useState("")
const [confirmPassword, setConfirmPassword] = useState("")

const navigate = useNavigate()

const handleRegister = async (e) => {
  e.preventDefault()

  if (password !== confirmPassword) {
    alert("Las contraseñas no coinciden")
    return
  }

  try {
    await register(
      nombre,
      email,
      password
    )

    alert("Usuario registrado correctamente")

    navigate("/")

  } catch (error) {
    alert(error.message)
  }
}
  return (
    <div className="login-page">
      <div className="login-card">

        <div className="login-logo">
          <h1>Eventia</h1>
          <p>Creá tu cuenta para empezar</p>
        </div>

        <form onSubmit={handleRegister}>

          <div className="form-group">
            <label>Nombre completo</label>
            <input
  type="text"
  placeholder="Juan Pérez"
  value={nombre}
  onChange={(e) => setNombre(e.target.value)}
/>
          </div>

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

          <div className="form-group">
            <label>Confirmar contraseña</label>
            <input
  type="password"
  placeholder="********"
  value={confirmPassword}
  onChange={(e) => setConfirmPassword(e.target.value)}
/>
          </div>

          <button
  type="submit"
  className="login-button"
>
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