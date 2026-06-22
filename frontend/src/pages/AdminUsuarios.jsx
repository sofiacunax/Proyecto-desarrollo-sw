import { useEffect, useState } from "react"
import { getUsuarios } from "../services/usuariosService"

export default function AdminUsuarios() {

  const [usuarios, setUsuarios] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {

    async function cargarUsuarios() {

      try {

        const token =
          localStorage.getItem("token")

        const data =
          await getUsuarios(token)

        setUsuarios(data)

      } catch (error) {

        console.error(error)

      } finally {

        setLoading(false)

      }
    }

    cargarUsuarios()

  }, [])

  if (loading) {
    return <p>Cargando usuarios...</p>
  }

  return (
    <div style={{ padding: "2rem" }}>
      <h1>Administración de Usuarios</h1>

      <table border="1" cellPadding="10">

        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
            <th>Email</th>
            <th>Rol</th>
          </tr>
        </thead>

        <tbody>

          {usuarios.map((usuario) => (

            <tr key={usuario.id}>

              <td>{usuario.id}</td>

              <td>{usuario.nombre}</td>

              <td>{usuario.email}</td>

              <td>{usuario.rol}</td>

            </tr>

          ))}

        </tbody>

      </table>

    </div>
  )
}