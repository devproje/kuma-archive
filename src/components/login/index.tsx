import { useRef, useState } from "react";

import "./login.scss";
import project from "../../assets/kuma.png";
import { DynamicIcon } from "lucide-react/dynamic";
import { AuthData, useAuthStore } from "../../store/auth";

function Login() {
	const [show, setShow] = useState(false);
	const auth = useAuthStore();

	const usernameRef = useRef<HTMLInputElement>(null);
	const passwordRef = useRef<HTMLInputElement>(null);
	const errnoRef = useRef<HTMLInputElement>(null);

	if (auth.token !== null)
		document.location.href = "/";

	return (
		<div className="ka-login">
			<form className="login-form">
				<div className="logo">
					<img src={project} />
					<h1>Kuma Archive Login</h1>
				</div>

				<div className="input-area">
					<div className="input">
						<div className="input-icon">
							<DynamicIcon name="user" size={20} />
						</div>
						<input ref={usernameRef} placeholder="Username" type="text" required />
						<div className="dummy"></div>
					</div>
					<div className="input">
						<div className="input-icon">
							<DynamicIcon name="key-round" size={20} />
						</div>
						<input ref={passwordRef} placeholder="Password" type={show ? "text" : "password"} required />
						<a onClick={ev => {
							ev.preventDefault();
							setShow(!show);
						}}>
							<DynamicIcon name={show ? "eye-off" : "eye"} size={20} />
						</a>
					</div>
				</div>

				<div className="submit-area">
					<span className="errno" ref={errnoRef}></span>

					<button type="submit" className="login-btn success" onClick={ev => {
						ev.preventDefault();
						const username = usernameRef.current?.value;
						const password = passwordRef.current?.value;

						if (!errnoRef)
							return;

						if (!username || !password) {
							alert("username or password is empty!");
							return;
						}
						
						if (username === "" || password === "") {
							alert("username or password is empty!");
							return;
						}

						const form = new URLSearchParams();
						form.append("username", username);
						form.append("password", password);

						errnoRef.current!.innerText = "";
						
						fetch("/api/auth/login", {
							method: "POST",
							body: form
						}).then((res) => {
							if (res.status !== 200) {
								errnoRef.current!.innerText = "username or password is not invalid";
								return;
							}

							res.json().then((data: AuthData) => {
								auth.setToken(data.token);
								window.location.href = "/";
							});
						});
					}}>
						<DynamicIcon name="log-in" size={20} />
						<span>Login</span>
					</button>
				</div>
			</form>
		</div>
	);
}

export default Login;
