const API_URL = "http://localhost:8080"

export async function obtenerMisEntradas(token) {
  const response = await fetch(
    `${API_URL}/private/mis-entradas`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  )

  return await response.json()
}

export async function cancelarEntrada(id, token) {
  const response = await fetch(
    `${API_URL}/private/entradas/${id}/cancelar`,
    {
      method: "PUT",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  )

  return await response.json()
}

export async function transferirEntrada(
  id,
  nuevoUsuarioID,
  token
) {
  const response = await fetch(
    `${API_URL}/private/entradas/${id}/transferir`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        nuevo_usuario_id: nuevoUsuarioID,
      }),
    }
  )

  return await response.json()
}

export async function comprarEntrada(
  eventoID,
  token
) {
  const response = await fetch(
    `${API_URL}/private/entradas`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        evento_id: eventoID,
      }),
    }
  )

  return await response.json()
}