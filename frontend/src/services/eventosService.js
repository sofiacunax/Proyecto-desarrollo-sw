const API_URL = "http://localhost:8080"

export async function getEventos(search = "") {
  const response = await fetch(
    `${API_URL}/eventos?search=${encodeURIComponent(search)}`
  )
if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}
  if (!response.ok) {
    throw new Error("No se pudieron cargar los eventos")
  }

  return response.json()
}

export async function getRankingEventos() {
  const response = await fetch(`${API_URL}/eventos/ranking`)
if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}
  if (!response.ok) {
    throw new Error("No se pudo cargar el ranking")
  }

  return response.json()
}
export async function getEventoPorId(id) {
  const response = await fetch(
    `http://localhost:8080/eventos/${id}`
  )

  return await response.json()
}

export async function eliminarEvento(id, token) {
  const response = await fetch(
    `${API_URL}/private/admin/eventos/${id}`,
    {
      method: "DELETE",
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
  if (!response.ok) {
    throw new Error(data.error || "No se pudo eliminar el evento")
  }

  return data
}

export async function crearEvento(
  evento,
  token
) {
  const response = await fetch(
    `${API_URL}/private/admin/eventos`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(evento),
    }
  )

  const data = await response.json()
if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}
  if (!response.ok) {
    throw new Error(
      data.error || "No se pudo crear el evento"
    )
  }

  return data
}

export async function actualizarEvento(id, evento, token) {
  const response = await fetch(
    `${API_URL}/private/admin/eventos/${id}`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(evento),
    }
  )

  const data = await response.json()
if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}
  if (!response.ok) {
    throw new Error(data.error || "No se pudo actualizar")
  }

  return data
}
export async function getReporteEvento(
  id,
  token
) {
  const response = await fetch(
    `http://localhost:8080/private/admin/eventos/${id}/reporte`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  )
if (response.status === 401) {
  localStorage.removeItem("token")
  window.location.href = "/"
  throw new Error("Sesión expirada")
}
  if (!response.ok) {
    throw new Error(
      "No se pudo obtener el reporte"
    )
  }

  return response.json()
}