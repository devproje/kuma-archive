import { create } from "zustand";

interface PathState {
	data: undefined;
	update(path: string): Promise<void>;
}

export const usePath = create<PathState>((set) => ({
	data: undefined,
	update: async (path: string) => {
		console.log(path);
		set({ data: undefined });
	}
}));
