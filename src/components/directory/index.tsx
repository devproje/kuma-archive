import { convert } from "../../util/unit";
import { DynamicIcon } from "lucide-react/dynamic";
import { DirEntry, usePath } from "../../store/path";

import "./directory.scss";
import Markdown from "react-markdown";
import { Suspense, useEffect, useState } from "react";
import { useLocation } from "react-router";
import { FileNavigator } from "../navigation";

function Directory() {
	const path = usePath();

	if (typeof path.data === "undefined")
		return <></>;

	return (
		<>
			<div className="ka-dir">
				<FileNavigator />
				<div className="ka-dir-row ka-dir-top">
					<div className="ka-dir-item"></div>
					<b className="ka-dir-item">Name</b>
					<b id="size" className="ka-dir-item">Size</b>
					<b id="date" className="ka-dir-item">Date</b>
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

			<Suspense fallback={<></>}>
				<Readme />
			</Suspense>
		</>
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
			<span id="size" className="ka-dir-item">
				{data.is_dir ? (
					"-"
				): convert(data.file_size)}
			</span>
			<span id="date" className="ka-dir-item">
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

function Readme() {
	const location = useLocation();
	const [load, setLoad] = useState(false);
	const [readme, setReadme] = useState<string>("");

	useEffect(() => {
		async function refresh() {
			const pathname = location.pathname;
			const res = await fetch(`/api/download${pathname}${pathname.endsWith("/") ? "" : "/"}README.md`, {
				cache: "no-cache"
			});

			if (res.status !== 200)
				return;

			setReadme(await res.text());
		}

		if (!load) {
			refresh().then(() => {
				setLoad(true);
			});
		}
	}, [load, location, setLoad, setReadme]);

	return (
		<>
			{load && readme !== "" ? (
				<div className="ka-readme">
					<Markdown>{readme}</Markdown>
				</div>
			) : <></>}
		</>
	);
}

export default Directory;