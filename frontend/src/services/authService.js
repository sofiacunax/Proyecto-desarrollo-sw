const API_URL = "http://localhost:8080"

export async function login(email, password) {
  const response = await fetch(
    `${API_URL}/auth/login`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email,
        password
      })
    }
  )

  const data = await response.json()

  if (!response.ok) {
    throw new Error(data.error)
  }

  return data
}

export async function register(
  nombre,
  email,
  password
) {
  const response = await fetch(
    `${API_URL}/auth/register`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        nombre,
        email,
        password
      })
    }
  )

  const data = await response.json()

  if (!response.ok) {
    throw new Error(data.error)
  }

  return data
}