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