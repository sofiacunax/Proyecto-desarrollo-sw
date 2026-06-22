import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import {
  getUsuarios,
  cambiarRol,
} from "../services/usuariosService"


export default function AdminUsuarios() {

  const [usuarios, setUsuarios] = useState([])
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()
 const handleCambiarRol = async (
  usuario
) => {

  const nuevoRol =
    usuario.rol === "ADMIN"
      ? "CLIENTE"
      : "ADMIN"

  try {

    const token =
      localStorage.getItem("token")

    await cambiarRol(
      usuario.id,
      nuevoRol,
      token
    )

    setUsuarios(
      usuarios.map((u) =>
        u.id === usuario.id
          ? {
              ...u,
              rol: nuevoRol,
            }
          : u
      )
    )

    alert(
      "Rol actualizado correctamente"
    )

  } catch (error) {

    console.error(error)

    alert(error.message)

  }
}
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

  <button
    onClick={() =>
      navigate("/admin/eventos")
    }
    style={{ marginBottom: "20px" }}
  >
    Volver a Eventos
  </button>

  <table>


        <thead>
          <tr>
            <th>ID</th>
            <th>Nombre</th>
            <th>Email</th>
            <th>Rol</th>
            <th>Acción</th>
          </tr>
        </thead>

        <tbody>

          {usuarios.map((usuario) => (

            <tr key={usuario.id}>

              <td>{usuario.id}</td>

              <td>{usuario.nombre}</td>

              <td>{usuario.email}</td>

              <td>{usuario.rol}</td>

              <td>
  <button
    onClick={() =>
      handleCambiarRol(usuario)
    }
  >
    Cambiar Rol
  </button>
</td>

            </tr>

          ))}

        </tbody>

      </table>

    </div>
  )
}