const authService = {

  login: async (email, password) => {
    console.log("Login:", email, password)
  },

  register: async (userData) => {
    console.log("Registro:", userData)
  },

  logout: () => {
    console.log("Logout")
  }

}

export default authService