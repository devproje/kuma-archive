import Login from "./components/login";
import Logout from "./components/logout";
import React, { useEffect, useState } from "react";
import Settings from "./components/settings";
import { useVersion } from "./store/version";
import NotFound from "./components/notfound";
import FileView from "./components/file-view";
import Directory from "./components/directory";
import { DirEntry, usePath } from "./store/path";
import { DynamicIcon } from "lucide-react/dynamic";
import { BrowserRouter, Route, Routes, useLocation } from "react-router";
import { AccountData, useAuthStore } from "./store/auth";

import "./App.scss";
import kuma from "./assets/kuma.png";

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/login" element={<Dashboard children={<Login />} />} />
				<Route path="/logout" element={<Logout />} />
				<Route path="/settings" element={<Dashboard children={<Settings />} />} />
				<Route path={"*"} element={<Dashboard children={<View />} />} />
			</Routes>
		</BrowserRouter>
	);
}

function Dashboard({ children }: { children: React.ReactNode }) {
	return (
		<main className="container-md ka-view">
			<Header />
			{children}
			<Footer />
		</main>
	);
}

function View() {
	const path = usePath();
	const auth = useAuthStore();
	const location = useLocation();
	const [load, setLoad] = useState(false);

	useEffect(() => {
		if (!load) {
			path.update(location.pathname.substring(1, location.pathname.length), auth.token)
				.then(() => {
					setLoad(true);
				});

			return;
		}
	}, [auth, load, path, location]);

	if (!load) {
		return <></>;
	}

	if (typeof path.data === "undefined") {
		return <NotFound />;
	}

	if (path.data.is_dir) {
		return <Directory />;
	}

	return <FileView />;
}

function Header() {
	const auth = useAuthStore();
	const [isAuth, setAuth] = useState(false);
	const [username, setUsername] = useState("undefined");

	useEffect(() => {
		if (auth.token === null) {
			return;
		}

		auth.checkToken(auth.token).then((ok) => {
			if (ok)
				setAuth(true);
		});

		fetch("/api/auth/read", {
			method: "GET",
			mode: "same-origin",
			headers: {
				"Authorization": `Basic ${auth.token}`
			}
		}).then(res => {
			if (res.status !== 200)
				return;

			return res.json();
		}).then((data: AccountData) => {
			setUsername(data.username);
		});
	}, [auth, isAuth]);

	return (
		<nav className="ka-nav">
			<a className="title" href="/">
				<img src={kuma} alt="" />
				<h4 className="title-content">Kuma Archive</h4>
			</a>

			<div className="action-row">
				<a className="link" href="https://git.wh64.net/devproje/kuma-archive">
					<DynamicIcon name="folder-git-2" size={15} />
				</a>
				<a className="link" href="https://projecttl.net">
					<DynamicIcon name="globe" size={15} />
				</a>
				
				{!isAuth ? (
					<a className="login-btn" href="/login">
						Login
					</a>
				) : (
					<>
						<a className="link" href="/settings">
							<DynamicIcon name="settings" size={15} />
						</a>
						<div className="login-info">
							<span className="username">Logged in as {username}</span>
							<a className="login-btn" href="/logout">
								Logout
							</a>
						</div>
					</>
				)}
				
			</div>
		</nav>
	);
}

function Footer() {
	const path = usePath();
	let file = 0;
	let dir = 0;

	const version = useVersion();
	const [load, setLoad] = useState(false);
	
	if (typeof path.data !== "undefined") {
		if (path.data.is_dir) {
			path.data.entries.forEach((entry: DirEntry) => {
				if (entry.is_dir) {
					dir += 1;
				} else {
					file += 1;
				}
			});
		}
	}

	useEffect(() => {
		if (!load) {
			version.update().then(() => {
				setLoad(true);
			});
			return;
		}
	}, [load, version]);

	return (
		<footer className="ka-footer">
			{path.data ? path.data.is_dir ? (
				<div className="searched">
					Found {dir === 1 ? `${dir} directory` : `${dir} directories`}, {file === 1 ? `${file} file` : `${file} files`}
				</div>
			) : <></> : <></>}
			
			<div className="footer">
				<span><b>Kuma Archive</b> {version.value}</span>
				<p>
					&copy; 2020-2025 <a href="https://git.wh64.net/devproje">Project_IO</a>. All rights reserved for images.
					<br />
					Code licensed under the <a href="https://git.wh64.net/devproje/kuma-archive/src/branch/master/LICENSE">MIT License</a>.
				</p>
				<span>Powered by WSERVER</span>
			</div>
		</footer>
	);
}

export default App;
