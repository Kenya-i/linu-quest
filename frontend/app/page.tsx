'use client'

import { useAuth } from "./context/AuthContext"

export default function Home() {
  const { user, loading } = useAuth()

  if (loading) return <p>読み込み中…</p>

  return (
    <div>
      {user ? (
        <p>ようこそ、{user.username}さん</p>
      ) : (
        <p>ログインしていません</p>
      )}
    </div>
  )
}