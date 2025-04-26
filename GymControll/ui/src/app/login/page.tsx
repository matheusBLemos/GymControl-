export default function LoginPage() {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100">
        <div className="bg-white p-8 rounded-2xl shadow-md w-full max-w-md">
          <h1 className="text-3xl font-bold text-center text-blue-900 mb-6">GymControll</h1>
          <form className="space-y-5">
            <div>
              <label className="block text-gray-700 mb-2" htmlFor="email">E-mail</label>
              <input
                type="email"
                id="email"
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Digite seu e-mail"
              />
            </div>
            <div>
              <label className="block text-gray-700 mb-2" htmlFor="password">Senha</label>
              <input
                type="password"
                id="password"
                className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Digite sua senha"
              />
            </div>
            <div className="flex items-center justify-between">
              <a href="#" className="text-sm text-blue-600 hover:underline">Esqueceu a senha?</a>
            </div>
            <button
              type="submit"
              className="w-full bg-blue-900 text-white py-2 rounded-lg hover:bg-blue-800 transition-colors"
            >
              Entrar
            </button>
          </form>
          <div className="mt-6 text-center">
            <p className="text-gray-600 text-sm">Não tem uma conta?</p>
            <a href="#" className="text-blue-600 hover:underline text-sm font-semibold">Registrar</a>
          </div>
        </div>
      </div>
    );
  }
  