const API_URL = "http://localhost:8080"

export async function getUsuarios(token) {
  const response = await fetch(
    `${API_URL}/private/admin/usuarios`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  )

  if (!response.ok) {
    throw new Error("No se pudieron cargar los usuarios")
  }

  return response.json()
}

export async function cambiarRol(
  id,
  rol,
  token
) {
  const response = await fetch(
    `http://localhost:8080/private/admin/usuarios/${id}/rol`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        rol,
      }),
    }
  )

  const data = await response.json()

  if (!response.ok) {
    throw new Error(
      data.error || "No se pudo cambiar el rol"
    )
  }

  return data
}