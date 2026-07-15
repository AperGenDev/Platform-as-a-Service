import { useState } from "react";
import { createTask } from "../api";

function TaskForm({ onCreated }) {

    const [title, setTitle] = useState("");

    const [description, setDescription] = useState("");

    const [priority, setPriority] = useState("medium");

    async function handleSubmit(e) {

        e.preventDefault();

        try {

            await createTask({
                title,
                description,
                priority,
            });
            console.log("Task created");

            setTitle("");
            setDescription("");
            setPriority("medium");

            onCreated();

        } catch (err) {
            console.error(err);
            alert(err.message);
        }

    }

    return (

        <section className="card">

            <h2>Create Task</h2>

            <form onSubmit={handleSubmit}>

                <input
                    type="text"
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    required
                />

                <textarea
                    placeholder="Description"
                    value={description}
                    onChange={(e)=>setDescription(e.target.value)}
                />

                <select
                    value={priority}
                    onChange={(e)=>setPriority(e.target.value)}
                >
                    <option value="low">Low</option>
                    <option value="medium">Medium</option>
                    <option value="high">High</option>
                </select>

                <button type="submit">

                    Create Task

                </button>

            </form>

        </section>

    );

}

export default TaskForm;