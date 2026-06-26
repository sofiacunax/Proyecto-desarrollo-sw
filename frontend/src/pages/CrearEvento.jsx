import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { FiArrowLeft, FiCalendar, FiDollarSign, FiImage, FiInfo, FiMapPin, FiUsers } from "react-icons/fi";
import { crearEvento } from "../services/eventosService";
import "../styles/CrearEvento.css";

export default function CrearEvento() {
  const [titulo, setTitulo] = useState("");
  const [descripcion, setDescripcion] = useState("");
  const [fecha, setFecha] = useState("");
  const [horario, setHorario] = useState("");
  const [duracion, setDuracion] = useState("");
  const [ubicacion, setUbicacion] = useState("");
  const [capacidad, setCapacidad] = useState("");
  const [precio, setPrecio] = useState("");
  const [categoria, setCategoria] = useState("");
  const [imagenURL, setImagenURL] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem("token");
      await crearEvento({ titulo, descripcion, fecha, horario, duracion: Number(duracion), ubicacion, capacidad: Number(capacidad), precio: Number(precio), categoria, imagen_url: imagenURL, estado: "ACTIVO" }, token);
      alert("Evento creado correctamente");
      navigate("/admin/eventos");
    } catch (error) {
      console.error(error);
      alert(error.message);
    }
  };

  return (
    <main className="event-form-page create-event-page">
      <div className="event-form-shell">
        <button className="form-back-button" type="button" onClick={() => navigate("/admin/eventos")}>
          <FiArrowLeft aria-hidden="true" /> Volver a eventos
        </button>
        <section className="event-form-card">
          <header className="event-form-header">
            <span className="form-eyebrow">Nuevo evento</span>
            <h1>Crear Nuevo Evento</h1>
            <p>Completá la información para publicar un evento</p>
          </header>
          <form onSubmit={handleSubmit}>
            <fieldset className="form-section">
              <legend><FiInfo aria-hidden="true" /><span>Información general</span></legend>
              <div className="form-field"><label htmlFor="titulo">Título</label><input id="titulo" placeholder="Ej. Festival de música" value={titulo} onChange={(e) => setTitulo(e.target.value)} /></div>
              <div className="form-field"><label htmlFor="descripcion">Descripción</label><textarea id="descripcion" rows="4" placeholder="Contá de qué se trata el evento" value={descripcion} onChange={(e) => setDescripcion(e.target.value)} /></div>
              <div className="form-field"><label htmlFor="categoria">Categoría</label><input id="categoria" placeholder="Ej. Música, teatro, deportes" value={categoria} onChange={(e) => setCategoria(e.target.value)} /></div>
            </fieldset>
            <fieldset className="form-section">
              <legend><FiCalendar aria-hidden="true" /><span>Fecha y horario</span></legend>
              <div className="form-row form-row-three">
                <div className="form-field"><label htmlFor="fecha">Fecha</label><input id="fecha" type="date" value={fecha} onChange={(e) => setFecha(e.target.value)} /></div>
                <div className="form-field"><label htmlFor="horario">Horario</label><input id="horario" type="time" value={horario} onChange={(e) => setHorario(e.target.value)} /></div>
                <div className="form-field"><label htmlFor="duracion">Duración</label><input id="duracion" type="number" placeholder="Minutos" value={duracion} onChange={(e) => setDuracion(e.target.value)} /></div>
              </div>
            </fieldset>
            <fieldset className="form-section">
              <legend><FiMapPin aria-hidden="true" /><span>Ubicación</span></legend>
              <div className="form-field"><label htmlFor="ubicacion">Dirección o lugar</label><input id="ubicacion" placeholder="Ej. Teatro Gran Rex, Buenos Aires" value={ubicacion} onChange={(e) => setUbicacion(e.target.value)} /></div>
            </fieldset>
            <fieldset className="form-section">
              <legend><FiUsers aria-hidden="true" /><span>Capacidad y precio</span></legend>
              <div className="form-row">
                <div className="form-field"><label htmlFor="capacidad">Capacidad</label><input id="capacidad" type="number" placeholder="Cantidad de personas" value={capacidad} onChange={(e) => setCapacidad(e.target.value)} /></div>
                <div className="form-field"><label htmlFor="precio"><FiDollarSign aria-hidden="true" /> Precio</label><input id="precio" type="number" placeholder="0,00" value={precio} onChange={(e) => setPrecio(e.target.value)} /></div>
              </div>
            </fieldset>
            <fieldset className="form-section">
              <legend><FiImage aria-hidden="true" /><span>Imagen</span></legend>
              <div className="form-field"><label htmlFor="imagenURL">URL de la imagen</label><input id="imagenURL" placeholder="https://ejemplo.com/imagen.jpg" value={imagenURL} onChange={(e) => setImagenURL(e.target.value)} /></div>
            </fieldset>
            <button className="event-submit-button" type="submit">Publicar Evento</button>
          </form>
        </section>
      </div>
    </main>
  );
}
