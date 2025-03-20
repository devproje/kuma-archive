import { useLocation } from "react-router";

import "./navigation.scss";

export function FileNavigator() {
	const location = useLocation();
	const split = location.pathname === "/" ? Array<string>() : location.pathname.substring(1, location.pathname.length).split("/");

	return (
		<div className="ka-navigator">
			{location.pathname === "/" ? (
				<span className="current">Index Directory</span>
			) : (
				<a href="/">
					<span>Index Directory</span>
				</a>
			)}
			{split.map((path, i) => {
				let route = "";
				split.forEach((str, j) => {
					if (j > i)
						return;

					route += `/${str}`;
				});

				return (
					<>
						<span className="heap">&gt;</span>
						{i === split.length - 1 ? (
							<div key={i}>
								<span className="current">{decodeURIComponent(path)}</span>
							</div>
						) : (
							<a key={i} href={route}>
								<span>{decodeURIComponent(path)}</span>
							</a>
						)}
					</>
				);
			})}
		</div>
	);
}
