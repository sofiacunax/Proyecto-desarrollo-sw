import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  obtenerMisEntradas,
  cancelarEntrada,
  transferirEntrada
} from "../services/entradasService"
import "../styles/MisEntradas.css";

function MisEntradas() {
  const navigate = useNavigate();

  const [entradas, setEntradas] = useState([]);
  const [loading, setLoading] = useState(true);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };
  const handleCancelar = async (id) => {
    try {
      const token = localStorage.getItem("token");

      const response = await cancelarEntrada(id, token);

      alert(response.message);

      const entradasActualizadas = await obtenerMisEntradas(token);

      setEntradas(entradasActualizadas || []);
    } catch (error) {
      console.error(error);
      alert("No se pudo cancelar la entrada");
    }
  };
  const handleTransferir = async (id) => {
  const nuevoUsuarioID = prompt(
    "Ingrese el ID del usuario destino"
  )

  if (!nuevoUsuarioID) return

  try {
    const token = localStorage.getItem("token")

    const response = await transferirEntrada(
      id,
      Number(nuevoUsuarioID),
      token
    )

    alert(response.message)

    const entradasActualizadas =
      await obtenerMisEntradas(token)

    setEntradas(entradasActualizadas || [])
  } catch (error) {
    console.error(error)
    alert("No se pudo transferir la entrada")
  }
}

  useEffect(() => {
    async function cargarEntradas() {
      try {
        const token = localStorage.getItem("token");

        const data = await obtenerMisEntradas(token);

        console.log(data);

        setEntradas(data || []);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    }

    cargarEntradas();
  }, []);

  return (
    <div className="tickets-page">
      <nav className="eventia-navbar">
        <div className="brand">
          <div>
            <h1>Eventia</h1>
            <p>Viví lo que te mueve</p>
          </div>
        </div>

        <div className="nav-links">
          <span onClick={() => navigate("/dashboard")}>Eventos</span>

          <span className="active">Mis Entradas</span>

          <span>Sobre Nosotros</span>
        </div>

        <button className="logout-btn" onClick={handleLogout}>
          Cerrar sesión
        </button>
      </nav>

      <main className="tickets-container">
        <section className="tickets-hero">
          <div className="hero-content">
            <h1>🎟 Mis Entradas</h1>

            <p>
              Gestioná tus entradas, transferilas a otros usuarios o cancelalas
              cuando lo necesites.
            </p>
          </div>
        </section>

        {loading ? (
          <p>Cargando entradas...</p>
        ) : (
          <section className="tickets-list">
            {entradas.length === 0 ? (
              <div className="ticket-card">
                <h3>No tenés entradas compradas</h3>
              </div>
            ) : (
              entradas.map((entrada) => (
                <div key={entrada.id} className="ticket-card">
                  <h3>{entrada.titulo}</h3>

                  <p>📅 {entrada.fecha}</p>

                  <p>📍 {entrada.ubicacion}</p>

                  <span
                    className={
                      entrada.estado === "ACTIVA"
                        ? "ticket-status active"
                        : "ticket-status cancelled"
                    }
                  >
                    {entrada.estado}
                  </span>

                  <div className="ticket-actions">
  {entrada.estado === "ACTIVA" && (
    <>
      <button
        onClick={() => handleTransferir(entrada.id)}
      >
        Transferir
      </button>

      <button
        className="cancel-btn"
        onClick={() => handleCancelar(entrada.id)}
      >
        Cancelar
      </button>
    </>
  )}
</div>
                </div>
              ))
            )}
          </section>
        )}
      </main>
    </div>
  );
}

export default MisEntradas;
