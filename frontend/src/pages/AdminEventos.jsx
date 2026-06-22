import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  getEventos,
  eliminarEvento,
} from "../services/eventosService";

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
  const confirmar = window.confirm(
    "¿Seguro que desea eliminar este evento?"
  )

  if (!confirmar) {
    return
  }

  try {
    const token = localStorage.getItem("token")

    await eliminarEvento(id, token)

    setEventos(
      eventos.filter(
        (evento) =>
          (evento.id || evento.ID) !== id
      )
    )

    alert("Evento eliminado correctamente")
  } catch (error) {
    console.error(error)

    alert(error.message)
  }
}

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Administración de Eventos</h1>

      <button
        onClick={() => navigate("/dashboard")}
        style={{ marginBottom: "20px" }}
      >
        Volver
      </button>
      <button
  onClick={() =>
    navigate("/admin/eventos/nuevo")
  }
>
  + Nuevo Evento
</button>
<button
  onClick={() => navigate("/admin/usuarios")}
>
  Administrar Usuarios
</button>

      {loading ? (
        <p>Cargando eventos...</p>
      ) : (
        
        <table border="1" cellPadding="10">
          <thead>
            <tr>
              <th>ID</th>
              <th>Título</th>
              <th>Fecha</th>
              <th>Categoría</th>
              <th>Estado</th>
              <th>Acciones</th>
            </tr>
          </thead>

          <tbody>
            {eventos.map((evento) => (
              <tr key={evento.id || evento.ID}>
                <td>{evento.id || evento.ID}</td>

                <td>
                  {evento.titulo || evento.Titulo}
                </td>

                <td>
                  {evento.fecha || evento.Fecha}
                </td>

                <td>
                  {evento.categoria || evento.Categoria}
                </td>

                <td>
                  {evento.estado || evento.Estado}
                </td>
                <td>

<button
  onClick={() =>
    navigate(
  `/admin/eventos/editar/${
    evento.id || evento.ID
  }`
)
  }
>
  Editar
</button>
  <button
  style={{
    marginLeft: "10px",
  }}
  onClick={() =>
    handleEliminar(
      evento.id || evento.ID
    )
  }
>
  Eliminar
</button>
</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}