import { DynamicIcon } from "lucide-react/dynamic";
import "./notfound.scss";

function NotFound() {
	return (
		<div className="not-found">
			<DynamicIcon className="icon" name="file-question" size={120} />
			<h1>404 Not Found</h1>

			<button className="primary" onClick={ev => {
				ev.preventDefault();
				document.location.href = "/";
			}}>Back to home</button>
		</div>
	);
}

export default NotFound;
