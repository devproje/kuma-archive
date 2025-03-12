import { BrowserRouter, Route, Routes } from "react-router";
import "./App.scss";

function App() {
	return (
		<BrowserRouter>
			<Routes>
				<Route path={"/"} element={<></>} />
			</Routes>
		</BrowserRouter>
	);
}

export default App;
