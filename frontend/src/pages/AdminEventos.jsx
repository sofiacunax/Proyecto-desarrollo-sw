import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { FiArrowLeft, FiBarChart2, FiCalendar, FiEdit2, FiPlus, FiTag, FiTrash2, FiUsers } from "react-icons/fi";
import { getEventos, eliminarEvento } from "../services/eventosService";
import "../styles/AdminEventos.css";

export default function AdminEventos() {
  const navigate = useNavigate();
  const [eventos, setEventos] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function cargarEventos() {
      try {
        const data = await getEventos("");
        setEventos(data || []);
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    }
    cargarEventos();
  }, []);

  const handleEliminar = async (id) => {
    const confirmar = window.confirm("¿Seguro que desea eliminar este evento?");
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
            <h1>Administración de Eventos</h1>
            <p>Gestioná todos los eventos publicados en Eventia</p>
          </div>
          <span className="events-count">{eventos.length} eventos</span>
        </header>

        <nav className="admin-actions" aria-label="Acciones de administración">
          <button className="admin-btn admin-btn-secondary" onClick={() => navigate("/dashboard")}>
            <FiArrowLeft aria-hidden="true" /> Volver al Dashboard
          </button>
          <button className="admin-btn admin-btn-primary" onClick={() => navigate("/admin/eventos/nuevo")}>
            <FiPlus aria-hidden="true" /> Nuevo Evento
          </button>
          <button className="admin-btn admin-btn-secondary" onClick={() => navigate("/admin/usuarios")}>
            <FiUsers aria-hidden="true" /> Administrar Usuarios
          </button>
        </nav>

        {loading ? (
          <div className="admin-feedback" role="status"><span className="admin-spinner" />Cargando eventos...</div>
        ) : eventos.length === 0 ? (
          <div className="admin-feedback admin-empty">
            <FiCalendar aria-hidden="true" />
            <h2>Todavía no hay eventos</h2>
            <p>Creá el primero para comenzar a gestionarlo.</p>
          </div>
        ) : (
          <section className="admin-events-grid" aria-label="Eventos publicados">
            {eventos.map((evento) => {
              const id = evento.id || evento.ID;
              const estado = evento.estado || evento.Estado || "Sin estado";
              return (
                <article className="admin-event-card" key={id}>
                  <div className="admin-event-card-top">
                    <span className={`event-status event-status-${estado.toLowerCase()}`}>{estado}</span>
                    <span className="event-id">#{id}</span>
                  </div>
                  <h2>{evento.titulo || evento.Titulo}</h2>
                  <div className="admin-event-meta">
                    <span>
  <FiCalendar aria-hidden="true" />
  {new Date(
    evento.fecha || evento.Fecha
  ).toLocaleDateString("es-AR")}
</span>
                    <span><FiTag aria-hidden="true" />{evento.categoria || evento.Categoria}</span>
                  </div>
                  <div className="admin-event-card-actions">
                    <button className="card-action edit-action" onClick={() => navigate(`/admin/eventos/editar/${id}`)}>
                      <FiEdit2 aria-hidden="true" /> Editar
                    </button>
                    <button className="card-action report-action" onClick={() => navigate(`/admin/eventos/${id}/reporte`)}>
                      <FiBarChart2 aria-hidden="true" /> Reporte
                    </button>
                    <button className="card-action delete-action" onClick={() => handleEliminar(id)}>
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
