import { create } from "zustand";
import { persist } from "zustand/middleware";

export interface AuthData {
	ok: number;
	token: string;
}


export interface AuthState {
	token: string | null;
	setToken: (token: string) => void;
	clearToken: () => void;
	checkToken: (token: string) => Promise<boolean>;
}

export const useAuthStore = create<AuthState>()(
	persist(
		(set) => ({
			token: null,
			setToken: (token) => set({ token }),
			clearToken: () => set({ token: null }),
			checkToken: async (token: string) => {
				const res = await fetch("/api/auth/check", {
					headers: {
						"Authorization": `Basic ${token}`
					}
				});

				return res.status === 200;
			}
		}),
		{
			name: "auth-storage"
		}
	)
);

