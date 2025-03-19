import React, { useEffect, useRef, useState } from "react";
import { AuthState, useAuthStore } from "../../store/auth";

import "./settings.scss";

function Settings() {
	const auth = useAuthStore();
	const [load, setLoad] = useState(false);

	useEffect(() => {
		if (auth.token === null) {
			document.location.href = "/";
			return;
		}

		auth.checkToken(auth.token).then((ok) => {
			if (!ok) {
				document.location.href = "/";
				return;
			}

			setLoad(true);
		});
	}, [auth, load]);

	if (!load) {
		return (
			<></>
		);
	}

	return (
		<div className="ka-settings">
			<h2>General</h2>

			<ChangePassword auth={auth} />
		</div>
	);
}

function SettingBox({ children }: { children: React.ReactNode }) {
	return (
		<div className="setting-box">
			{children}
		</div>
	);
}

function ChangePassword({ auth }: { auth: AuthState }) {
	const orRef = useRef<HTMLInputElement>(null);
	const pwRef = useRef<HTMLInputElement>(null);
	const ckRef = useRef<HTMLInputElement>(null);

	return (
		<SettingBox>
			<h4>Change Password</h4>
			<span>If you change your password, you will need to log in again.</span>
			<hr className="line" />
			<form className="box-col" id="pw-change">
				<input type="password" ref={orRef} placeholder="Password" required />
				<input type="password" ref={pwRef} placeholder="New Password" required />
				<input type="password" ref={ckRef} placeholder="Check Password" required />

				<button type="submit" className="danger" onClick={ev => {
					ev.preventDefault();
					const origin = orRef.current?.value;
					const password = pwRef.current?.value;
					const check = ckRef.current?.value;

					if (!origin || !password || !check) {
						return;
					}

					if (origin === "" || password === "" || check === "") {
						alert("You must need to write all inputs");
						return;
					}

					if (password !== check) {
						alert("New password is not matches!");
						return;
					}

					const form = new URLSearchParams();
					form.append("password", origin);
					form.append("new_password", password);

					fetch("/api/auth/update", {
						body: form,
						method: "PATCH",
						headers: {
							"Authorization": `Basic ${auth.token}`
						}
					}).then((res) => {
						if (res.status !== 200) {
							alert(`${res.status} ${res.statusText}`);
							return;
						}
						
						alert("password changed!");
						document.location.href = "/logout";
					});
				}}>Change Password</button>
			</form>
		</SettingBox>
	);
}

export default Settings;
