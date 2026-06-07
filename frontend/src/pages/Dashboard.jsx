import { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";
import { getEventos, getRankingEventos } from "../services/eventosService";
import "../styles/Dashboard.css";

export default function Dashboard() {
  const navigate = useNavigate();

  const [eventos, setEventos] = useState([]);
  const [ranking, setRanking] = useState([]);
  const [busqueda, setBusqueda] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    async function cargarDatos() {
      try {
        const eventosData = await getEventos();
        const rankingData = await getRankingEventos();

        setEventos(eventosData || []);
        setRanking(rankingData || []);
      } catch {
        setError("No se pudieron cargar los eventos");
      } finally {
        setLoading(false);
      }
    }

    cargarDatos();
  }, []);

  const eventosFiltrados = useMemo(() => {
    return eventos.filter((evento) => {
      const texto = `
        ${evento.titulo || evento.Titulo || evento.nombre || evento.Nombre || ""}
        ${evento.ubicacion || evento.Ubicacion || ""}
        ${evento.categoria || evento.Categoria || ""}
      `.toLowerCase();

      return texto.includes(busqueda.toLowerCase());
    });
  }, [eventos, busqueda]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  return (
    <div className="eventia-page">
      <nav className="eventia-navbar">
        <div className="brand">
          <div>
            <h1>Eventia</h1>
            <p>Viví lo que te mueve</p>
          </div>
        </div>

        <div className="nav-links">
          <span className="active">Eventos</span>

          <span onClick={() => navigate("/mis-entradas")}>Mis Entradas</span>

          <span>Sobre Nosotros</span>
        </div>

        <button className="logout-btn" onClick={handleLogout}>
          Cerrar sesión
        </button>
      </nav>

      <main className="eventia-container">
        <section className="hero">
          <div className="hero-text">
            <h3>¡Hola! </h3>
            <h2>
              Descubrí eventos <br />
              que <span>te mueven</span>
            </h2>
            <p>
              Explorá experiencias únicas, puntuá tus favoritos y encontrá los
              más populares.
            </p>
          </div>
        </section>

        <section className="filter-panel">
          <div className="main-search">
            <span>🔎</span>
            <input
              placeholder="Buscar eventos por nombre, ubicación o categoría..."
              value={busqueda}
              onChange={(e) => setBusqueda(e.target.value)}
            />
          </div>

          <p className="status-text">
            Mostrando <b>{eventosFiltrados.length}</b> eventos
          </p>
        </section>

        {loading && <p className="loading">Cargando eventos...</p>}
        {error && <p className="error-message">{error}</p>}

        {!loading && !error && (
          <>
            <section className="popular-header">
              <div>
                <div className="section-icon"></div>
                <div>
                  <h2>Eventos más populares</h2>
                  <p>
                    Top 3 eventos mas destacados segun valoraciones de nuestra
                    comunidad
                  </p>
                </div>
              </div>
            </section>

            <section className="events-grid">
              {ranking.length === 0 ? (
                <p className="empty-message">
                  Todavía no hay eventos puntuados.
                </p>
              ) : (
                ranking
                  .slice(0, 3)
                  .map((evento, index) => (
                    <RankingCard
                      key={evento.id || evento.ID || evento.evento_id || index}
                      evento={evento}
                      rank={index + 1}
                    />
                  ))
              )}
            </section>

            <section className="popular-header">
              <div>
                <div className="section-icon">📅</div>
                <div>
                  <h2>Todos los eventos</h2>
                </div>
              </div>
            </section>

            <section className="events-grid">
              {eventosFiltrados.length === 0 ? (
                <p className="empty-message">No se encontraron eventos.</p>
              ) : (
                eventosFiltrados.map((evento) => (
                  <EventCard key={evento.id || evento.ID} evento={evento} />
                ))
              )}
            </section>
          </>
        )}
      </main>
    </div>
  );
}

function EventCard({ evento }) {
  const titulo =
    evento.titulo ||
    evento.Titulo ||
    evento.nombre ||
    evento.Nombre ||
    "Evento";
  const descripcion = evento.descripcion || evento.Descripcion || "";
  const fecha = evento.fecha || evento.Fecha || "";
  const horario = evento.horario || evento.Horario || "";
  const ubicacion = evento.ubicacion || evento.Ubicacion || "";
  const categoria = evento.categoria || evento.Categoria || "Evento";
  const precio = evento.precio || evento.Precio || 0;
  const imagen =
    evento.imagen_url || evento.ImagenURL || evento.imagen || evento.Imagen;
  const [mostrarPuntuar, setMostrarPuntuar] = useState(false);
  const [puntuacion, setPuntuacion] = useState(0);
  
  const enviarPuntuacion = async (valor) => {
    setPuntuacion(valor);

    const token = localStorage.getItem("token");
    const eventoID = evento.id || evento.ID;

    console.log("Mandando puntuación");
    console.log("Token:", token);
    console.log("Evento ID:", eventoID);
    console.log("Valor:", valor);

    try {
      const response = await fetch(
        "http://localhost:8080/private/puntuaciones",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify({
            evento_id: eventoID,
            puntuacion: valor,
          }),
        },
      );

      const data = await response.json();

      console.log("Status:", response.status);
      console.log("Respuesta:", data);

      if (!response.ok) {
        alert(data.error || "No se pudo guardar la puntuación");
        return;
      }

      alert("Puntuación guardada");
    } catch (error) {
      console.error(error);
      alert(error.message);
    }
  };

  return (
    <article
     className="event-card"
      
      >
      <img
        src={
          imagen ||
          "https://images.unsplash.com/photo-1501281668745-f7f57925c3b4?q=80&w=1200&auto=format&fit=crop"
        }
        alt={titulo}
      />

      <div className="event-info">
        <button
          className="event-state score-btn"
          onClick={() => setMostrarPuntuar(!mostrarPuntuar)}
        >
          ★ Puntuar
        </button>

        {mostrarPuntuar && (
          <div className="rating-bar">
            {[1, 2, 3, 4, 5].map((valor) => (
              <button
                key={valor}
                onClick={() => enviarPuntuacion(valor)}
                className={puntuacion >= valor ? "star active-star" : "star"}
              >
                ★
              </button>
            ))}
          </div>
        )}

        <h3>{titulo}</h3>

        <p className="event-description">{descripcion}</p>

        <div className="meta">
          <span>{fecha}</span>
          <span>{horario}</span>
          <span>{ubicacion}</span>
        </div>

        <div className="tag-row">
          <span>{categoria}</span>
          <b>${Number(precio).toLocaleString("es-AR")}</b>
        </div>

        <div className="card-bottom">
  <button
  onClick={() =>
    window.location.href = `/comprar/${evento.id || evento.ID}`
  }
>
  Comprar
</button>
</div>
      </div>
      
    </article>
  );
}

function RankingCard({ evento, rank }) {
  const navigate = useNavigate()
  const titulo =
    evento.titulo ||
    evento.Titulo ||
    evento.nombre ||
    evento.Nombre ||
    "Evento";
  const promedio =
    evento.promedio || evento.Promedio || evento.puntuacion_promedio || 0;
  const votos =
    evento.votos || evento.Votos || evento.cantidad_puntuaciones || "";

  return (
    <article className="event-card ranking-card">
      <div className="rank-badge">{rank}</div>

      <div className="event-info">
        <h3>{titulo}</h3>

        <div className="rating">
          <strong>⭐ {Number(promedio).toFixed(1)}</strong>
          {votos && <small>({votos} votos)</small>}
          <p>Puntuación promedio</p>
        </div>

        <div className="card-bottom">
          <button
  onClick={() =>
    navigate(
      `/comprar/${evento.id || evento.ID || evento.evento_id}`
    )
  }
>
  Comprar
</button>
        </div>
      </div>
    </article>
  );
}

