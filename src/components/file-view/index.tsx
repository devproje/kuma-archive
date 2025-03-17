import { useLocation } from "react-router";
import { useEffect, useState } from "react";
import { convert } from "../../util/unit";
import { usePath } from "../../store/path";
import { DynamicIcon, IconName } from "lucide-react/dynamic";

import "./fview.scss";

function FileView() {
	const path = usePath();
	const location = useLocation();
	const [load, setLoad] = useState(false);
	const [type, setType] = useState<IconName>("file");

	useEffect(() => {
		if (!load) {
			path.update(location.pathname.substring(1, location.pathname.length)).then(() => {
				setLoad(true);

				switch (true) {
				case path.data?.path.endsWith(".zip"):
				case path.data?.path.endsWith(".tar"):
				case path.data?.path.endsWith(".tar.gz"):
				case path.data?.path.endsWith(".tar.xz"):
				case path.data?.path.endsWith(".7z"):
				case path.data?.path.endsWith(".rar"):
					setType("file-archive");
					break;
				case path.data?.path.endsWith(".pdf"):
					setType("file-pen-line");
					break;
				case path.data?.path.endsWith(".doc"):
				case path.data?.path.endsWith(".docx"):
					setType("file-chart-pie");
					break;
				case path.data?.path.endsWith(".xls"):
				case path.data?.path.endsWith(".xlsx"):
					setType("file-spreadsheet");
					break;
				case path.data?.path.endsWith(".ppt"):
				case path.data?.path.endsWith(".pptx"):
					setType("file-sliders");
					break;
				case path.data?.path.endsWith(".jpg"):
				case path.data?.path.endsWith(".jpeg"):
				case path.data?.path.endsWith(".png"):
				case path.data?.path.endsWith(".gif"):
					setType("file-image");
					break;
				case path.data?.path.endsWith(".mp3"):
				case path.data?.path.endsWith(".wav"):
					setType("file-audio");
					break;
				case path.data?.path.endsWith(".mp4"):
				case path.data?.path.endsWith(".mkv"):
					setType("file-video");
					break;
				case path.data?.path.endsWith(".c"):
				case path.data?.path.endsWith(".cpp"):
				case path.data?.path.endsWith(".js"):
				case path.data?.path.endsWith(".ts"):
				case path.data?.path.endsWith(".jsx"):
				case path.data?.path.endsWith(".tsx"):
				case path.data?.path.endsWith(".py"):
				case path.data?.path.endsWith(".java"):
				case path.data?.path.endsWith(".rb"):
				case path.data?.path.endsWith(".go"):
				case path.data?.path.endsWith(".rs"):
				case path.data?.path.endsWith(".php"):
				case path.data?.path.endsWith(".html"):
				case path.data?.path.endsWith(".css"):
				case path.data?.path.endsWith(".scss"):
					setType("file-code");
					break;
				case path.data?.path.endsWith(".sh"):
				case path.data?.path.endsWith(".bat"):
					setType("file-terminal");
					break;
				case path.data?.path.endsWith(".json"):
					setType("file-json");
					break;
				default:
					setType("file");
					break;
				}
			});
			return;
		}
	}, [path, location, load]);

	if (typeof path.data === "undefined")
		return <></>;

	return (
		<div className="ka-fileview">
			<div className="action-row">
				<a className="link" onClick={ev => {
					ev.preventDefault();
					document.location.href = document.location.href.substring(0, document.location.href.lastIndexOf("/"));
				}}>
					<DynamicIcon name="chevron-left" />
					<span>Back to directory</span>
				</a>
			</div>
			<DynamicIcon id="icon" name={type} size={120} />
			<b>{path.data.path}</b>
			{convert(path.data.total)}

			<div id="download">
				<button className="primary" id="download" onClick={ev => {
					ev.preventDefault();
					const link = document.createElement("a");
					link.href = `/api/download${path.data?.path}`;
					link.download = `/api/download${path.data?.path}`;
					
					document.body.appendChild(link);
					link.click();
					document.body.removeChild(link);
				}}>
					<DynamicIcon name="cloud-download" />
					<span>Download</span>
				</button>
			</div>
		</div>
	);
}

export default FileView;
