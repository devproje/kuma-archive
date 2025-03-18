import { useAuthStore } from "../../store/auth";

function Logout() {
	const auth = useAuthStore();
	auth.clearToken();
	
	document.location.href = "/";
	return <></>;
}

export default Logout;
