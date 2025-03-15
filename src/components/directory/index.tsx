import { convert } from "../../util/unit";
import { DynamicIcon } from "lucide-react/dynamic";
import { DirEntry, usePath } from "../../store/path";

import "./directory.scss";

function Directory() {
	const path = usePath();
	if (typeof path.data === "undefined")
		return <></>;

	return (
		<div className="ka-dir">
			<div className="ka-dir-row ka-dir-top">
				<div className="ka-dir-item"></div>
				<b className="ka-dir-item">Name</b>
				<b className="ka-dir-item">Size</b>
				<b className="ka-dir-item">Date</b>
			</div>

			{path.data.path === "/" ? <></> : (
				<DirItem data={{
					name: "../",
					path: path.data.path.endsWith("/") ? path.data.path += ".." : path.data.path += "/..",
					date: -1,
					file_size: -1,
					is_dir: true,
				}} />
			)}

			{path.data.entries.map((entry, key) => {
				return <DirItem data={entry} key={key} />;
			})}
		</div>
	);
}

function DirItem({ data }: { data: DirEntry }) {
	return (
		<div className="ka-dir-row">
			<div className="ka-dir-item">
				{data.is_dir ? (
					<DynamicIcon name="folder" size={18} />
				) : <></>}
			</div>
			<a className="ka-dir-item" href={data.path}>
				{data.name}
			</a>
			<span className="ka-dir-item">
				{data.is_dir ? (
					"-"
				): convert(data.file_size)}
			</span>
			<span className="ka-dir-item">
				{data.date === -1 ? "-" : new Date(data.date * 1000).toLocaleString("en-US", {
					weekday: "short",
					year: "numeric",
					month: "short",
					day: "numeric",
					hour: "2-digit",
					minute: "2-digit",
					hour12: false
				}).replace(/,/g, "")}
			</span>
		</div>
	);
}

export default Directory;