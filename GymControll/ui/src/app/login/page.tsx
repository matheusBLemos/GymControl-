"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";
import { loginUser } from "@/services/fatches";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const data = await loginUser(email, password);
      if (data.user) {
        Cookies.set("token", data.user);
        router.push("/");
      } else {
        alert("Falha no login");
      }
    } catch (error) {
      alert("Erro ao tentar fazer login.");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-2xl shadow-md w-full max-w-md">
        <h1 className="text-3xl font-bold text-center text-blue-900 mb-6">
          GymControll
        </h1>
        <form className="space-y-5" onSubmit={handleSubmit}>
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="email">
              E-mail
            </label>
            <input
              type="email"
              id="email"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Digite seu e-mail"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="password">
              Senha
            </label>
            <input
              type="password"
              id="password"
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Digite sua senha"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="flex items-center justify-between">
            <a href="#" className="text-sm text-blue-600 hover:underline">
              Esqueceu a senha?
            </a>
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
          <a
            href="#"
            className="text-blue-600 hover:underline text-sm font-semibold"
          >
            Registrar
          </a>
        </div>
      </div>
    </div>
  );
}
