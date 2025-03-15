import { BrowserRouter, Route, Routes, useLocation } from "react-router";
import { usePath } from "./store/path";
import { useEffect, useState } from "react";

import "./App.scss";
import kuma from "./assets/kuma.png";
import { Menu, PanelLeftClose } from "lucide-react";

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
			}
	
			const id = setInterval(() => {
				path.update(location.pathname.substring(1, location.pathname.length));
			}, 5000);
	
		return () => clearInterval(id);
	}, [load, path, location]);

	return (
		<main className="container-md ka-view">
			<Header />
		</main>
	);
}

function Header() {
	const [open, setOpen] = useState(false);

	return (
		<nav className="ka-nav">
			<a className="title">
				<img src={kuma} alt="" />
				<h4 className="title-content">Kuma Archive</h4>
			</a>

			<a onClick={ev => {
				ev.preventDefault();
				setOpen(!open);
			}}>
				<Menu />
			</a>
			<MenuView open={open} setOpen={setOpen} />
		</nav>
	);
}

// TODO: create menu modal
function MenuView({ open, setOpen }: { open: boolean; setOpen: (value: boolean) => void }) {
	return (
		<div className={`ka-menu ${open ? "open" : ""}`}>
			<a onClick={ev => {
				ev.preventDefault();
				setOpen(false);
			}}>
				<PanelLeftClose />
				<span>Close</span>
			</a>
		</div>
	);
}

export default App;
