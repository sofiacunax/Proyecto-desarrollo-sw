import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"

import {
  getEventoPorId,
  actualizarEvento,
} from "../services/eventosService"

export default function EditarEvento() {

  const { id } = useParams()
  const navigate = useNavigate()

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

  useEffect(() => {
    async function cargarEvento() {

      try {

        const evento =
          await getEventoPorId(id)

        setTitulo(
          evento.titulo || evento.Titulo || ""
        )

        setDescripcion(
          evento.descripcion ||
          evento.Descripcion ||
          ""
        )

        setFecha(
          evento.fecha || evento.Fecha || ""
        )

        setHorario(
          evento.horario ||
          evento.Horario ||
          ""
        )

        setDuracion(
          evento.duracion ||
          evento.Duracion ||
          ""
        )

        setUbicacion(
          evento.ubicacion ||
          evento.Ubicacion ||
          ""
        )

        setCapacidad(
          evento.capacidad ||
          evento.Capacidad ||
          ""
        )

        setPrecio(
          evento.precio ||
          evento.Precio ||
          ""
        )

        setCategoria(
          evento.categoria ||
          evento.Categoria ||
          ""
        )

        setImagenURL(
          evento.imagen_url ||
          evento.ImagenURL ||
          ""
        )

      } catch (error) {
        console.error(error)
      }
    }

    cargarEvento()

  }, [id])

  const handleSubmit = async (e) => {

    e.preventDefault()

    try {

      const token =
        localStorage.getItem("token")

      await actualizarEvento(
        id,
        {
          titulo,
          descripcion,
          fecha,
          horario,
          duracion: Number(duracion),
          ubicacion,
          capacidad: Number(capacidad),
          precio: Number(precio),
          categoria,
          imagen_url: imagenURL,
          estado: "ACTIVO",
        },
        token
      )

      alert(
        "Evento actualizado correctamente"
      )

      navigate("/admin/eventos")

    } catch (error) {

      console.error(error)

      alert(error.message)
    }
  }

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Editar Evento</h1>

      <form onSubmit={handleSubmit}>

        <input
          value={titulo}
          onChange={(e) =>
            setTitulo(e.target.value)
          }
          placeholder="Título"
        />

        <br /><br />

        <input
          value={descripcion}
          onChange={(e) =>
            setDescripcion(e.target.value)
          }
          placeholder="Descripción"
        />

        <br /><br />

        <input
          type="date"
          value={fecha}
          onChange={(e) =>
            setFecha(e.target.value)
          }
        />

        <br /><br />

        <input
          type="time"
          value={horario}
          onChange={(e) =>
            setHorario(e.target.value)
          }
        />

        <br /><br />

        <input
          value={duracion}
          onChange={(e) =>
            setDuracion(e.target.value)
          }
          placeholder="Duración"
        />

        <br /><br />

        <input
          value={ubicacion}
          onChange={(e) =>
            setUbicacion(e.target.value)
          }
          placeholder="Ubicación"
        />

        <br /><br />

        <input
          value={capacidad}
          onChange={(e) =>
            setCapacidad(e.target.value)
          }
          placeholder="Capacidad"
        />

        <br /><br />

        <input
          value={precio}
          onChange={(e) =>
            setPrecio(e.target.value)
          }
          placeholder="Precio"
        />

        <br /><br />

        <input
          value={categoria}
          onChange={(e) =>
            setCategoria(e.target.value)
          }
          placeholder="Categoría"
        />

        <br /><br />

        <input
          value={imagenURL}
          onChange={(e) =>
            setImagenURL(e.target.value)
          }
          placeholder="URL Imagen"
        />

        <br /><br />

        <button type="submit">
          Guardar cambios
        </button>

      </form>
    </div>
  )
}