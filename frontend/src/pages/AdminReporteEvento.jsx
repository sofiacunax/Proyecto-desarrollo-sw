import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  FiArrowLeft,
  FiBarChart2,
  FiHash,
  FiMail,
  FiPieChart,
  FiTag,
  FiUser,
  FiUsers,
} from "react-icons/fi";
import { getReporteEvento } from "../services/eventosService";
import "../styles/AdminReporteEvento.css";

export default function AdminReporteEvento() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [reporte, setReporte] = useState(null);

  useEffect(() => {
    async function cargarReporte() {
      try {
        const token = localStorage.getItem("token");
        const data = await getReporteEvento(id, token);
        setReporte(data);
      } catch (error) {
        console.error(error);
      }
    }

    cargarReporte();
  }, [id]);

  if (!reporte) {
    return (
      <main className="report-page">
        <div className="report-container">
          <div className="report-loading" role="status">
            <span className="report-spinner" aria-hidden="true" />
            <span>Cargando reporte...</span>
          </div>
        </div>
      </main>
    );
  }

  return (
    <main className="report-page">
      <div className="report-container">
        <button
          className="report-back-button"
          type="button"
          onClick={() => navigate("/admin/eventos")}
        >
          <FiArrowLeft aria-hidden="true" />
          Volver a Eventos
        </button>

        <header className="report-header">
          <div>
            <span className="report-eyebrow">Panel administrativo</span>
            <h1>Reporte de Evento</h1>
            <p>Análisis de ocupación y compradores</p>
          </div>
          <div className="report-header-icon" aria-hidden="true">
            <FiBarChart2 />
          </div>
        </header>

        <section className="event-summary-card" aria-labelledby="event-summary-title">
          <div className="event-summary-icon" aria-hidden="true">
            <FiTag />
          </div>
          <div className="event-summary-main">
            <span>Evento analizado</span>
            <h2 id="event-summary-title">{reporte.titulo}</h2>
          </div>
          <div className="event-summary-details">
            <span><FiHash aria-hidden="true" /> ID {reporte.evento_id}</span>
            <span><FiUsers aria-hidden="true" /> Capacidad {reporte.capacidad}</span>
          </div>
        </section>

        <section className="report-section" aria-labelledby="metrics-title">
          <div className="report-section-heading">
            <div>
              <span className="report-section-kicker">Resumen</span>
              <h2 id="metrics-title">Métricas principales</h2>
            </div>
          </div>

          <div className="metrics-grid">
            <article className="metric-card">
              <div className="metric-icon capacity-icon" aria-hidden="true"><FiUsers /></div>
              <span>Capacidad Total</span>
              <strong>{reporte.capacidad}</strong>
              <small>lugares disponibles</small>
            </article>

            <article className="metric-card">
              <div className="metric-icon sales-icon" aria-hidden="true"><FiBarChart2 /></div>
              <span>Entradas Vendidas</span>
              <strong>{reporte.entradas_vendidas}</strong>
              <small>entradas registradas</small>
            </article>

            <article className="metric-card occupancy-card">
              <div className="metric-icon occupancy-icon" aria-hidden="true"><FiPieChart /></div>
              <span>Ocupación</span>
              <strong>{reporte.porcentaje_ocupado.toFixed(2)}%</strong>
              <div className="occupancy-track" aria-hidden="true">
                <span style={{ width: `${Math.min(reporte.porcentaje_ocupado, 100)}%` }} />
              </div>
            </article>
          </div>
        </section>

        <section className="report-section buyers-section" aria-labelledby="buyers-title">
          <div className="report-section-heading buyers-heading">
            <div>
              <span className="report-section-kicker">Asistentes</span>
              <h2 id="buyers-title">Compradores</h2>
            </div>
            <span className="buyers-count">{reporte.compradores?.length || 0} registrados</span>
          </div>

          {reporte.compradores?.length === 0 ? (
            <div className="buyers-empty">
              <div className="buyers-empty-icon" aria-hidden="true"><FiUsers /></div>
              <h3>Sin compradores por el momento</h3>
              <p>No existen compradores registrados para este evento.</p>
            </div>
          ) : (
            <div className="buyers-grid">
              {reporte.compradores?.map((comprador) => (
                <article className="buyer-card" key={comprador.id}>
                  <div className="buyer-avatar" aria-hidden="true"><FiUser /></div>
                  <div className="buyer-info">
                    <h3>{comprador.nombre}</h3>
                    <a href={`mailto:${comprador.email}`}>
                      <FiMail aria-hidden="true" />
                      <span>{comprador.email}</span>
                    </a>
                  </div>
                </article>
              ))}
            </div>
          )}
        </section>
      </div>
    </main>
  );
}
