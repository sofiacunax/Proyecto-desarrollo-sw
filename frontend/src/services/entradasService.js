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

  const data = await response.json()
  if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}

  if (response.status === 401) {
    localStorage.removeItem("token")
    window.location.href = "/"
    throw new Error("Sesión expirada")
  }

  if (!response.ok) {
    throw new Error(data.error || "Error al obtener entradas")
  }

  return data
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

  const data = await response.json()
  if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}

  if (response.status === 401) {
    localStorage.removeItem("token")
    window.location.href = "/"
    throw new Error("Sesión expirada")
  }

  if (!response.ok) {
    throw new Error(data.error || "Error al cancelar entrada")
  }

  return data
}

export async function transferirEntrada(
  id,
  email,
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
  email: email
}),
    }
  )

  const data = await response.json()
  if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}

  if (response.status === 401) {
    localStorage.removeItem("token")
    window.location.href = "/"
    throw new Error("Sesión expirada")
  }

  if (!response.ok) {
    throw new Error(data.error || "Error al transferir entrada")
  }

  return data
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

  const data = await response.json()
  if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}

if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}

if (!response.ok) {
  throw new Error(data.error || "Error al comprar entrada")
}

return data
}

