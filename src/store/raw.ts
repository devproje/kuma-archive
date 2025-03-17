import { create } from "zustand";

interface RawState {
	data: string | undefined;
	update(path: string): Promise<void>;
}

export const useRaw = create<RawState>((set) => ({
	data: undefined,
	update: async (path: string) => {
		const res = await fetch(`/api/raw/${path}`, {
			cache: "no-cache"
		});
		if (res.status !== 200 && res.status !== 304) {
			set({ data: undefined });
			return;
		}

		const contentType = res.headers.get("Content-Type");
		if (!contentType || !contentType.includes("text")) {
			set({ data: "this file is not supported showing preview" });
			return;
		}
		
		const text = await res.text();
		set({ data: text });	
	}
}));
