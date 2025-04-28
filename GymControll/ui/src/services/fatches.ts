//aqui ficam os fatchees:
export async function loginUser(email: string, password: string) {
    try {
      const response = await fetch("http://localhost:5000/gymcontoll/api/v1/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });
  
      if (!response.ok) {
        throw new Error("Falha na requisição");
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error("Erro no login:", error);
      throw error;
    }
  }
  