import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  FiArrowLeft,
  FiBarChart2,
  FiCalendar,
  FiEdit2,
  FiPlus,
  FiRefreshCw,
  FiSlash,
  FiTag,
  FiTrash2,
  FiUsers,
} from "react-icons/fi";
import {
  cambiarEstadoEvento,
  eliminarEvento,
  getEventosAdmin,
} from "../services/eventosService";
import "../styles/AdminEventos.css";

export default function AdminEventos() {
  const navigate = useNavigate();
  const [eventos, setEventos] = useState([]);
  const [loading, setLoading] = useState(true);
  const [estadoFiltro, setEstadoFiltro] = useState("TODOS");
  const [error, setError] = useState("");
  const [actualizandoId, setActualizandoId] = useState(null);

  useEffect(() => {
    async function cargarEventos() {
      try {
        const token = localStorage.getItem("token");
        const data = await getEventosAdmin("", token);
        setEventos(data || []);
      } catch (error) {
        console.error(error);
        setError(error.message || "No se pudieron cargar los eventos");
      } finally {
        setLoading(false);
      }
    }
    cargarEventos();
  }, []);

  const eventosFiltrados = eventos.filter((evento) => {
    if (estadoFiltro === "TODOS") return true;
    return (evento.estado || evento.Estado || "ACTIVO") === estadoFiltro;
  });

  const handleCambiarEstado = async (evento) => {
    const id = evento.id || evento.ID;
    const estadoActual = evento.estado || evento.Estado || "ACTIVO";
    const nuevoEstado = estadoActual === "CANCELADO" ? "ACTIVO" : "CANCELADO";

    try {
      setActualizandoId(id);
      const token = localStorage.getItem("token");
      await cambiarEstadoEvento(id, nuevoEstado, token);
      setEventos((items) =>
        items.map((item) =>
          (item.id || item.ID) === id
            ? { ...item, estado: nuevoEstado, Estado: nuevoEstado }
            : item
        )
      );
    } catch (error) {
      console.error(error);
      alert(error.message);
    } finally {
      setActualizandoId(null);
    }
  };

  const handleEliminar = async (id) => {
    const confirmar = window.confirm(
      "Eliminar borra el evento definitivamente solo si no tiene entradas ni puntuaciones. Para quitarlo de la venta, use Cancelar evento. Desea continuar?"
    );
    if (!confirmar) return;

    try {
      const token = localStorage.getItem("token");
      await eliminarEvento(id, token);
      setEventos(eventos.filter((evento) => (evento.id || evento.ID) !== id));
      alert("Evento eliminado correctamente");
    } catch (error) {
      console.error(error);
      alert(error.message);
    }
  };

  return (
    <main className="admin-events-page">
      <div className="admin-events-container">
        <header className="admin-events-header">
          <div>
            <span className="admin-eyebrow">Panel administrativo</span>
            <h1>Administracion de Eventos</h1>
            <p>Gestiona todos los eventos publicados en Eventia</p>
          </div>
          <span className="events-count">{eventosFiltrados.length} eventos</span>
        </header>

        <nav className="admin-actions" aria-label="Acciones de administracion">
          <button
            className="admin-btn admin-btn-secondary"
            onClick={() => navigate("/dashboard")}
          >
            <FiArrowLeft aria-hidden="true" /> Volver al Dashboard
          </button>
          <button
            className="admin-btn admin-btn-primary"
            onClick={() => navigate("/admin/eventos/nuevo")}
          >
            <FiPlus aria-hidden="true" /> Nuevo Evento
          </button>
          <button
            className="admin-btn admin-btn-secondary"
            onClick={() => navigate("/admin/usuarios")}
          >
            <FiUsers aria-hidden="true" /> Administrar Usuarios
          </button>
        </nav>

        <div className="status-filter" role="tablist" aria-label="Filtrar eventos por estado">
          {[
            ["TODOS", "Todos"],
            ["ACTIVO", "Activos"],
            ["CANCELADO", "Cancelados"],
          ].map(([value, label]) => (
            <button
              key={value}
              type="button"
              className={estadoFiltro === value ? "active" : ""}
              onClick={() => setEstadoFiltro(value)}
            >
              {label}
            </button>
          ))}
        </div>

        {error && <div className="admin-alert">{error}</div>}

        {loading ? (
          <div className="admin-feedback" role="status">
            <span className="admin-spinner" />Cargando eventos...
          </div>
        ) : eventosFiltrados.length === 0 ? (
          <div className="admin-feedback admin-empty">
            <FiCalendar aria-hidden="true" />
            <h2>Todavia no hay eventos</h2>
            <p>No hay eventos para el filtro seleccionado.</p>
          </div>
        ) : (
          <section className="admin-events-grid" aria-label="Eventos publicados">
            {eventosFiltrados.map((evento) => {
              const id = evento.id || evento.ID;
              const estado = evento.estado || evento.Estado || "ACTIVO";
              const cancelado = estado === "CANCELADO";
              const imagen = evento.imagen_url || evento.ImagenURL;

              return (
                <article
                  className={`admin-event-card ${
                    cancelado ? "admin-event-card-cancelled" : ""
                  }`}
                  key={id}
                >
                  <div className="admin-event-image">
                    <img
                      src={
                        imagen ||
                        "https://images.unsplash.com/photo-1501281668745-f7f57925c3b4?q=80&w=1200&auto=format&fit=crop"
                      }
                      alt={evento.titulo || evento.Titulo}
                    />
                  </div>
                  <div className="admin-event-card-top">
                    <span className={`event-status event-status-${estado.toLowerCase()}`}>
                      {estado}
                    </span>
                    <span className="event-id">#{id}</span>
                  </div>
                  <h2>{evento.titulo || evento.Titulo}</h2>
                  <div className="admin-event-meta">
                    <span>
                      <FiCalendar aria-hidden="true" />
                      {new Date(evento.fecha || evento.Fecha).toLocaleDateString("es-AR")}
                    </span>
                    <span>
                      <FiTag aria-hidden="true" />
                      {evento.categoria || evento.Categoria}
                    </span>
                  </div>
                  <div className="admin-event-card-actions">
                    <button
                      className={`card-action state-action ${
                        cancelado ? "reactivate-action" : "cancel-action"
                      }`}
                      onClick={() => handleCambiarEstado(evento)}
                      disabled={actualizandoId === id}
                    >
                      {cancelado ? (
                        <FiRefreshCw aria-hidden="true" />
                      ) : (
                        <FiSlash aria-hidden="true" />
                      )}
                      {cancelado ? "Reactivar" : "Cancelar"}
                    </button>
                    <button
                      className="card-action edit-action"
                      onClick={() => navigate(`/admin/eventos/editar/${id}`)}
                    >
                      <FiEdit2 aria-hidden="true" /> Editar
                    </button>
                    <button
                      className="card-action report-action"
                      onClick={() => navigate(`/admin/eventos/${id}/reporte`)}
                    >
                      <FiBarChart2 aria-hidden="true" /> Reporte
                    </button>
                    <button
                      className="card-action delete-action"
                      onClick={() => handleEliminar(id)}
                    >
                      <FiTrash2 aria-hidden="true" /> Eliminar
                    </button>
                  </div>
                </article>
              );
            })}
          </section>
        )}
      </div>
    </main>
  );
}
