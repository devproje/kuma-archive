import { useEffect, useState } from "react";
import { useRaw } from "../../store/raw";
import "./fview.scss";
import { useLocation } from "react-router";
import { DynamicIcon } from "lucide-react/dynamic";

function FileView() {
	const raw = useRaw();
	const location = useLocation();
	const [load, setLoad] = useState(false);

	useEffect(() => {
		if (!load) {
			raw.update(location.pathname.substring(1, location.pathname.length));
			setLoad(true);
			return;
		}
	}, [raw, location, load]);

	return (
		<div className="ka-fileview">
			<span className="title">
				<div className="name">
					<a className="link" href={location.pathname.endsWith("/") ? location.pathname + ".." : location.pathname + "/.."}>
						<DynamicIcon name="chevron-left" size={21} />
					</a>
					<b>{location.pathname}</b>
				</div>
				<div className="action-row">
					<a className="btn link" href={`/api/raw${location.pathname}`}>
						<DynamicIcon name="file" size={21} />
					</a>
					<a className="btn link" onClick={ev => {
						ev.preventDefault();
						navigator.clipboard.writeText(`${document.location.origin}/api/raw${location.pathname}`)
							.then(() => alert("url copied to clipboard"))
							.catch(err => console.error("Failed to copy text: ", err));
					}}>
						<DynamicIcon name="link" size={21} />
					</a>
					<a className="btn link" onClick={ev => {
						ev.preventDefault();
						fetch(`/api/download${location.pathname}`)
							.then(response => response.blob())
							.then(blob => {
								const url = window.URL.createObjectURL(blob);
								const a = document.createElement("a");

								a.style.display = "none";
								a.href = url;
								a.download = location.pathname.split("/").pop() || "download";

								document.body.appendChild(a);
								a.click();

								window.URL.revokeObjectURL(url);
							})
							.catch(error => console.error("Download failed:", error));
					}}>
						<DynamicIcon name="download" size={21} />
					</a>
				</div>
			</span>
			<pre>{raw.data}</pre>
		</div>
	);
}

export default FileView;
