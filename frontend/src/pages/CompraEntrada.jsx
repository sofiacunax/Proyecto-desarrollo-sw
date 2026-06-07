import { useParams, useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"
import { getEventoPorId } from "../services/eventosService"
import { comprarEntrada } from "../services/entradasService"
import "../styles/CompraEntrada.css"

function CompraEntrada() {
  const { id } = useParams()
  const navigate = useNavigate()

  const [evento, setEvento] = useState(null)
  const [cantidad, setCantidad] = useState(1)
  const handleComprar = async () => {
  try {
    const token = localStorage.getItem("token")

    for (let i = 0; i < cantidad; i++) {
      await comprarEntrada(
        evento.id || evento.ID,
        token
      )
    }

    alert("Compra realizada correctamente")

    navigate("/mis-entradas")
  } catch (error) {
    console.error(error)

    alert("No se pudo completar la compra")
  }
}

  useEffect(() => {
    async function cargarEvento() {
      try {
        const data = await getEventoPorId(id)
        setEvento(data)
      } catch (error) {
        console.error(error)
      }
    }

    cargarEvento()
  }, [id])

  if (!evento) {
    return <p>Cargando evento...</p>
  }

  return (
    <div className="purchase-page">
      <div className="purchase-container">

        <button
          className="back-btn"
          onClick={() => navigate("/dashboard")}
        >
          ← Volver a eventos
        </button>

        <h1>Confirmar compra</h1>

        <div className="purchase-card">

          <img
            src={
              evento.imagen_url ||
              evento.ImagenURL ||
              "https://images.unsplash.com/photo-1501281668745-f7f57925c3b4?q=80&w=1200&auto=format&fit=crop"
            }
            alt={evento.titulo || evento.Titulo}
          />

          <div className="purchase-info">

            <h2>
              {evento.titulo || evento.Titulo}
            </h2>

            <div className="category-pill">
              {evento.categoria || evento.Categoria}
            </div>

            <div className="event-description-box">
              {evento.descripcion || evento.Descripcion}
            </div>

            <p>
              📅 {evento.fecha || evento.Fecha}
            </p>

            <p>
              🕒 {evento.horario || evento.Horario}
            </p>

            <p>
              📍 {evento.ubicacion || evento.Ubicacion}
            </p>

            <p>
              💲 Precio: $
              {Number(
                evento.precio || evento.Precio || 0
              ).toLocaleString("es-AR")}
            </p>

            <div className="quantity-box">
              <label>Cantidad</label>

              <select
  value={cantidad}
  onChange={(e) =>
    setCantidad(Number(e.target.value))
  }
>
  <option value={1}>1</option>
  <option value={2}>2</option>
  <option value={3}>3</option>
  <option value={4}>4</option>
  <option value={5}>5</option>
</select>
            </div>

            <div className="purchase-summary">

              <h3>Resumen de compra</h3>

             <div className="purchase-row">
  <span>Precio unitario</span>

  <span>
    $
    {Number(
      evento.precio || evento.Precio || 0
    ).toLocaleString("es-AR")}
  </span>
</div>

             <div className="purchase-row">
  <span>Cantidad</span>

  <span>{cantidad}</span>
</div>

              <div className="purchase-total">
  <span>Total</span>

  <span>
    $
    {Number(
      (evento.precio || evento.Precio || 0) *
      cantidad
    ).toLocaleString("es-AR")}
  </span>
</div>

            </div>

           <button
  className="confirm-btn"
  onClick={handleComprar}
>
  Confirmar compra
</button>

          </div>

        </div>

      </div>
    </div>
  )
}

export default CompraEntrada