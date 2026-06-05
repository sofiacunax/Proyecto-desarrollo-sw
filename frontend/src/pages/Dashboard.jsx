import { Link } from "react-router-dom"
export default function Dashboard() {
  return (
    <div style={{ padding: "40px" }}>
      <h1>Bienvenida a Eventia 🎉</h1>

      <p>
        Esta será la página principal después del login.
      </p>

      <Link to="/">
  <button>
    Cerrar sesión
  </button>
</Link>
    </div>
  )
}