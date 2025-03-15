import { useEffect, useState } from "react";
import Directory from "./components/directory";
import { DirEntry, usePath } from "./store/path";
import { DynamicIcon } from "lucide-react/dynamic";
import { BrowserRouter, Route, Routes, useLocation } from "react-router";

import "./App.scss";
import kuma from "./assets/kuma.png";
import FileView from "./components/file-view";

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
			path.update(location.pathname.substring(1, location.pathname.length));
			setLoad(true);

			return;
		}

		const id = setInterval(() => {
			path.update(location.pathname.substring(1, location.pathname.length));
		}, 5000);
	
		return () => clearInterval(id);
	}, [load, path, location]);

	if (!load) {
		return <></>;
	}

	return (
		<main className="container-md ka-view">
			<Header />
			{typeof path.data !== "undefined" ? path.data.is_dir ? <Directory /> : <FileView /> : (
				<>
					<h1>404 Not Found</h1>

					<button className="primary" onClick={ev => {
						ev.preventDefault();
						document.location.href = "/";
					}}>Back to home</button>
				</>
			)}
			
			<Footer />
		</main>
	);
}

function Header() {
	// const [open, setOpen] = useState(false);

	return (
		<nav className="ka-nav">
			<a className="title" href="/">
				<img src={kuma} alt="" />
				<h4 className="title-content">Kuma Archive</h4>
			</a>

			<a onClick={ev => {
				ev.preventDefault();
				console.log("not provide features");
				// setOpen(!open);
			}}>
				<DynamicIcon className="link" name="more-vertical" />
			</a>
			{/* <MenuView open={open} setOpen={setOpen} /> */}
		</nav>
	);
}

// TODO: create menu modal
// function MenuView({ open, setOpen }: { open: boolean; setOpen: (value: boolean) => void }) {
// 	return (
// 		<div className={`ka-menu ${open ? "open" : ""}`}>
// 			<MenuItem icon="panel-left-close" name="Close" block={() => {
// 				setOpen(false);
// 			}} />
// 		</div>
// 	);
// }

// function MenuItem({ icon, name, block }: { icon: IconName, name: string, block?: () => void }) {
// 	return (
// 		<a className={"ka-menu-item link"} onClick={ev => {
// 			ev.preventDefault();
// 			if (typeof block === "undefined")
// 				return;

// 			block();
// 		}}>
// 			<DynamicIcon name={icon} />
// 			<span>{name}</span>
// 		</a>
// 	);
// }

function Footer() {
	const path = usePath();
	let file = 0;
	let dir = 0;

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

	return (
		<footer className="ka-footer">
			{path.data ? path.data.is_dir ? (
				<div className="searched">
					Found {dir === 1 ? `${dir} directory` : `${dir} directories`}, {file === 1 ? `${file} file` : `${file} files`}
				</div>
			) : <></> : <></>}
			
			<div className="footer">
				&copy; 2020-2025 Project_IO. MIT License. Powered by WSERVER.
			</div>
		</footer>
	);
}

export default App;
