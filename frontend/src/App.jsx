import "./App.css";
import { useEffect, useState } from "react";

import Header from "./components/Header";
import HealthStatus from "./components/HealthStatus";
import TaskForm from "./components/TaskForm";
import TaskList from "./components/TaskList";

import { getTasks } from "./api";

function App() {

    const [tasks, setTasks] = useState([]);

    async function loadTasks() {
        try {
            const data = await getTasks();
            setTasks(data);
        } catch (err) {
            console.error(err);
        }
    }

    useEffect(() => {
        loadTasks();
    }, []);

    return (
        <div className="container">

            <Header />

            <HealthStatus />

            <TaskForm
                onCreated={loadTasks}
            />

            <TaskList
                tasks={tasks}
                onDeleted={loadTasks}
            />

        </div>
    );
}

export default App;