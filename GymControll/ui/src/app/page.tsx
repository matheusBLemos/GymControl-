"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";

export default function Home() {
  const router = useRouter();
  const [loading, setLoading] = useState(true); // novo estado

  useEffect(() => {
    const token = Cookies.get("token");
    if (!token) {
      router.push("/login");
    } else {
      setLoading(false); // só libera a renderização se tiver token
    }
  }, [router]);

  if (loading) {
    return null; // ou pode colocar um "Loading..." se quiser
  }

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center p-8">
      <header className="w-full max-w-6xl flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-blue-900">GymControll</h1>
        <nav>
          <ul className="flex gap-6">
            <li><a href="#" className="text-blue-600 hover:underline">Dashboard</a></li>
            <li><a href="#" className="text-blue-600 hover:underline">Perfil</a></li>
            <li>
              <a
                href="#"
                className="text-blue-600 hover:underline"
                onClick={() => {
                  Cookies.remove("token");
                  router.push("/login");
                }}
              >
                Sair
              </a>
            </li>
          </ul>
        </nav>
      </header>
      <main className="w-full max-w-6xl grid grid-cols-1 md:grid-cols-2 gap-8">
        {/* Seus cards aqui */}
        <div className="bg-white p-6 rounded-2xl shadow-md flex flex-col items-start">
          <h2 className="text-2xl font-bold text-blue-800 mb-4">Cadastrar Personal</h2>
          <p className="text-gray-600 mb-6">Adicione novos personais à sua academia.</p>
          <button className="bg-blue-900 text-white px-4 py-2 rounded-lg hover:bg-blue-800">Cadastrar</button>
        </div>
        <div className="bg-white p-6 rounded-2xl shadow-md flex flex-col items-start">
          <h2 className="text-2xl font-bold text-blue-800 mb-4">Cadastrar Nutricionista</h2>
          <p className="text-gray-600 mb-6">Adicione nutricionistas para montar dietas para seus alunos.</p>
          <button className="bg-blue-900 text-white px-4 py-2 rounded-lg hover:bg-blue-800">Cadastrar</button>
        </div>
        <div className="bg-white p-6 rounded-2xl shadow-md flex flex-col items-start">
          <h2 className="text-2xl font-bold text-blue-800 mb-4">Cadastrar Aluno</h2>
          <p className="text-gray-600 mb-6">Cadastre novos alunos e acompanhe seus treinos e dietas.</p>
          <button className="bg-blue-900 text-white px-4 py-2 rounded-lg hover:bg-blue-800">Cadastrar</button>
        </div>
        <div className="bg-white p-6 rounded-2xl shadow-md flex flex-col items-start">
          <h2 className="text-2xl font-bold text-blue-800 mb-4">Gerenciar Treinos e Dietas</h2>
          <p className="text-gray-600 mb-6">Visualize e edite os planos de treino e dieta dos seus alunos.</p>
          <button className="bg-blue-900 text-white px-4 py-2 rounded-lg hover:bg-blue-800">Acessar</button>
        </div>
      </main>
    </div>
  );
}
