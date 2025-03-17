import { useEffect, useState } from "react";
import { useVersion } from "./store/version";
import FileView from "./components/file-view";
import Directory from "./components/directory";
import { DirEntry, usePath } from "./store/path";
import { DynamicIcon } from "lucide-react/dynamic";
import { BrowserRouter, Route, Routes, useLocation } from "react-router";

import "./App.scss";
import kuma from "./assets/kuma.png";
import NotFound from "./components/notfound";

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path={"*"} element={<Dashboard />} />
			</Routes>
		</BrowserRouter>
	);
}

function Dashboard() {
	const path = usePath();
	const location = useLocation();
	const [load, setLoad] = useState(false);

	useEffect(() => {
		if (!load) {
			path.update(location.pathname.substring(1, location.pathname.length)).then(() => {
				setLoad(true);
			});

			return;
		}
	}, [load, path, location]);

	if (!load) {
		return <></>;
	}

	return (
		<main className="container-md ka-view">
			<Header />
			{typeof path.data !== "undefined" ? path.data.is_dir ? <Directory /> : <FileView /> : <NotFound />}
			<Footer />
		</main>
	);
}

function Header() {
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
