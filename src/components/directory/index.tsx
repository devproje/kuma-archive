import { usePath } from "../../store/path";
import { useEffect } from "react";

import "./directory.scss";

function Directory() {
	const path = usePath();

	useEffect(() => {
		path.update(location.pathname);
	});

	return (
		<div></div>
	);
}

export default Directory;