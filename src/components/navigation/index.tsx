import { useLocation } from "react-router";

export function FileNavigator() {
	const location = useLocation();
	const split = location.pathname === "/" ? Array<string>() : location.pathname.substring(1, location.pathname.length).split("/");

	return (
		<div>
			<a href="/">
				<span>Index Directory</span>
			</a>
			{split.map((path, i) => {
				if (i === split.length - 1) {
					return (
						<a key={i}>
							<span>{path}</span>
						</a>
					);
				}

				return (
					<a key={i} href="">
						<span>{path}</span>
						<span>&gt;&gt;</span>
					</a>
				);
			})}
		</div>
	);
}
