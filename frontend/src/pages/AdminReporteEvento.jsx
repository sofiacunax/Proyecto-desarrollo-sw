import { useEffect, useState } from "react"
import { useNavigate, useParams } from "react-router-dom"

import {
  getReporteEvento,
} from "../services/eventosService"

export default function AdminReporteEvento() {

  const { id } = useParams()

  const navigate = useNavigate()

  const [reporte, setReporte] =
    useState(null)

  useEffect(() => {

    async function cargarReporte() {

      try {

        const token =
          localStorage.getItem("token")

        const data =
          await getReporteEvento(
            id,
            token
          )

        setReporte(data)

      } catch (error) {

        console.error(error)

      }
    }

    cargarReporte()

  }, [id])

  if (!reporte) {
    return <p>Cargando reporte...</p>
  }

  return (
    <div style={{ padding: "2rem" }}>

      <button
        onClick={() =>
          navigate("/admin/eventos")
        }
      >
        Volver
      </button>

      <h1>
        Reporte de Evento
      </h1>

      <h2>
        {reporte.titulo}
      </h2>

      <hr />

      <h3>Métricas</h3>

      <p>
        Capacidad:
        {" "}
        {reporte.capacidad}
      </p>

      <p>
        Entradas vendidas:
        {" "}
        {reporte.entradas_vendidas}
      </p>

      <p>
        Ocupación:
        {" "}
        {reporte.porcentaje_ocupado.toFixed(2)}
        %
      </p>

      <hr />

      <h3>
        Compradores
      </h3>

      {reporte.compradores?.length === 0 ? (
        <p>
          No hay compradores
        </p>
      ) : (
        <ul>

          {reporte.compradores?.map(
            (comprador) => (
              <li
                key={comprador.id}
              >
                {comprador.nombre}
                {" - "}
                {comprador.email}
              </li>
            )
          )}

        </ul>
      )}

    </div>
  )
}