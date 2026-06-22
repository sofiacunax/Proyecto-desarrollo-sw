import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { crearEvento } from "../services/eventosService"

export default function CrearEvento() {

  const [titulo, setTitulo] = useState("")
  const [descripcion, setDescripcion] = useState("")
  const [fecha, setFecha] = useState("")
  const [horario, setHorario] = useState("")
  const [duracion, setDuracion] = useState("")
  const [ubicacion, setUbicacion] = useState("")
  const [capacidad, setCapacidad] = useState("")
  const [precio, setPrecio] = useState("")
  const [categoria, setCategoria] = useState("")
  const [imagenURL, setImagenURL] = useState("")
  const navigate = useNavigate()
  const handleSubmit = async (e) => {
  e.preventDefault()

  try {

    const token =
      localStorage.getItem("token")

    await crearEvento(
      {
        titulo: titulo,
        descripcion: descripcion,
        fecha: fecha,
        horario: horario,
        duracion: Number(duracion),
        ubicacion: ubicacion,
        capacidad: Number(capacidad),
        precio: Number(precio),
        categoria: categoria,
        imagen_url: imagenURL,
        estado: "ACTIVO",
      },
      token
    )

    alert(
      "Evento creado correctamente"
    )

    navigate("/admin/eventos")

  } catch (error) {

    console.error(error)

    alert(error.message)
  }
}

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Crear Evento</h1>

      <form onSubmit={handleSubmit}>

        <input
          placeholder="Título"
          value={titulo}
          onChange={(e) => setTitulo(e.target.value)}
        />

        <br /><br />

        <input
          placeholder="Descripción"
          value={descripcion}
          onChange={(e) => setDescripcion(e.target.value)}
        />

        <br /><br />

        <input
          type="date"
          value={fecha}
          onChange={(e) => setFecha(e.target.value)}
        />

        <br /><br />

        <input
          type="time"
          value={horario}
          onChange={(e) => setHorario(e.target.value)}
        />

        <br /><br />

        <input
          type="number"
          placeholder="Duración"
          value={duracion}
          onChange={(e) => setDuracion(e.target.value)}
        />

        <br /><br />

        <input
          placeholder="Ubicación"
          value={ubicacion}
          onChange={(e) => setUbicacion(e.target.value)}
        />

        <br /><br />

        <input
          type="number"
          placeholder="Capacidad"
          value={capacidad}
          onChange={(e) => setCapacidad(e.target.value)}
        />

        <br /><br />

        <input
          type="number"
          placeholder="Precio"
          value={precio}
          onChange={(e) => setPrecio(e.target.value)}
        />

        <br /><br />

        <input
          placeholder="Categoría"
          value={categoria}
          onChange={(e) => setCategoria(e.target.value)}
        />

        <br /><br />

        <input
          placeholder="URL Imagen"
          value={imagenURL}
          onChange={(e) => setImagenURL(e.target.value)}
        />

        <br /><br />

        <button type="submit">
          Crear Evento
        </button>

      </form>
    </div>
  )
}