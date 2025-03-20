import { useLocation } from "react-router";

export function FileNavigator() {
	const location = useLocation();
	const split = location.pathname.substring(1, location.pathname.length).split("/");

	return <div>
		
	</div>;
}
