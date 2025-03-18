import { useAuthStore } from "../../store/auth";

function Logout() {
	const auth = useAuthStore();
	if (auth.token !== null)
		auth.clearToken();

	document.location.href = "/";

	return <>Redirecting...</>;
}

export default Logout;
