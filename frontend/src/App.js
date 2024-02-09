import './App.css';
import { Routes, Route, Link } from "react-router-dom";
import Tasks from "./components/tasks/tasks";
import Settings from "./components/settings/settings";
import Servers from "./components/servers/servers";

function App() {
    return (
        <div className="h-screen w-screen overflow-x-hidden">
            <div className="text-white rounded-lg lg:w-[900px] m-auto my-5">
                <header className="text-black font-semibold text-xl p-5 flex justify-between items-center bg-white mb-5 rounded-md block-shadow">
                    <div className="w-1/2 flex gap-3">
                        <Link to={"/"} className="w-32 text-center rounded-md py-2 text-sm">ЗАДАЧИ</Link>
                        <Link to={"/settings"} className="w-32 text-center rounded-md py-2 text-sm">НАСТРОЙКИ</Link>
                        <Link to={"/servers"} className="w-32 text-center rounded-md py-2 text-sm">СЕРВЕРА</Link>
                    </div>
                    <a href="https://github.com/xALEGORx" className="text-gray text-sm font-normal">by ALEGOR</a>
                </header>
                <div className="p-5 bg-white rounded-md block-shadow">
                    <Routes>
                        <Route path="/" element={<Tasks />} />
                        <Route path="/settings" element={<Settings />} />
                        <Route path="/servers" element={<Servers />} />
                    </Routes>
                </div>
            </div>
        </div>
    );
}

export default App;
