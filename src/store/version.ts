import { create } from "zustand";

interface VersionState {
	value: string | undefined;
	update(): Promise<void>;
}

export const useVersion = create<VersionState>((set) => ({
	value: undefined,
	update: async () => {
		const res = await fetch("/api/version", {
			cache: "no-cache"
		});
		if (res.status !== 200 && res.status !== 304) {
			set({ value: undefined });
			return;
		}

		const text = await res.text();
		set({ value: text });
	}
}));
