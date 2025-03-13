import { create } from "zustand";

interface PathState {
	data: PathResponse | string | undefined;
	update(path: string): Promise<void>;
}

interface PathResponse {
	ok: number;
	path: string;
	entries: Array<DirEntry>
}

export interface DirEntry {
	name: string;
	file_size: number;
}

export const usePath = create<PathState>((set) => ({
	data: undefined,
	update: async (path: string) => {
		const res = await fetch(`/api/path/${path}`);
		if (res.status !== 200 && res.status !== 304) {
			set({ data: undefined });
			return;
		}

		try {
			set({ data: await res.json() });
		} catch {
			set({ data: `/api/path/${path}` });
		}
	}
}));
