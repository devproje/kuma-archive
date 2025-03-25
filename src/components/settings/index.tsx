import React, { useEffect, useRef, useState } from "react";
import { useAuthStore } from "../../store/auth";

import "./settings.scss";
import { DynamicIcon } from "lucide-react/dynamic";

function Settings() {
	const auth = useAuthStore();
	const [load, setLoad] = useState(false);

	useEffect(() => {
		if (auth.token === null) {
			document.location.href = "/login";
			return;
		}

		auth.checkToken(auth.token).then((ok) => {
			if (!ok) {
				document.location.href = "/login";
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
			<h2 className="ka-title">General</h2>
			<AccountSetting />

			<h2 className="ka-title">Private Directory</h2>
			<PrivateDirectory />
		</div>
	);
}

function AccountSetting() {
	return (
		<SettingBox>
			<h4>Account Setting</h4>
			<span>You can modify your account. This is a sensitive option. Please reconsider if you want to change your account information.</span>
			<hr className="line" />

			<ChangePassword />
			<RemoveAccount />
		</SettingBox>
	);
}

function PrivateDirectory() {
	return (
		<SettingBox>
			<h4>Directory Management</h4>
			<span>You can manage your private directories here. Add, remove, or modify directories as needed.</span>
			<hr className="line" />

			<CreateDirectory />
			<DirectoryTable />
		</SettingBox>
	);
}

function ChangePassword() {
	const auth = useAuthStore();
	const orRef = useRef<HTMLInputElement>(null);
	const pwRef = useRef<HTMLInputElement>(null);
	const ckRef = useRef<HTMLInputElement>(null);

	return (
		<div className="box-row">
			<div className="box-col">
				<h6>Change Password</h6>
				<span>If you change your password, you will need to log in again.</span>
			</div>

			<form className="box-col form" id="pw-change" onSubmit={ev => {
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
			}}>
				<PasswordInput placeholder="Password" ref={orRef} />
				<PasswordInput placeholder="New Password" ref={pwRef} />
				<PasswordInput placeholder="Check Password" ref={ckRef} />

				<button type="submit" className="danger">Change Password</button>
			</form>
		</div>
	);
}

function RemoveAccount() {
	const auth = useAuthStore();
	const [remove, setRemove] = useState(false);

	return (
		<div className="box-row">
			<div className="box-col">
				<h6>Delete Account</h6>
				<span>You can delete account. This action is irreversible. Please proceed with caution.</span>
			</div>

			<form className="box-col form" onSubmit={ev => {
				ev.preventDefault();

				fetch("/api/auth/delete", {
					method: "DELETE",
					headers: {
						"Authorization": `Basic ${auth.token}`
					}
				}).then((res) => {
					if (res.status !== 200) {
						alert(`${res.status} ${res.statusText}`);
						return;
					}

					alert("Your account has been deactivated and removed");
					document.location.href = "/logout";
				});
			}}>
				<label className="checkbox">
					<input type="checkbox" onChange={ev => {
						setRemove(ev.target.checked);
					}} />
					<span>I agree to remove my account.</span>
				</label>

				<button type="submit" className="danger" disabled={!remove}>Remove Account</button>
			</form>
		</div>
	);
}

function PasswordInput({ placeholder, ref }: { placeholder: string; ref: React.RefObject<HTMLInputElement | null> }) {
	const [show, setShow] = useState(false);

	return (
		<div className="input-pass">
			<input type={!show ? "password" : "text"} ref={ref} placeholder={placeholder} required />
			<a className="pw-btn" onClick={ev => {
				ev.preventDefault();
				setShow(!show);
			}}>
				<DynamicIcon name={show ? "eye-off" : "eye"} size={17} />
			</a>
		</div>
	);
}

function CreateDirectory() {
	const pathRef = useRef<HTMLInputElement>(null);
	const [disabled, setDisabled] = useState(true);

	return (
		<div className="box-row">
			<div className="box-col">
				<h6>Add Directory</h6>
				<span>You can add a private directory here.</span>
			</div>

			<form className="box-col form" onSubmit={ev => {
				ev.preventDefault();
			}}>
				<input type="text" placeholder="Directory Path" ref={pathRef} required onChange={ev => {
					setDisabled(ev.target.value === "");
				}} />
				<button type="submit" className="success" disabled={disabled}>Add Directory</button>
			</form>
		</div>
	);
}

function DirectoryTable() {
	return (
		<div className="box-col">
			<div className="box-col">
				<h6>Directories</h6>
				<span>You can view and manage your directories.</span>
			</div>

			<div className="box-col">
				{/* TODO: create table here */}
			</div>
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

export default Settings;
