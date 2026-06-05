import { Link } from "react-router-dom"
import "../styles/Dashboard.css"

export default function Dashboard() {
  return (
    <div className="dashboard-page">

      <div className="dashboard-navbar">
        <h1 className="dashboard-logo">Eventia</h1>

        <Link to="/" className="logout-btn">
          Cerrar sesión
        </Link>
      </div>

      <div className="dashboard-content">

        <div className="dashboard-title">
          <h1>¡Hola! 👋</h1>
          <p>Bienvenida nuevamente a Eventia.</p>
        </div>

        <div className="cards">

          <div className="card">
            <h2>12</h2>
            <p>Eventos disponibles</p>
          </div>

          <div className="card">
            <h2>3</h2>
            <p>Mis entradas</p>
          </div>

          <div className="card">
            <h2>1</h2>
            <p>Mi perfil</p>
          </div>

        </div>

        <div className="events">

          <h2>Próximos eventos</h2>

          <div className="event-item">
            🎵 Rock Fest 2026
          </div>

          <div className="event-item">
            🎭 Festival Cultural
          </div>

          <div className="event-item">
            💻 Tech Summit
          </div>

        </div>

      </div>

    </div>
  )
}