import { useEffect, useState } from "react";
import Directory from "../directory";
import { usePath } from "../../store/path";
import { useLocation } from "react-router";

import "./dashboard.scss";

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
		<Directory />
	);
}

export default Dashboard;
