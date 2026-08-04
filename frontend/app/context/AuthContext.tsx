'use client'

import { API_URL } from "@/lib/api"
import { createContext, useContext, useState, useEffect, ReactNode } from "react"

type User = {
    id: string
    username: string
    email: string
}

type AuthContextType = {
    user: User | null
    loading: boolean
    refetch: () => void
}

const AuthContext = createContext<AuthContextType | null>(null)

export function AuthProvider({ children }: {children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null)
    const [loading, setLoading] = useState(true)

    const fetchMe = () => {
        fetch(`${API_URL}/me`, { credentials: 'include' })
            .then((res) => {
                if (!res.ok) throw new Error()
                return res.json()
            })
            .then((data) => setUser(data))
            .catch(() => setUser(null))
            .finally(() => setLoading(false))
    }

    useEffect(() => {
        fetchMe()
    }, [])

    return (
        <AuthContext.Provider value={{ user, loading, refetch: fetchMe }}>
            {children}
        </AuthContext.Provider>
    )
}

export function useAuth() {
    const context = useContext(AuthContext)
    if (!context) throw new Error('useAuth must be used within AuthProvider')
    return context
}