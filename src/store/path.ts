import { create } from "zustand";

interface PathState {
	data: PathResponse | undefined;
	update(path: string): Promise<void>;
}

interface PathResponse {
	ok: number;
	path: string;
	total: number;
	is_dir: boolean;
	entries: Array<DirEntry>
}

export interface DirEntry {
	name: string;
	path: string;
	date: number;
	file_size: number;
	is_dir: boolean;
}

export const usePath = create<PathState>((set) => ({
	data: undefined,
	update: async (path: string) => {
		const res = await fetch(`/api/path/${path}`);
		if (res.status !== 200 && res.status !== 304) {
			set({ data: undefined });
			return;
		}

		set({ data: await res.json() });
	}
}));
